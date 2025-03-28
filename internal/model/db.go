package model

import (
	"database/sql"

	"golang.org/x/crypto/bcrypt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB(dbPath string) error {
	var err error
	DB, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		return err
	}

	// 创建表
	_, err = DB.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT UNIQUE NOT NULL,
			password_hash TEXT NOT NULL,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		);

		CREATE TABLE IF NOT EXISTS products (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			price DECIMAL(10,2) NOT NULL,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			status TEXT NOT NULL DEFAULT 'enabled'
		);

		CREATE TABLE IF NOT EXISTS licenses (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			product_id INTEGER,
			license_key TEXT UNIQUE NOT NULL,
			duration_days INTEGER NOT NULL,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			activated_at DATETIME,
			expires_at DATETIME,
			activation_status TEXT NOT NULL DEFAULT 'inactive',
			enable_status TEXT NOT NULL DEFAULT 'enabled',
			remark TEXT,
			FOREIGN KEY (product_id) REFERENCES products(id)
		);
	`)
	if err != nil {
		return err
	}

	// 检查是否存在管理员用户
	var count int
	err = DB.QueryRow("SELECT COUNT(*) FROM users").Scan(&count)
	if err != nil {
		return err
	}

	// 如果不存在管理员用户，创建默认管理员
	if count == 0 {
		// 生成密码哈希
		passwordHash, err := bcrypt.GenerateFromPassword([]byte("zhi583379"), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		// 插入默认管理员用户
		_, err = DB.Exec(
			"INSERT INTO users (username, password_hash) VALUES (?, ?)",
			"xifan", string(passwordHash),
		)
		if err != nil {
			return err
		}
	}

	return nil
}
