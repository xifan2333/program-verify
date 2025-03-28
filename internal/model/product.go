package model

import (
	"time"
)

// ProductStatus 产品状态
type ProductStatus string

const (
	ProductStatusEnabled  ProductStatus = "enabled"  // 启用
	ProductStatusDisabled ProductStatus = "disabled" // 禁用
)

// Product 产品模型
type Product struct {
	ID        int64         `json:"id" db:"id"`
	Name      string        `json:"name" db:"name"`
	Price     float64       `json:"price" db:"price"`
	CreatedAt time.Time     `json:"created_at" db:"created_at"`
	Status    ProductStatus `json:"status" db:"status"`
}

// CreateProductRequest 创建产品请求
type CreateProductRequest struct {
	Name  string  `json:"name" binding:"required"`
	Price float64 `json:"price" binding:"required,gt=0"`
}

// UpdateProductRequest 更新产品请求
type UpdateProductRequest struct {
	Name   string        `json:"name,omitempty"`
	Price  float64       `json:"price,omitempty"`
	Status ProductStatus `json:"status,omitempty"`
}

// ProductWithStats 包含产品信息和许可证统计的结构体
type ProductWithStats struct {
	ID               int64         `json:"id" db:"id"`
	Name             string        `json:"name" db:"name"`
	Price            float64       `json:"price" db:"price"`
	CreatedAt        time.Time     `json:"created_at" db:"created_at"`
	Status           ProductStatus `json:"status" db:"status"`
	TotalLicenses    int           `json:"total_licenses" db:"total_licenses"`
	ActiveLicenses   int           `json:"active_licenses" db:"active_licenses"`
	InactiveLicenses int           `json:"inactive_licenses" db:"inactive_licenses"`
}
