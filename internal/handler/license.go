package handler

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"program-verify/internal/model"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func generateLicenseKey() (string, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes), nil
}

// GetLicenses 获取许可证列表
func GetLicenses(c *gin.Context) {
	// 获取筛选参数
	productID := c.Query("product_id")
	licenseKey := c.Query("license_key")
	activatedStartDate := c.Query("activated_start_date")
	activatedEndDate := c.Query("activated_end_date")
	expiresStartDate := c.Query("expires_start_date")
	expiresEndDate := c.Query("expires_end_date")
	activationStatus := c.Query("activation_status")
	enableStatus := c.Query("enable_status")
	remark := c.Query("remark")

	// 构建查询条件
	whereClause := "WHERE 1=1"
	args := []interface{}{}
	if activationStatus != "" {
		whereClause += " AND l.activation_status = ?"
		args = append(args, activationStatus)
	}
	if enableStatus != "" {
		whereClause += " AND l.enable_status = ?"
		args = append(args, enableStatus)
	}
	if productID != "" {
		if productIDInt, err := strconv.ParseInt(productID, 10, 64); err == nil {
			whereClause += " AND l.product_id = ?"
			args = append(args, productIDInt)
		}
	}
	if licenseKey != "" {
		whereClause += " AND l.license_key LIKE ?"
		args = append(args, "%"+licenseKey+"%")
	}
	if activatedStartDate != "" {
		whereClause += " AND l.activated_at >= ?"
		args = append(args, activatedStartDate)
	}
	if activatedEndDate != "" {
		whereClause += " AND l.activated_at <= ?"
		args = append(args, activatedEndDate)
	}
	if expiresStartDate != "" {
		whereClause += " AND l.expires_at >= ?"
		args = append(args, expiresStartDate)
	}
	if expiresEndDate != "" {
		whereClause += " AND l.expires_at <= ?"
		args = append(args, expiresEndDate)
	}
	if remark != "" {
		whereClause += " AND l.remark LIKE ?"
		args = append(args, "%"+remark+"%")
	}

	// 获取总数
	var total int
	countQuery := "SELECT COUNT(*) FROM licenses l " + whereClause
	err := model.DB.QueryRow(countQuery, args...).Scan(&total)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Error(http.StatusInternalServerError, "获取许可证总数失败"))
		return
	}

	// 检查是否有分页参数
	pageStr := c.Query("page")
	pageSizeStr := c.Query("page_size")

	// 如果没有分页参数，返回所有数据
	if pageStr == "" && pageSizeStr == "" {
		query := `
			SELECT 
				l.id, l.license_key, l.product_id, l.activation_status, 
				l.enable_status, l.activated_at, l.expires_at, l.remark,
				p.name as product_name
			FROM licenses l
			LEFT JOIN products p ON l.product_id = p.id
			` + whereClause + `
			ORDER BY l.created_at DESC
		`

		rows, err := model.DB.Query(query, args...)
		if err != nil {
			c.JSON(http.StatusInternalServerError, model.Error(http.StatusInternalServerError, "获取许可证列表失败"))
			return
		}
		defer rows.Close()

		var licenses []model.LicenseWithProduct
		for rows.Next() {
			var l model.LicenseWithProduct
			var productName string
			if err := rows.Scan(
				&l.ID, &l.Key, &l.Product.ID, &l.ActivationStatus,
				&l.EnableStatus, &l.ActivatedAt, &l.ExpiresAt, &l.Remark,
				&productName,
			); err != nil {
				c.JSON(http.StatusInternalServerError, model.Error(http.StatusInternalServerError, "解析许可证数据失败"))
				return
			}
			l.Product.Name = productName
			licenses = append(licenses, l)
		}

		c.JSON(http.StatusOK, model.Success(gin.H{
			"total": total,
			"items": licenses,
		}, "获取许可证列表成功"))
		return
	}

	// 有分页参数时的处理逻辑
	page, _ := strconv.Atoi(pageStr)
	pageSize, _ := strconv.Atoi(pageSizeStr)
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	// 计算总页数
	totalPages := (total + pageSize - 1) / pageSize
	if page > totalPages {
		page = totalPages
	}

	offset := (page - 1) * pageSize

	// 获取分页数据
	query := `
		SELECT 
			l.id, l.license_key, l.product_id, l.activation_status, 
			l.enable_status, l.activated_at, l.expires_at, l.remark,
			p.name as product_name
		FROM licenses l
		LEFT JOIN products p ON l.product_id = p.id
		` + whereClause + `
		ORDER BY l.created_at DESC
		LIMIT ? OFFSET ?
	`
	args = append(args, pageSize, offset)

	rows, err := model.DB.Query(query, args...)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Error(http.StatusInternalServerError, "获取许可证列表失败"))
		return
	}
	defer rows.Close()

	var licenses []model.LicenseWithProduct
	for rows.Next() {
		var l model.LicenseWithProduct
		var productName string
		if err := rows.Scan(
			&l.ID, &l.Key, &l.Product.ID, &l.ActivationStatus,
			&l.EnableStatus, &l.ActivatedAt, &l.ExpiresAt, &l.Remark,
			&productName,
		); err != nil {
			c.JSON(http.StatusInternalServerError, model.Error(http.StatusInternalServerError, "解析许可证数据失败"))
			return
		}
		l.Product.Name = productName
		licenses = append(licenses, l)
	}

	c.JSON(http.StatusOK, model.Success(gin.H{
		"total":       total,
		"total_pages": totalPages,
		"page":        page,
		"page_size":   pageSize,
		"items":       licenses,
	}, "获取许可证列表成功"))
}

