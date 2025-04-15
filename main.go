package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"program-verify/internal/config"
	"program-verify/internal/model"
	"program-verify/internal/router"
)

//go:embed frontend/dist
var buildFS embed.FS

//go:embed frontend/dist/index.html
var indexPage []byte

func main() {
	// 加载配置
	cfg := config.LoadConfig()

	// 确保数据目录存在
	if err := os.MkdirAll(filepath.Dir(cfg.DBPath), 0777); err != nil {
		log.Fatalf("创建数据目录失败: %v", err)
	}

	// 检查目录权限
	if err := checkDirectoryPermissions(filepath.Dir(cfg.DBPath)); err != nil {
		log.Fatalf("目录权限检查失败: %v", err)
	}

	// 初始化数据库
	if err := model.InitDB(cfg.DBPath); err != nil {
		log.Fatalf("初始化数据库失败: %v", err)
	}

	// 设置嵌入的文件系统
	fsys, err := fs.Sub(buildFS, "frontend/dist")
	if err != nil {
		log.Fatalf("设置文件系统失败: %v", err)
	}

	// 设置路由
	r := router.SetupRouter(fsys, indexPage)

	// 启动服务器
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("启动服务器失败: %v", err)
	}
}

func checkDirectoryPermissions(path string) error {
	// 检查目录是否存在
	if _, err := os.Stat(path); err != nil {
		return fmt.Errorf("目录不存在: %v", err)
	}

	// 检查目录权限
	if err := os.Chmod(path, 0777); err != nil {
		return fmt.Errorf("设置目录权限失败: %v", err)
	}

	return nil
}
