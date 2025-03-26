package models

import (
	"time"

	"gorm.io/gorm"
)

type License struct {
	gorm.Model
	LicenseKey  string    `gorm:"uniqueIndex;not null" json:"license_key"`
	Status      string    `gorm:"not null" json:"status"` // unused, used, expired
	ExpiredAt   time.Time `json:"expired_at"`
	UsedAt      time.Time `json:"used_at,omitempty"`
	MachineCode string    `json:"machine_code,omitempty"`
}