// GetLicense 获取单个许可证
func GetLicense(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Error(http.StatusBadRequest, "无效的许可证ID"))
		return
	}

	var license model.LicenseWithProduct
	err = model.DB.QueryRow(`
		SELECT 
			l.id, 
			l.license_key, 
			l.duration_days,
			l.created_at,
			l.activated_at,
			l.expires_at,
			l.activation_status,
			l.enable_status,
			l.remark,
			p.id as product_id,
			p.name as product_name,
			p.price as product_price,
			p.created_at as product_created_at,
			p.status as product_status
		FROM licenses l
		LEFT JOIN products p ON l.product_id = p.id
		WHERE l.id = ?
	`, id).Scan(
		&license.ID,
		&license.Key,
		&license.DurationDays,
		&license.CreatedAt,
		&license.ActivatedAt,
		&license.ExpiresAt,
		&license.ActivationStatus,
		&license.EnableStatus,
		&license.Remark,
		&license.Product.ID,
		&license.Product.Name,
		&license.Product.Price,
		&license.Product.CreatedAt,
		&license.Product.Status,
	)

	if err != nil {
		c.JSON(http.StatusNotFound, model.Error(http.StatusNotFound, "许可证不存在"))
		return
	}

	c.JSON(http.StatusOK, model.Success(license, "获取许可证成功"))
}

// CreateLicense 创建许可证
func CreateLicense(c *gin.Context) {
	var req model.CreateLicenseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Error(http.StatusBadRequest, "请求参数错误"))
		return
	}

	// 验证产品是否存在
	var exists bool
	err := model.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM products WHERE id = ?)", req.ProductID).Scan(&exists)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Error(http.StatusInternalServerError, "检查产品失败"))
		return
	}
	if !exists {
		c.JSON(http.StatusBadRequest, model.Error(http.StatusBadRequest, "产品不存在"))
		return
	}

	// 开始事务
	tx, err := model.DB.Begin()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Error(http.StatusInternalServerError, "开始事务失败"))
		return
	}

	// 批量生成许可证
	var licenses []gin.H
	now := time.Now()
	for i := 0; i < req.Count; i++ {
		// 生成许可证密钥
		licenseKey, err := generateLicenseKey()
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, model.Error(http.StatusInternalServerError, "生成许可证密钥失败"))
			return
		}

		// 创建许可证
		result, err := tx.Exec(
			"INSERT INTO licenses (product_id, license_key, duration_days, created_at, activation_status, enable_status, remark) VALUES (?, ?, ?, ?, ?, ?, ?)",
			req.ProductID, licenseKey, req.DurationDays, now, model.LicenseActivationStatusInactive, model.LicenseEnableStatusEnabled, req.Remark,
		)
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, model.Error(http.StatusInternalServerError, "创建许可证失败"))
			return
		}

		id, err := result.LastInsertId()
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, model.Error(http.StatusInternalServerError, "获取许可证ID失败"))
			return
		}

		licenses = append(licenses, gin.H{
			"id":                id,
			"license_key":       licenseKey,
			"duration_days":     req.DurationDays,
			"created_at":        now,
			"activation_status": model.LicenseActivationStatusInactive,
			"enable_status":     model.LicenseEnableStatusEnabled,
			"remark":            req.Remark,
		})
	}

	// 提交事务
	if err := tx.Commit(); err != nil {
		c.JSON(http.StatusInternalServerError, model.Error(http.StatusInternalServerError, "提交事务失败"))
		return
	}

	c.JSON(http.StatusCreated, model.Success(gin.H{
		"count":    req.Count,
		"licenses": licenses,
	}, "创建许可证成功"))
}

