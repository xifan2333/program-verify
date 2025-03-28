package router

import (
	"program-verify/internal/config"
	"program-verify/internal/handler"
	"program-verify/internal/service"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// SetupRouter 设置路由
func SetupRouter() *gin.Engine {
	// 加载配置
	cfg := config.LoadConfig()

	// 创建认证服务
	authService := service.NewAuthService(cfg.JWTSecret)

	// 创建Gin引擎
	r := gin.Default()
	gin.SetMode(gin.DebugMode)

	// 配置CORS
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	config.MaxAge = 12 * time.Hour
	r.Use(cors.New(config))

	// 静态文件服务
	r.Static("/assets", "./static/assets")
	r.StaticFile("/favicon.ico", "./static/favicon.ico")

	// API v1 路由组
	v1 := r.Group("/api/v1")
	{
		// 公开路由
		v1.POST("/auth/login", authService.Login)
		v1.POST("/licenses/verify", handler.VerifyLicense)

		// 需要认证的路由
		authorized := v1.Group("")
		authorized.Use(authService.AuthMiddleware())
		{
			// 产品相关路由
			product := authorized.Group("/products")
			{
				product.GET("", handler.GetProducts)
				product.POST("", handler.CreateProduct)
				product.GET("/:id", handler.GetProduct)
				product.PUT("/:id", handler.UpdateProduct)
				product.GET("/stats", handler.GetProductStats)
			}

			// 许可证相关路由
			license := authorized.Group("/licenses")
			{
				license.GET("", handler.GetLicenses)
				license.POST("", handler.CreateLicense)
				license.GET("/:id", handler.GetLicense)
				license.PUT("/:id", handler.UpdateLicense)
				license.GET("/stats", handler.GetLicenseStats)
			}

			// 数据分析相关路由
			analytics := authorized.Group("/analytics")
			{
				analytics.GET("/revenue-trend", handler.GetRevenueTrend)
				analytics.GET("/product-activation", handler.GetProductActivation)
				analytics.GET("/stats", handler.GetStats)
			}
		}
	}

	// 处理前端路由
	r.NoRoute(func(c *gin.Context) {
		c.File("./static/index.html")
	})

	return r
}
