package database

import "gorm.io/gorm"

var DB *gorm.DB

// SetDB 设置数据库实例
func SetDB(db *gorm.DB) {
	DB = db
}

// GetDB 获取数据库实例
func GetDB() *gorm.DB {
	return DB
}