// UpdateLicense 更新许可证
func UpdateLicense(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Error(http.StatusBadRequest, "无效的许可证ID"))
		return
	}

	var req model.UpdateLicenseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Error(http.StatusBadRequest, "请求参数错误"))
		return
	}

	// 检查许可证是否存在
	var exists bool
	err = model.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM licenses WHERE id = ?)", id).Scan(&exists)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Error(http.StatusInternalServerError, "检查许可证失败"))
		return
	}
	if !exists {
		c.JSON(http.StatusNotFound, model.Error(http.StatusNotFound, "许可证不存在"))
		return
	}

	// 构建更新SQL
	var updateFields []string
	var args []interface{}

	// 处理启用状态更新
	if req.EnableStatus != "" {
		updateFields = append(updateFields, "enable_status = ?")
		args = append(args, req.EnableStatus)
	}

	// 处理备注更新
	if req.Remark != "" {
		updateFields = append(updateFields, "remark = ?")
		args = append(args, req.Remark)
	}

	// 处理过期时间更新
	if req.ExpiresAt != nil {
		updateFields = append(updateFields, "expires_at = ?")
		args = append(args, req.ExpiresAt)

		// 如果许可证已激活，重新计算有效期天数
		var activatedAt *time.Time
		err = model.DB.QueryRow("SELECT activated_at FROM licenses WHERE id = ?", id).Scan(&activatedAt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, model.Error(http.StatusInternalServerError, "获取激活时间失败"))
			return
		}

		if activatedAt != nil {
			durationDays := int(req.ExpiresAt.Sub(*activatedAt).Hours() / 24)
			updateFields = append(updateFields, "duration_days = ?")
			args = append(args, durationDays)
		}
	}

	// 如果没有要更新的字段，返回错误
	if len(updateFields) == 0 {
		c.JSON(http.StatusBadRequest, model.Error(http.StatusBadRequest, "没有要更新的字段"))
		return
	}

	// 添加ID到参数列表
	args = append(args, id)

	// 构建并执行更新SQL
	query := "UPDATE licenses SET " + strings.Join(updateFields, ", ") + " WHERE id = ?"
	_, err = model.DB.Exec(query, args...)

	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Error(http.StatusInternalServerError, "更新许可证失败"))
		return
	}

	// 获取更新后的许可证信息
	var license model.LicenseWithProduct
	err = model.DB.QueryRow(`
		SELECT 
			l.id, 
			l.license_key, 
			l.duration_days,
			l.created_at,
			l.activated_at,
			l.expires_at,
			l.activation_status,
			l.enable_status,
			l.remark,
			p.name as product_name,
			p.price as product_price,
			p.created_at as product_created_at,
			p.status as product_status
		FROM licenses l
		LEFT JOIN products p ON l.product_id = p.id
		WHERE l.id = ?
	`, id).Scan(
		&license.ID,
		&license.Key,
		&license.DurationDays,
		&license.CreatedAt,
		&license.ActivatedAt,
		&license.ExpiresAt,
		&license.ActivationStatus,
		&license.EnableStatus,
		&license.Remark,
		&license.Product.Name,
		&license.Product.Price,
		&license.Product.CreatedAt,
		&license.Product.Status,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Error(http.StatusInternalServerError, "获取更新后的许可证信息失败"))
		return
	}

	// 根据更新的字段返回不同的提示信息
	var message string
	if req.EnableStatus != "" {
		if req.EnableStatus == string(model.LicenseEnableStatusEnabled) {
			message = "启用许可证成功"
		} else if req.EnableStatus == string(model.LicenseEnableStatusDisabled) {
			message = "禁用许可证成功"
		}
	} else if req.ExpiresAt != nil {
		message = "更新过期时间成功"
	} else if req.Remark != "" {
		message = "更新备注成功"
	} else {
		message = "更新许可证成功"
	}

	c.JSON(http.StatusOK, model.Success(license, message))
}

