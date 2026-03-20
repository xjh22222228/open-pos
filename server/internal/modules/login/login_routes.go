package login

import "github.com/gin-gonic/gin"

func Routes(g *gin.RouterGroup) {
	g.POST("/login", SignController)
	g.POST("/register", RegisterController)
}
