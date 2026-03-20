package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/xjh22222228/open-erp/server/database"
	"github.com/xjh22222228/open-erp/server/internal/httputils"
	"github.com/xjh22222228/open-erp/server/internal/models"
)

// AuthMiddleware 校验 Token 并注入用户信息到上下文
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		resp := httputils.NewResponse(c)

		// 1. 获取请求头 Authorization: Bearer <token>
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			resp.Error(http.StatusUnauthorized, "未登录或登录已过期")
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			resp.Error(http.StatusUnauthorized, "Authorization 格式错误")
			c.Abort()
			return
		}

		token := parts[1]

		// 2. 从 Redis 校验 Token
		key := fmt.Sprintf("session:%s", token)
		userData, err := database.RedisClient.Get(context.Background(), key).Result()
		if err != nil {
			resp.Error(http.StatusUnauthorized, "登录会话已过期，请重新登录")
			c.Abort()
			return
		}

		// 3. 反序列化并注入 Context
		var u models.ErpUser
		if err := json.Unmarshal([]byte(userData), &u); err != nil {
			resp.Error(http.StatusInternalServerError, "系统解析会话失败")
			c.Abort()
			return
		}

		c.Set("currentUser", &u)
		c.Next()
	}
}

// GetCurrentUser 从 Gin 上下文获取当前用户信息
func GetCurrentUser(c *gin.Context) *models.ErpUser {
	u, exists := c.Get("currentUser")
	if !exists {
		return nil
	}
	return u.(*models.ErpUser)
}