// VerifyLicense 验证/激活许可证
func VerifyLicense(c *gin.Context) {
	var req model.VerifyLicenseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Error(http.StatusBadRequest, "请求参数错误"))
		return
	}

	var license model.License
	var productName string
	var productStatus model.ProductStatus
	err := model.DB.QueryRow(`
		SELECT l.id, l.product_id, l.license_key, l.activation_status, l.expires_at,
			   l.duration_days, p.name as product_name, p.status as product_status
		FROM licenses l
		JOIN products p ON l.product_id = p.id
		WHERE l.license_key = ?
	`, req.LicenseKey).Scan(&license.ID, &license.ProductID, &license.Key,
		&license.ActivationStatus, &license.ExpiresAt, &license.DurationDays, &productName,
		&productStatus)

	if err != nil {
		c.JSON(http.StatusOK, model.Success(model.VerifyLicenseResponse{
			Valid:   false,
			Message: "无效的许可证",
		}, "许可证验证失败"))
		return
	}

	// 检查许可证是否被禁用
	if license.EnableStatus == model.LicenseEnableStatusDisabled {
		c.JSON(http.StatusOK, model.Success(model.VerifyLicenseResponse{
			Valid:   false,
			Message: "许可证已被禁用",
		}, "许可证验证失败"))
		return
	}

	// 检查产品是否被禁用
	if productStatus == model.ProductStatusDisabled {
		c.JSON(http.StatusOK, model.Success(model.VerifyLicenseResponse{
			Valid:   false,
			Message: "产品已被禁用",
		}, "许可证验证失败"))
		return
	}

	// 检查许可证状态
	if license.ActivationStatus == model.LicenseActivationStatusExpired {
		c.JSON(http.StatusOK, model.Success(model.VerifyLicenseResponse{
			Valid:   false,
			Message: "许可证已过期",
		}, "许可证验证失败"))
		return
	}

	// 如果未激活，则激活许可证
	if license.ActivationStatus == model.LicenseActivationStatusInactive {
		now := time.Now()
		expiresAt := now.AddDate(0, 0, license.DurationDays)

		// 更新许可证状态
		_, err = model.DB.Exec(`
			UPDATE licenses 
			SET activated_at = ?, expires_at = ?, activation_status = ?
			WHERE id = ?
		`, now, expiresAt, model.LicenseActivationStatusActivated, license.ID)

		if err != nil {
			c.JSON(http.StatusInternalServerError, model.Error(http.StatusInternalServerError, "激活许可证失败"))
			return
		}

		c.JSON(http.StatusOK, model.Success(model.VerifyLicenseResponse{
			Valid: true,
			License: model.LicenseInfo{
				ProductName: productName,
				ActivatedAt: now,
				ExpiresAt:   expiresAt,
			},
			Message: "许可证激活成功",
		}, "许可证验证成功"))
		return
	}

	// 检查是否过期
	if time.Now().After(*license.ExpiresAt) {
		// 更新状态为过期
		_, err = model.DB.Exec("UPDATE licenses SET activation_status = ? WHERE id = ?",
			model.LicenseActivationStatusExpired, license.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, model.Error(http.StatusInternalServerError, "更新许可证状态失败"))
			return
		}

		c.JSON(http.StatusOK, model.Success(model.VerifyLicenseResponse{
			Valid:   false,
			Message: "许可证已过期",
		}, "许可证验证失败"))
		return
	}

	c.JSON(http.StatusOK, model.Success(model.VerifyLicenseResponse{
		Valid: true,
		License: model.LicenseInfo{
			ProductName: productName,
			ActivatedAt: *license.ActivatedAt,
			ExpiresAt:   *license.ExpiresAt,
		},
	}, "许可证验证成功"))
}

