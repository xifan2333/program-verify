package model

import (
	"time"
)

// LicenseActivationStatus 许可证激活状态
type LicenseActivationStatus string

const (
	LicenseActivationStatusInactive  LicenseActivationStatus = "inactive"  // 未激活
	LicenseActivationStatusActivated LicenseActivationStatus = "activated" // 已激活
	LicenseActivationStatusExpired   LicenseActivationStatus = "expired"   // 已过期
)

// LicenseEnableStatus 许可证启用状态
type LicenseEnableStatus string

const (
	LicenseEnableStatusEnabled  LicenseEnableStatus = "enabled"  // 启用
	LicenseEnableStatusDisabled LicenseEnableStatus = "disabled" // 禁用
)

// License 许可证
type License struct {
	ID               int64                   `json:"id" db:"id"`
	ProductID        int64                   `json:"product_id" db:"product_id"`
	Key              string                  `json:"license_key" db:"license_key"`
	DurationDays     int                     `json:"duration_days" db:"duration_days"`
	CreatedAt        time.Time               `json:"created_at" db:"created_at"`
	ActivatedAt      *time.Time              `json:"activated_at,omitempty" db:"activated_at"`
	ExpiresAt        *time.Time              `json:"expires_at,omitempty" db:"expires_at"`
	ActivationStatus LicenseActivationStatus `json:"activation_status" db:"activation_status"`
	EnableStatus     LicenseEnableStatus     `json:"enable_status" db:"enable_status"`
	Remark           string                  `json:"remark,omitempty" db:"remark"`
}

// LicenseWithProduct 包含产品信息的许可证
type LicenseWithProduct struct {
	ID               int64                   `json:"id"`
	Key              string                  `json:"license_key"`
	DurationDays     int                     `json:"duration_days"`
	CreatedAt        time.Time               `json:"created_at"`
	ActivatedAt      *time.Time              `json:"activated_at,omitempty"`
	ExpiresAt        *time.Time              `json:"expires_at,omitempty"`
	ActivationStatus LicenseActivationStatus `json:"activation_status"`
	EnableStatus     LicenseEnableStatus     `json:"enable_status"`
	Remark           string                  `json:"remark,omitempty"`
	Product          Product                 `json:"product"`
}

// LicenseInfo 许可证信息
type LicenseInfo struct {
	ProductName string    `json:"product_name"`
	ActivatedAt time.Time `json:"activated_at"`
	ExpiresAt   time.Time `json:"expires_at"`
}

// GenerateLicenseRequest 生成许可证请求
type GenerateLicenseRequest struct {
	ProductID    int64  `json:"product_id" binding:"required"`
	DurationDays int    `json:"duration_days" binding:"required,min=1"`
	Remark       string `json:"remark,omitempty"`
}

// BatchGenerateLicenseRequest 批量生成许可证请求
type BatchGenerateLicenseRequest struct {
	ProductID    int64  `json:"product_id" binding:"required"`
	Count        int    `json:"count" binding:"required,min=1,max=100"`
	DurationDays int    `json:"duration_days" binding:"required,min=1"`
	Remark       string `json:"remark,omitempty"`
}

type ActivateLicenseRequest struct {
	LicenseKey string `json:"license_key" binding:"required"`
}

// VerifyLicenseRequest 验证许可证请求
type VerifyLicenseRequest struct {
	LicenseKey string `json:"license_key" binding:"required"`
}

// VerifyLicenseResponse 验证许可证响应
type VerifyLicenseResponse struct {
	Valid   bool        `json:"valid"`
	License LicenseInfo `json:"license,omitempty"`
	Message string      `json:"message,omitempty"`
}

// CreateLicenseRequest 创建许可证请求
type CreateLicenseRequest struct {
	ProductID    int64  `json:"product_id" binding:"required"`
	DurationDays int    `json:"duration_days" binding:"required,min=1"`
	Count        int    `json:"count" binding:"required,min=1,max=100"`
	Remark       string `json:"remark"`
}

// UpdateLicenseRequest 更新许可证请求
type UpdateLicenseRequest struct {
	EnableStatus string     `json:"enable_status,omitempty" binding:"omitempty,oneof=enabled disabled"`
	Remark       string     `json:"remark,omitempty"`
	ExpiresAt    *time.Time `json:"expires_at,omitempty"`
}
