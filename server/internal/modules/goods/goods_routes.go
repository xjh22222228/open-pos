package goods

import "github.com/gin-gonic/gin"

func Routes(g *gin.RouterGroup) {
	g.POST("/goods/create", CreateController)
	g.POST("/goods/update", UpdateController)
	g.POST("/goods/delete", DeleteController)
	g.POST("/goods/list", ListController)
}
