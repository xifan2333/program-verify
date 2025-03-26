package models

import (
	"time"

	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	Username     string    `gorm:"uniqueIndex;not null" json:"username"`
	PasswordHash string    `gorm:"not null" json:"-"`
	CreatedAt    time.Time `json:"created_at"`
}
