package main

import (
	"log"
	"os"
	"path/filepath"
	"program-verify/internal/config"
	"program-verify/internal/model"
	"program-verify/internal/router"
)

func main() {
	// 加载配置
	cfg := config.LoadConfig()

	// 确保数据目录存在
	if err := os.MkdirAll(filepath.Dir(cfg.DBPath), 0755); err != nil {
		log.Fatalf("创建数据目录失败: %v", err)
	}

	// 初始化数据库
	if err := model.InitDB(cfg.DBPath); err != nil {
		log.Fatalf("初始化数据库失败: %v", err)
	}

	// 设置路由
	r := router.SetupRouter()

	// 启动服务器
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("启动服务器失败: %v", err)
	}
}
