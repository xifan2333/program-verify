package service

import (
	"net/http"
	"program-verify/internal/model"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

// AuthService 认证服务
type AuthService struct {
	jwtSecret string
}

// NewAuthService 创建认证服务实例
func NewAuthService(jwtSecret string) *AuthService {
	return &AuthService{
		jwtSecret: jwtSecret,
	}
}

// Login 处理登录请求
func (s *AuthService) Login(c *gin.Context) {
	var req model.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Error(http.StatusBadRequest, "请求参数错误"))
		return
	}

	// 获取用户
	user, err := model.GetUserByUsername(req.Username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, model.Error(http.StatusUnauthorized, "用户名或密码错误"))
		return
	}

	// 验证密码
	if !s.VerifyPassword(req.Password, user.PasswordHash) {
		c.JSON(http.StatusUnauthorized, model.Error(http.StatusUnauthorized, "用户名或密码错误"))
		return
	}

	// 生成JWT token
	tokenString, err := s.GenerateToken(user.ID, user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Error(http.StatusInternalServerError, "生成令牌失败"))
		return
	}

	c.JSON(http.StatusOK, model.Success(gin.H{
		"token": tokenString,
	}, "登录成功"))
}

// VerifyPassword 验证密码
func (s *AuthService) VerifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// HashPassword 对密码进行哈希处理
func (s *AuthService) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// GenerateToken 生成JWT令牌
func (s *AuthService) GenerateToken(userID int64, username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  userID,
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})
	return token.SignedString([]byte(s.jwtSecret))
}

// ValidateToken 验证JWT令牌
func (s *AuthService) ValidateToken(tokenString string) (*jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.jwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return &claims, nil
	}

	return nil, jwt.ErrSignatureInvalid
}

// AuthMiddleware JWT认证中间件
func (s *AuthService) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取Authorization头
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, model.Error(http.StatusUnauthorized, "未提供认证令牌"))
			c.Abort()
			return
		}

		// 检查Bearer前缀
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, model.Error(http.StatusUnauthorized, "无效的认证令牌格式"))
			c.Abort()
			return
		}

		// 验证token
		claims, err := s.ValidateToken(parts[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, model.Error(http.StatusUnauthorized, "无效的认证令牌"))
			c.Abort()
			return
		}

		// 将用户信息存储到上下文中
		c.Set("user_id", (*claims)["user_id"])
		c.Set("username", (*claims)["username"])

		c.Next()
	}
}
