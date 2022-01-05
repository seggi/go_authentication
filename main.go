package main

import (
	"fuko_backend/api/controller"
	"fuko_backend/api/repository"
	"fuko_backend/api/routes"
	"fuko_backend/api/service"
	"fuko_backend/infrastructure"
	"fuko_backend/models"
)

func init() {
	infrastructure.LoadEnv()
}

func main() {
	router := infrastructure.FukoGinRouter()
	db := infrastructure.FukoDatabase()

	postRepository := repository.FukoUserRepository(db)
	postService := service.FukoUserService(postRepository)
	postController := controller.FukoUserController(postService)
	postRoute := routes.FukoUserRoute(postController, router)
	postRoute.Setup()

	db.DB.AutoMigrate(&models.User{})

	router.Gin.Run(":8000")
}
