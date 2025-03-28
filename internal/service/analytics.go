package service

import (
	"database/sql"
	"time"
)

// RevenueTrendData 收益趋势数据
type RevenueTrendData struct {
	Dates           []string  `json:"dates"`
	Revenue         []float64 `json:"revenue"`
	TotalRevenue    float64   `json:"total_revenue"`
	AvgDailyRevenue float64   `json:"avg_daily_revenue"`
}

// ProductActivationData 产品激活数据
type ProductActivationData struct {
	ID                uint    `json:"id"`
	Name              string  `json:"name"`
	Price             float64 `json:"price"`
	TotalLicenses     int     `json:"total_licenses"`
	ActivatedLicenses int     `json:"activated_licenses"`
	Revenue           float64 `json:"revenue"`
	ActivationRate    float64 `json:"activation_rate"`
}

// ProductActivationSummary 产品激活汇总数据
type ProductActivationSummary struct {
	Products []ProductActivationData `json:"products"`
	Summary  struct {
		TotalLicenses     int     `json:"total_licenses"`
		TotalActivated    int     `json:"total_activated"`
		TotalRevenue      float64 `json:"total_revenue"`
		AvgActivationRate float64 `json:"avg_activation_rate"`
	} `json:"summary"`
}

// StatsSummary 统计数据汇总
type StatsSummary struct {
	Products struct {
		Total    int `json:"total"`
		Enabled  int `json:"enabled"`
		Disabled int `json:"disabled"`
	} `json:"products"`
	Licenses struct {
		Total     int `json:"total"`
		Activated int `json:"activated"`
		Inactive  int `json:"inactive"`
		Expired   int `json:"expired"`
		Disabled  int `json:"disabled"`
	} `json:"licenses"`
	Revenue struct {
		Total     float64 `json:"total"`
		Today     float64 `json:"today"`
		ThisMonth float64 `json:"this_month"`
	} `json:"revenue"`
}

// CalculateRevenueTrend 计算收益趋势数据
func CalculateRevenueTrend(db *sql.DB, startDate, endDate time.Time, interval string) (*RevenueTrendData, error) {
	var query string
	var dateFormat string
	var groupBy string

	switch interval {
	case "day":
		dateFormat = "%Y-%m-%d"
		groupBy = "strftime('%Y-%m-%d', l.created_at)"
	case "week":
		dateFormat = "%Y-W%W"
		groupBy = "strftime('%Y-W%W', l.created_at)"
	case "month":
		dateFormat = "%Y-%m"
		groupBy = "strftime('%Y-%m', l.created_at)"
	default:
		dateFormat = "%Y-%m-%d"
		groupBy = "strftime('%Y-%m-%d', l.created_at)"
	}

	query = `
		SELECT 
			strftime(?, l.created_at) as date,
			COALESCE(SUM(p.price), 0) as daily_revenue
		FROM licenses l
		JOIN products p ON l.product_id = p.id
		WHERE datetime(l.created_at) BETWEEN datetime(?) AND datetime(?)
		AND l.activation_status IN ('activated', 'expired')
		AND l.enable_status = 'enabled'
		GROUP BY ` + groupBy + `
		ORDER BY date ASC
	`

	rows, err := db.Query(query, dateFormat, startDate.Format("2006-01-02 15:04:05"), endDate.Format("2006-01-02 15:04:05"))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// 初始化切片
	dates := make([]string, 0)
	revenue := make([]float64, 0)
	var totalRevenue float64
	var count int

	for rows.Next() {
		var date string
		var dailyRevenue float64
		if err := rows.Scan(&date, &dailyRevenue); err != nil {
			return nil, err
		}
		dates = append(dates, date)
		revenue = append(revenue, dailyRevenue)
		totalRevenue += dailyRevenue
		count++
	}

	avgDailyRevenue := 0.0
	if count > 0 {
		avgDailyRevenue = totalRevenue / float64(count)
	}

	return &RevenueTrendData{
		Dates:           dates,
		Revenue:         revenue,
		TotalRevenue:    totalRevenue,
		AvgDailyRevenue: avgDailyRevenue,
	}, nil
}

