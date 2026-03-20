package server

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/xjh22222228/open-erp/server/config"
	"github.com/xjh22222228/open-erp/server/internal/middleware"
	"github.com/xjh22222228/open-erp/server/internal/modules/login"
)

func RouterRun() {
	r := gin.Default()

	// 基础 API V1 组
	v1 := r.Group("/api/v1")

	// 1. 公开接口 (无需登录)
	{
		login.Routes(v1)
	}

	// 2. 受保护接口 (需要认证)
	auth := v1.Group("/")
	auth.Use(middleware.AuthMiddleware())
	{
		// 此处后续可以添加如：
		// user.Routes(auth)
		// store.Routes(auth)
	}

	port := config.GlobalConfig.Server.Port
	err := r.Run(fmt.Sprintf("0.0.0.0:%d", port))
	if err != nil {
		log.Panicf("启动服务器错误: %v", err)
	}
}
