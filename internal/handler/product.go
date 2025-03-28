package handler

import (
	"net/http"
	"program-verify/internal/model"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// GetProducts 获取产品列表
func GetProducts(c *gin.Context) {
	// 获取筛选参数
	status := c.Query("status")
	name := c.Query("name")
	minPrice := c.Query("min_price")
	maxPrice := c.Query("max_price")
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	// 构建查询条件
	whereClause := "WHERE 1=1"
	args := []interface{}{}
	if status != "" {
		whereClause += " AND status = ?"
		args = append(args, status)
	}
	if name != "" {
		whereClause += " AND name LIKE ?"
		args = append(args, "%"+name+"%")
	}
	if minPrice != "" {
		if minPriceFloat, err := strconv.ParseFloat(minPrice, 64); err == nil {
			whereClause += " AND price >= ?"
			args = append(args, minPriceFloat)
		}
	}
	if maxPrice != "" {
		if maxPriceFloat, err := strconv.ParseFloat(maxPrice, 64); err == nil {
			whereClause += " AND price <= ?"
			args = append(args, maxPriceFloat)
		}
	}
	if startDate != "" {
		whereClause += " AND created_at >= ?"
		args = append(args, startDate)
	}
	if endDate != "" {
		whereClause += " AND created_at <= ?"
		args = append(args, endDate)
	}

	// 获取总数
	var total int
	countQuery := "SELECT COUNT(*) FROM products " + whereClause
	err := model.DB.QueryRow(countQuery, args...).Scan(&total)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Error(http.StatusInternalServerError, "获取产品总数失败"))
		return
	}

	// 检查是否有分页参数
	pageStr := c.Query("page")
	pageSizeStr := c.Query("page_size")

	// 如果没有分页参数，返回所有数据
	if pageStr == "" && pageSizeStr == "" {
		query := `
			SELECT id, name, price, created_at, status
			FROM products
			` + whereClause + `
			ORDER BY created_at DESC
		`

		rows, err := model.DB.Query(query, args...)
		if err != nil {
			c.JSON(http.StatusInternalServerError, model.Error(http.StatusInternalServerError, "获取产品列表失败"))
			return
		}
		defer rows.Close()

		var products []model.Product
		for rows.Next() {
			var p model.Product
			if err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.CreatedAt, &p.Status); err != nil {
				c.JSON(http.StatusInternalServerError, model.Error(http.StatusInternalServerError, "解析产品数据失败"))
				return
			}
			products = append(products, p)
		}

		c.JSON(http.StatusOK, model.Success(gin.H{
			"total": total,
			"items": products,
		}, "获取产品列表成功"))
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
		SELECT id, name, price, created_at, status
		FROM products
		` + whereClause + `
		ORDER BY created_at DESC
		LIMIT ? OFFSET ?
	`
	args = append(args, pageSize, offset)

	rows, err := model.DB.Query(query, args...)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Error(http.StatusInternalServerError, "获取产品列表失败"))
		return
	}
	defer rows.Close()

	var products []model.Product
	for rows.Next() {
		var p model.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.CreatedAt, &p.Status); err != nil {
			c.JSON(http.StatusInternalServerError, model.Error(http.StatusInternalServerError, "解析产品数据失败"))
			return
		}
		products = append(products, p)
	}

	c.JSON(http.StatusOK, model.Success(gin.H{
		"total":       total,
		"total_pages": totalPages,
		"page":        page,
		"page_size":   pageSize,
		"items":       products,
	}, "获取产品列表成功"))
}

// GetProduct 获取单个产品
func GetProduct(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Error(http.StatusBadRequest, "无效的产品ID"))
		return
	}

	var product model.Product
	err = model.DB.QueryRow(`
		SELECT id, name, price, created_at, status
		FROM products
		WHERE id = ?
	`, id).Scan(&product.ID, &product.Name, &product.Price, &product.CreatedAt, &product.Status)

	if err != nil {
		c.JSON(http.StatusNotFound, model.Error(http.StatusNotFound, "产品不存在"))
		return
	}

	c.JSON(http.StatusOK, model.Success(product, "获取产品成功"))
}

// CreateProduct 创建产品
func CreateProduct(c *gin.Context) {
	var req model.CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Error(http.StatusBadRequest, "请求参数错误"))
		return
	}

	now := time.Now()
	result, err := model.DB.Exec(
		"INSERT INTO products (name, price, status, created_at) VALUES (?, ?, ?, ?)",
		req.Name, req.Price, model.ProductStatusEnabled, now,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Error(http.StatusInternalServerError, "创建产品失败"))
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Error(http.StatusInternalServerError, "获取产品ID失败"))
		return
	}

	c.JSON(http.StatusCreated, model.Success(gin.H{
		"id":         id,
		"name":       req.Name,
		"price":      req.Price,
		"status":     model.ProductStatusEnabled,
		"created_at": now,
	}, "创建产品成功"))
}

// UpdateProduct 更新产品
func UpdateProduct(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Error(http.StatusBadRequest, "无效的产品ID"))
		return
	}

	var req model.UpdateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Error(http.StatusBadRequest, "请求参数错误"))
		return
	}

	// 检查产品是否存在
	var exists bool
	err = model.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM products WHERE id = ?)", id).Scan(&exists)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Error(http.StatusInternalServerError, "检查产品失败"))
		return
	}
	if !exists {
		c.JSON(http.StatusNotFound, model.Error(http.StatusNotFound, "产品不存在"))
		return
	}

	// 构建更新语句
	var updateFields []string
	var args []interface{}

	if req.Name != "" {
		updateFields = append(updateFields, "name = ?")
		args = append(args, req.Name)
	}
	if req.Price > 0 {
		updateFields = append(updateFields, "price = ?")
		args = append(args, req.Price)
	}
	if req.Status != "" {
		updateFields = append(updateFields, "status = ?")
		args = append(args, req.Status)
	}

	if len(updateFields) == 0 {
		c.JSON(http.StatusBadRequest, model.Error(http.StatusBadRequest, "没有要更新的字段"))
		return
	}

	// 添加ID到参数列表
	args = append(args, id)

	// 构建并执行更新语句
	query := "UPDATE products SET " + strings.Join(updateFields, ", ") + " WHERE id = ?"
	_, err = model.DB.Exec(query, args...)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Error(http.StatusInternalServerError, "更新产品失败"))
		return
	}

	// 获取更新后的产品信息
	var product model.Product
	err = model.DB.QueryRow(`
		SELECT id, name, price, created_at, status
		FROM products
		WHERE id = ?
	`, id).Scan(&product.ID, &product.Name, &product.Price, &product.CreatedAt, &product.Status)

	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Error(http.StatusInternalServerError, "获取更新后的产品信息失败"))
		return
	}

	// 根据更新的字段返回不同的提示信息
	var message string
	if req.Status != "" {
		if req.Status == model.ProductStatusEnabled {
			message = "启用产品成功"
		} else if req.Status == model.ProductStatusDisabled {
			message = "禁用产品成功"
		}
	} else {
		message = "更新产品成功"
	}

	c.JSON(http.StatusOK, model.Success(product, message))
}

// DisableProduct 禁用产品
func DisableProduct(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Error(http.StatusBadRequest, "无效的产品ID"))
		return
	}

	// 检查产品是否存在且未禁用
	var exists bool
	err = model.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM products WHERE id = ? AND status != ?)", id, model.ProductStatusDisabled).Scan(&exists)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Error(http.StatusInternalServerError, "检查产品失败"))
		return
	}
	if !exists {
		c.JSON(http.StatusNotFound, model.Error(http.StatusNotFound, "产品不存在或已被禁用"))
		return
	}

	// 执行禁用
	_, err = model.DB.Exec(`
		UPDATE products 
		SET status = ?
		WHERE id = ?
	`, model.ProductStatusDisabled, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Error(http.StatusInternalServerError, "禁用产品失败"))
		return
	}

	// 获取更新后的产品信息
	var product model.Product
	err = model.DB.QueryRow(`
		SELECT id, name, price, created_at, status
		FROM products
		WHERE id = ?
	`, id).Scan(&product.ID, &product.Name, &product.Price, &product.CreatedAt, &product.Status)

	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Error(http.StatusInternalServerError, "获取更新后的产品信息失败"))
		return
	}

	c.JSON(http.StatusOK, model.Success(product, "禁用产品成功"))
}

// EnableProduct 启用产品
func EnableProduct(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Error(http.StatusBadRequest, "无效的产品ID"))
		return
	}

	// 检查产品是否存在且已禁用
	var exists bool
	err = model.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM products WHERE id = ? AND status = ?)", id, model.ProductStatusDisabled).Scan(&exists)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Error(http.StatusInternalServerError, "检查产品失败"))
		return
	}
	if !exists {
		c.JSON(http.StatusNotFound, model.Error(http.StatusNotFound, "产品不存在或未被禁用"))
		return
	}

	// 执行启用
	_, err = model.DB.Exec(`
		UPDATE products 
		SET status = ?
		WHERE id = ?
	`, model.ProductStatusEnabled, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Error(http.StatusInternalServerError, "启用产品失败"))
		return
	}

	// 获取更新后的产品信息
	var product model.Product
	err = model.DB.QueryRow(`
		SELECT id, name, price, created_at, status
		FROM products
		WHERE id = ?
	`, id).Scan(&product.ID, &product.Name, &product.Price, &product.CreatedAt, &product.Status)

	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Error(http.StatusInternalServerError, "获取更新后的产品信息失败"))
		return
	}

	c.JSON(http.StatusOK, model.Success(product, "启用产品成功"))
}

// GetProductStats 获取产品统计信息
func GetProductStats(c *gin.Context) {
	// 获取产品统计信息
	query := `
		SELECT 
			COUNT(*) as total_products,
			COUNT(CASE WHEN status = 'enabled' THEN 1 END) as enabled_products,
			COUNT(CASE WHEN status = 'disabled' THEN 1 END) as disabled_products,
			COALESCE(SUM(price), 0) as total_price,
			COALESCE(AVG(CASE WHEN status = 'enabled' THEN price END), 0) as avg_enabled_price
		FROM products
	`

	var stats struct {
		TotalProducts    int     `json:"total_products"`    // 产品总数
		EnabledProducts  int     `json:"enabled_products"`  // 已启用产品数
		DisabledProducts int     `json:"disabled_products"` // 已禁用产品数
		TotalPrice       float64 `json:"total_price"`       // 所有产品总价
		AvgEnabledPrice  float64 `json:"avg_enabled_price"` // 已启用产品平均价格
	}

	err := model.DB.QueryRow(query).Scan(
		&stats.TotalProducts,
		&stats.EnabledProducts,
		&stats.DisabledProducts,
		&stats.TotalPrice,
		&stats.AvgEnabledPrice,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Error(http.StatusInternalServerError, "获取统计信息失败"))
		return
	}

	c.JSON(http.StatusOK, model.Success(stats, "获取统计信息成功"))
}
