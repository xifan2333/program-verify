package config

import (
	"os"
)

type Config struct {
	// 服务器配置
	Port   string
	DBPath string

	// JWT配置
	JWTSecret string

	// 管理员配置
	AdminUsername string
	AdminPassword string
}

func LoadConfig() *Config {
	cfg := &Config{}

	// 从环境变量加载配置
	cfg.Port = getEnvOrDefault("PORT", "8080")
	cfg.DBPath = getEnvOrDefault("DB_PATH", "data/program_verify.db")
	cfg.JWTSecret = getEnvOrDefault("JWT_SECRET", "your-secret-key")
	cfg.AdminUsername = getEnvOrDefault("ADMIN_USERNAME", "admin")
	cfg.AdminPassword = getEnvOrDefault("ADMIN_PASSWORD", "password")

	return cfg
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
