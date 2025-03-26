package main

import (
	"log"
	"os"

	"license-verify/database"
	"license-verify/handlers"
	"license-verify/middleware"
	"license-verify/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func init() {
	// 加载环境变量
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// 初始化数据库
	db, err := gorm.Open(sqlite.Open("license.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// 自动迁移数据库表
	db.AutoMigrate(&models.Admin{}, &models.License{})

	// 将数据库实例传递给handlers
	database.SetDB(db)
}

func main() {
	r := gin.Default()

	// 设置跨域中间件
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// 路由组
	api := r.Group("/api")
	{
		// 管理员路由
		admin := api.Group("/admin")
		{
			admin.POST("/login", handlers.AdminLogin)
			admin.Use(middleware.AuthRequired())
			admin.POST("/licenses/batch", handlers.CreateLicenses)
			admin.GET("/licenses", handlers.GetLicenses)
			admin.DELETE("/licenses/:id", handlers.DeleteLicense)
		}

		// 验证路由
		verify := api.Group("/verify")
		{
			verify.POST("", handlers.VerifyLicense)
		}
	}

	// 启动服务器
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
