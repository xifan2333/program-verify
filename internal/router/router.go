package router

import (
	"io/fs"
	"net/http"
	"program-verify/internal/config"
	"program-verify/internal/handler"
	"program-verify/internal/service"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// SetupRouter 设置路由
func SetupRouter(fsys fs.FS, indexPage []byte) *gin.Engine {
	// 加载配置
	cfg := config.LoadConfig()

	// 创建认证服务
	authService := service.NewAuthService(cfg.JWTSecret)

	// 创建Gin引擎
	r := gin.Default()

	// 配置CORS
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	config.MaxAge = 12 * time.Hour
	r.Use(cors.New(config))

	// API v1 路由组
	v1 := r.Group("/api/v1")
	{
		// 公开路由
		v1.POST("/auth/login", authService.Login)
		v1.POST("/licenses/verify", handler.VerifyLicense)
		v1.GET("/auth/verify", authService.VerifyToken)

		// 需要认证的路由
		authorized := v1.Group("")
		authorized.Use(authService.AuthMiddleware())
		{
			// 用户相关路由
			user := authorized.Group("/user")
			{
				user.PUT("/update", handler.UpdateUserInfo(authService))
			}

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

	// 设置静态文件服务
	r.StaticFS("/static", http.FS(fsys))

	// 设置 NoRoute 处理器
	r.NoRoute(func(c *gin.Context) {
		// 如果是 API 请求，返回 404
		if strings.HasPrefix(c.Request.RequestURI, "/api") {
			c.Status(http.StatusNotFound)
			return
		}

		// 如果是静态资源请求，尝试从文件系统提供
		if strings.HasPrefix(c.Request.RequestURI, "/static") {
			c.FileFromFS(c.Request.RequestURI, http.FS(fsys))
			return
		}

		// 其他请求返回 index.html
		c.Header("Cache-Control", "no-cache")
		c.Data(http.StatusOK, "text/html; charset=utf-8", indexPage)
	})

	return r
}
