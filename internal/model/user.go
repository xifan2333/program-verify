package model

import (
	"database/sql"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// User 用户模型
type User struct {
	ID           int64     `json:"id" db:"id"`
	Username     string    `json:"username" db:"username"`
	PasswordHash string    `json:"-" db:"password_hash"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
}

// LoginRequest 登录请求
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginResponse 登录响应
type LoginResponse struct {
	Token    string `json:"token"`
	Username string `json:"username"`
}

// UpdateUserRequest 更新用户信息请求
type UpdateUserRequest struct {
	CurrentPassword string `json:"current_password" binding:"required"`
	NewUsername     string `json:"new_username" binding:"required"`
	NewPassword     string `json:"new_password" binding:"required"`
}

// GetUserByUsername 根据用户名获取用户
func GetUserByUsername(username string) (*User, error) {
	var user User
	err := DB.QueryRow(`
		SELECT id, username, password_hash, created_at
		FROM users
		WHERE username = ?
	`, username).Scan(&user.ID, &user.Username, &user.PasswordHash, &user.CreatedAt)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// UpdateUser 更新用户信息
func UpdateUser(userID int64, currentPassword, newUsername, newPassword string) error {
	// 首先获取用户信息
	var currentPasswordHash string
	err := DB.QueryRow("SELECT password_hash FROM users WHERE id = ?", userID).Scan(&currentPasswordHash)
	if err != nil {
		return err
	}

	// 验证当前密码
	if err := bcrypt.CompareHashAndPassword([]byte(currentPasswordHash), []byte(currentPassword)); err != nil {
		return err
	}

	// 检查新用户名是否已存在
	var count int
	err = DB.QueryRow("SELECT COUNT(*) FROM users WHERE username = ? AND id != ?", newUsername, userID).Scan(&count)
	if err != nil {
		return err
	}
	if count > 0 {
		return fmt.Errorf("username already exists")
	}

	// 生成新密码哈希
	newPasswordHash, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// 更新用户信息
	_, err = DB.Exec(
		"UPDATE users SET username = ?, password_hash = ? WHERE id = ?",
		newUsername, string(newPasswordHash), userID,
	)
	return err
}
