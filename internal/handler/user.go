package handler

import (
	"net/http"
	"program-verify/internal/model"
	"program-verify/internal/service"

	"github.com/gin-gonic/gin"
)

// UpdateUserInfo 更新用户信息
func UpdateUserInfo(authService *service.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取当前用户ID
		userID, exists := c.Get("user_id")
		if !exists {
			c.JSON(http.StatusUnauthorized, model.Error(http.StatusUnauthorized, "未提供认证令牌"))
			return
		}

		// 将 float64 转换为 int64
		userIDInt64, ok := userID.(float64)
		if !ok {
			c.JSON(http.StatusInternalServerError, model.Error(http.StatusInternalServerError, "无效的用户ID类型"))
			return
		}

		var req model.UpdateUserRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, model.Error(http.StatusBadRequest, "请求参数错误"))
			return
		}

		// 更新用户信息
		err := model.UpdateUser(
			int64(userIDInt64),
			req.CurrentPassword,
			req.NewUsername,
			req.NewPassword,
		)

		if err != nil {
			if err.Error() == "username already exists" {
				c.JSON(http.StatusConflict, model.Error(http.StatusConflict, "用户名已存在"))
				return
			}
			c.JSON(http.StatusInternalServerError, model.Error(http.StatusInternalServerError, "更新用户信息失败"))
			return
		}

		// 生成新的token
		token, err := authService.GenerateToken(int64(userIDInt64), req.NewUsername)
		if err != nil {
			c.JSON(http.StatusInternalServerError, model.Error(http.StatusInternalServerError, "生成token失败"))
			return
		}

		c.JSON(http.StatusOK, model.Success(gin.H{
			"token":    token,
			"username": req.NewUsername,
		}, "用户信息更新成功"))
	}
}