// BatchGenerateLicense 批量生成许可证
func BatchGenerateLicense(c *gin.Context) {
	var req model.BatchGenerateLicenseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Error(http.StatusBadRequest, "请求参数错误"))
		return
	}

	// 验证产品是否存在
	var exists bool
	err := model.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM products WHERE id = ?)", req.ProductID).Scan(&exists)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Error(http.StatusInternalServerError, "检查产品失败"))
		return
	}
	if !exists {
		c.JSON(http.StatusBadRequest, model.Error(http.StatusBadRequest, "产品不存在"))
		return
	}

	// 开始事务
	tx, err := model.DB.Begin()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Error(http.StatusInternalServerError, "开始事务失败"))
		return
	}

	// 批量生成许可证
	var licenses []gin.H
	for i := 0; i < req.Count; i++ {
		// 生成许可证密钥
		licenseKey, err := generateLicenseKey()
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, model.Error(http.StatusInternalServerError, "生成许可证密钥失败"))
			return
		}

		// 创建许可证
		result, err := tx.Exec(
			"INSERT INTO licenses (product_id, license_key, duration_days, activation_status, enable_status, remark) VALUES (?, ?, ?, ?, ?, ?)",
			req.ProductID, licenseKey, req.DurationDays, model.LicenseActivationStatusInactive, model.LicenseEnableStatusEnabled, req.Remark,
		)
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, model.Error(http.StatusInternalServerError, "创建许可证失败"))
			return
		}

		id, err := result.LastInsertId()
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, model.Error(http.StatusInternalServerError, "获取许可证ID失败"))
			return
		}

		licenses = append(licenses, gin.H{
			"id":                id,
			"license_key":       licenseKey,
			"activation_status": model.LicenseActivationStatusInactive,
			"enable_status":     model.LicenseEnableStatusEnabled,
			"remark":            req.Remark,
		})
	}

	// 提交事务
	if err := tx.Commit(); err != nil {
		c.JSON(http.StatusInternalServerError, model.Error(http.StatusInternalServerError, "提交事务失败"))
		return
	}

	c.JSON(http.StatusCreated, model.Success(gin.H{
		"count":    req.Count,
		"licenses": licenses,
	}, "生成许可证成功"))
}

// GetLicenseStats 获取许可证统计信息
func GetLicenseStats(c *gin.Context) {
	// 获取活跃许可证数量和总收益
	query := `
		SELECT 
			COUNT(CASE WHEN l.activation_status = 'activated' AND l.enable_status = 'enabled' THEN 1 END) as active_licenses,
			COALESCE(SUM(CASE WHEN l.activation_status = 'activated' AND l.enable_status = 'enabled' THEN p.price END), 0) as total_revenue
		FROM licenses l
		LEFT JOIN products p ON l.product_id = p.id
	`

	var stats struct {
		ActiveLicenses int     `json:"active_licenses"`
		TotalRevenue   float64 `json:"total_revenue"`
	}

	err := model.DB.QueryRow(query).Scan(&stats.ActiveLicenses, &stats.TotalRevenue)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Error(http.StatusInternalServerError, "获取统计信息失败"))
		return
	}

	c.JSON(http.StatusOK, model.Success(stats, "获取统计信息成功"))
}

