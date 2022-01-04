package main

import (
	"fuko_backend/infrastructure"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(context *gin.Context) {
		infrastructure.LoadEnv()
		infrastructure.FukoDatabase()
		context.JSON(http.StatusOK, gin.H{"data": "Fuko Banckend"})
	})

	router.Run(":8000")
}
