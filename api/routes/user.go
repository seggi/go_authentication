package routes

import (
	"fuko_backend/api/controller"
	"fuko_backend/infrastructure"
)

// Route for user module
type UserRoute struct {
	Handler    infrastructure.GinRouter
	Controller controller.UserController
}

// Initializes new instance of UserRoute
func FukoUserRoute(
	controller controller.UserController,
	handler infrastructure.GinRouter,
) UserRoute {
	return UserRoute{
		Handler:    handler,
		Controller: controller,
	}
}

// Setups user routes
func (u UserRoute) Setup() {
	user := u.Handler.Gin.Group("/auth")
	{
		user.POST("/register", u.Controller.CreateUser)
		user.POST("/login", u.Controller.Login)
	}
}