// GetLicenseAnalytics 获取许可证分析数据
func GetLicenseAnalytics(c *gin.Context) {
	// 获取产品转化率
	productConversionQuery := `
		SELECT 
			p.id,
			p.name,
			COUNT(l.id) as total_licenses,
			COUNT(CASE WHEN l.activation_status = 'activated' THEN 1 END) as activated_licenses,
			COALESCE(SUM(CASE WHEN l.activation_status = 'activated' THEN p.price END), 0) as revenue
		FROM products p
		LEFT JOIN licenses l ON p.id = l.product_id
		WHERE p.status = 'enabled'
		GROUP BY p.id, p.name
		ORDER BY revenue DESC
		LIMIT 10
	`

	// 获取时间趋势
	trendQuery := `
		SELECT 
			DATE(l.created_at) as date,
			COUNT(*) as new_licenses,
			COUNT(CASE WHEN l.activation_status = 'activated' THEN 1 END) as activated_licenses,
			COALESCE(SUM(CASE WHEN l.activation_status = 'activated' THEN p.price END), 0) as daily_revenue
		FROM licenses l
		LEFT JOIN products p ON l.product_id = p.id
		WHERE l.created_at >= DATE_SUB(CURDATE(), INTERVAL 30 DAY)
		GROUP BY DATE(l.created_at)
		ORDER BY date DESC
	`

	// 获取有效期分布
	durationQuery := `
		SELECT 
			duration_days,
			COUNT(*) as count
		FROM licenses
		GROUP BY duration_days
		ORDER BY count DESC
		LIMIT 5
	`

	// 获取过期预警
	expirationQuery := `
		SELECT 
			COUNT(*) as expiring_soon
		FROM licenses
		WHERE activation_status = 'activated'
		AND enable_status = 'enabled'
		AND expires_at BETWEEN NOW() AND DATE_ADD(NOW(), INTERVAL 7 DAY)
	`

	var analytics struct {
		ProductConversion []struct {
			ID                int     `json:"id"`
			Name              string  `json:"name"`
			TotalLicenses     int     `json:"total_licenses"`
			ActivatedLicenses int     `json:"activated_licenses"`
			Revenue           float64 `json:"revenue"`
			ConversionRate    float64 `json:"conversion_rate"`
		} `json:"product_conversion"`

		TimeTrend []struct {
			Date              string  `json:"date"`
			NewLicenses       int     `json:"new_licenses"`
			ActivatedLicenses int     `json:"activated_licenses"`
			DailyRevenue      float64 `json:"daily_revenue"`
		} `json:"time_trend"`

		DurationDistribution []struct {
			DurationDays int `json:"duration_days"`
			Count        int `json:"count"`
		} `json:"duration_distribution"`

		ExpirationWarning struct {
			ExpiringSoon int `json:"expiring_soon"`
		} `json:"expiration_warning"`
	}

	// 执行查询
	rows, err := model.DB.Query(productConversionQuery)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Error(http.StatusInternalServerError, "获取产品转化率失败"))
		return
	}
	defer rows.Close()

	for rows.Next() {
		var item struct {
			ID                int     `json:"id"`
			Name              string  `json:"name"`
			TotalLicenses     int     `json:"total_licenses"`
			ActivatedLicenses int     `json:"activated_licenses"`
			Revenue           float64 `json:"revenue"`
			ConversionRate    float64 `json:"conversion_rate"`
		}
		err := rows.Scan(&item.ID, &item.Name, &item.TotalLicenses, &item.ActivatedLicenses, &item.Revenue)
		if err != nil {
			continue
		}
		if item.TotalLicenses > 0 {
			item.ConversionRate = float64(item.ActivatedLicenses) / float64(item.TotalLicenses) * 100
		}
		analytics.ProductConversion = append(analytics.ProductConversion, item)
	}

	// 执行时间趋势查询
	rows, err = model.DB.Query(trendQuery)
	if err == nil {
		defer rows.Close()
		for rows.Next() {
			var item struct {
				Date              string  `json:"date"`
				NewLicenses       int     `json:"new_licenses"`
				ActivatedLicenses int     `json:"activated_licenses"`
				DailyRevenue      float64 `json:"daily_revenue"`
			}
			if err := rows.Scan(&item.Date, &item.NewLicenses, &item.ActivatedLicenses, &item.DailyRevenue); err == nil {
				analytics.TimeTrend = append(analytics.TimeTrend, item)
			}
		}
	}

	// 执行有效期分布查询
	rows, err = model.DB.Query(durationQuery)
	if err == nil {
		defer rows.Close()
		for rows.Next() {
			var item struct {
				DurationDays int `json:"duration_days"`
				Count        int `json:"count"`
			}
			if err := rows.Scan(&item.DurationDays, &item.Count); err == nil {
				analytics.DurationDistribution = append(analytics.DurationDistribution, item)
			}
		}
	}

	// 执行过期预警查询
	err = model.DB.QueryRow(expirationQuery).Scan(&analytics.ExpirationWarning.ExpiringSoon)
	if err != nil {
		analytics.ExpirationWarning.ExpiringSoon = 0
	}

	c.JSON(http.StatusOK, model.Success(analytics, "获取分析数据成功"))
}
