package main

import (
	"log"
	"os"

	"license-verify/database"
	"license-verify/models"
	"license-verify/utils"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// 加载环境变量
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// 连接数据库
	db, err := gorm.Open(sqlite.Open("license.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// 设置数据库实例
	database.SetDB(db)

	// 创建管理员表
	db.AutoMigrate(&models.Admin{})

	// 检查是否已存在管理员
	var count int64
	db.Model(&models.Admin{}).Count(&count)
	if count > 0 {
		log.Println("Admin already exists")
		return
	}

	// 创建管理员账户
	hashedPassword, err := utils.HashPassword(os.Getenv("ADMIN_PASSWORD"))
	if err != nil {
		log.Fatal("Failed to hash password:", err)
	}

	admin := models.Admin{
		Username:     os.Getenv("ADMIN_USERNAME"),
		PasswordHash: hashedPassword,
	}

	if err := db.Create(&admin).Error; err != nil {
		log.Fatal("Failed to create admin:", err)
	}

	log.Println("Admin created successfully")
}
