package login

import (
	"github.com/gin-gonic/gin"
	"github.com/xjh22222228/open-erp/server/internal/cryptoutils"
	"github.com/xjh22222228/open-erp/server/internal/httputils"
	"github.com/xjh22222228/open-erp/server/internal/modules/tenant"
)

type SignRequest struct {
	TenantCode string `json:"tenantCode" binding:"required"`
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
}

func SignController(c *gin.Context) {
	resp := httputils.NewResponse(c)

	var req SignRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		resp.BadRequest(err.Error())
		return
	}

	loginSrv := NewLoginService()
	u, err := loginSrv.Authenticate(req.TenantCode, req.Username, req.Password)
	if err != nil {
		resp.Error(401, err.Error())
		return
	}

	token := cryptoutils.GenerateToken()

	// 存入 Redis 会话
	if err := loginSrv.CreateSession(token, u); err != nil {
		resp.Error(500, "登录失败，无法创建会话")
		return
	}

	resp.Success(gin.H{
		"token": token,
		"user": gin.H{
			"id":       u.ID,
			"username": u.Username,
			"realName": u.RealName,
		},
	}, "登录成功")
}

func RegisterController(c *gin.Context) {
	resp := httputils.NewResponse(c)

	var req SignRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		resp.BadRequest(err.Error())
		return
	}

	tenantSrv := tenant.NewTenantService()
	if err := tenantSrv.Register(req.TenantCode, req.Username, req.Password); err != nil {
		resp.Error(500, err.Error())
		return
	}

	resp.Success(nil, "注册成功")
}
