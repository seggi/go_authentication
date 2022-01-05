package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Gin Router
type GinRouter struct {
	Gin *gin.Engine
}

func FukoGinRouter() GinRouter {
	httpRouter := gin.Default()
	httpRouter.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Up and Running..."})
	})
	return GinRouter{
		Gin: httpRouter,
	}
}