// CalculateProductActivation 计算产品激活数据
func CalculateProductActivation(db *sql.DB, startDate, endDate time.Time) (*ProductActivationSummary, error) {
	var summary ProductActivationSummary
	// 初始化切片
	summary.Products = make([]ProductActivationData, 0)

	// 构建查询
	query := `
		SELECT 
			p.id,
			p.name,
			p.price,
			COUNT(CASE WHEN datetime(l.created_at) BETWEEN datetime(?) AND datetime(?) THEN l.id ELSE NULL END) as total_licenses,
			SUM(CASE WHEN l.activation_status = 'activated' AND datetime(l.created_at) BETWEEN datetime(?) AND datetime(?) THEN 1 ELSE 0 END) as activated_licenses,
			COALESCE(SUM(CASE WHEN l.activation_status IN ('activated', 'expired') AND datetime(l.created_at) BETWEEN datetime(?) AND datetime(?) THEN p.price ELSE 0 END), 0) as revenue
		FROM products p
		LEFT JOIN licenses l ON p.id = l.product_id AND l.enable_status = 'enabled'
		WHERE p.status = 'enabled'
		GROUP BY p.id, p.name, p.price
		ORDER BY revenue DESC
	`

	rows, err := db.Query(query,
		startDate.Format("2006-01-02 15:04:05"), endDate.Format("2006-01-02 15:04:05"),
		startDate.Format("2006-01-02 15:04:05"), endDate.Format("2006-01-02 15:04:05"),
		startDate.Format("2006-01-02 15:04:05"), endDate.Format("2006-01-02 15:04:05"),
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var totalLicenses, totalActivated int
	var totalRevenue float64

	for rows.Next() {
		var data ProductActivationData
		if err := rows.Scan(
			&data.ID,
			&data.Name,
			&data.Price,
			&data.TotalLicenses,
			&data.ActivatedLicenses,
			&data.Revenue,
		); err != nil {
			return nil, err
		}

		// 计算激活率
		if data.TotalLicenses > 0 {
			data.ActivationRate = float64(data.ActivatedLicenses) / float64(data.TotalLicenses) * 100
		}

		summary.Products = append(summary.Products, data)
		totalLicenses += data.TotalLicenses
		totalActivated += data.ActivatedLicenses
		totalRevenue += data.Revenue
	}

	// 设置汇总数据
	summary.Summary.TotalLicenses = totalLicenses
	summary.Summary.TotalActivated = totalActivated
	summary.Summary.TotalRevenue = totalRevenue
	if totalLicenses > 0 {
		summary.Summary.AvgActivationRate = float64(totalActivated) / float64(totalLicenses) * 100
	}

	return &summary, nil
}

// CalculateStats 计算统计数据
func CalculateStats(db *sql.DB) (StatsSummary, error) {
	var stats StatsSummary

	// 计算产品统计
	productRows, err := db.Query(`
		SELECT status, COUNT(*) as count 
		FROM products 
		GROUP BY status
	`)
	if err != nil {
		return stats, err
	}
	defer productRows.Close()

	for productRows.Next() {
		var status string
		var count int
		if err := productRows.Scan(&status, &count); err != nil {
			return stats, err
		}
		stats.Products.Total += count
		if status == "enabled" {
			stats.Products.Enabled = count
		} else {
			stats.Products.Disabled = count
		}
	}

	// 计算许可证统计
	licenseRows, err := db.Query(`
		SELECT enable_status, activation_status, COUNT(*) as count 
		FROM licenses 
		GROUP BY enable_status, activation_status
	`)
	if err != nil {
		return stats, err
	}
	defer licenseRows.Close()

	for licenseRows.Next() {
		var enableStatus, activationStatus string
		var count int
		if err := licenseRows.Scan(&enableStatus, &activationStatus, &count); err != nil {
			return stats, err
		}
		stats.Licenses.Total += count
		if enableStatus == "disabled" {
			stats.Licenses.Disabled += count
		} else {
			switch activationStatus {
			case "activated":
				stats.Licenses.Activated += count
			case "inactive":
				stats.Licenses.Inactive += count
			case "expired":
				stats.Licenses.Expired += count
			}
		}
	}

	// 计算收益统计
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	monthStart := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())

	// 计算总收益 - 包括已激活和已过期的许可证
	revenueRows, err := db.Query(`
		SELECT 
			COALESCE(SUM(p.price), 0) as total_revenue,
			COALESCE(SUM(CASE WHEN l.created_at >= ? THEN p.price ELSE 0 END), 0) as today_revenue,
			COALESCE(SUM(CASE WHEN l.created_at >= ? THEN p.price ELSE 0 END), 0) as month_revenue
		FROM licenses l
		JOIN products p ON l.product_id = p.id
		WHERE l.activation_status IN ('activated', 'expired')
		AND l.enable_status = 'enabled'
	`, today, monthStart)
	if err != nil {
		return stats, err
	}
	defer revenueRows.Close()

	if revenueRows.Next() {
		if err := revenueRows.Scan(&stats.Revenue.Total, &stats.Revenue.Today, &stats.Revenue.ThisMonth); err != nil {
			return stats, err
		}
	}

	return stats, nil
}
