package category

import "github.com/gin-gonic/gin"

func Routes(g *gin.RouterGroup) {
	g.POST("/categories/tree", TreeController)
	g.POST("/categories/create", CreateController)
	g.POST("/categories/update", UpdateController)
	g.POST("/categories/delete", DeleteController)
}
