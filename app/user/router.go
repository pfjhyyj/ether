package user

import (
	"github.com/gin-gonic/gin"
	"github.com/pfjhyyj/ether/app/user/controller"
	"github.com/pfjhyyj/ether/app/user/service"
	"github.com/pfjhyyj/ether/middleware"
)

func SetUserRouter(r *gin.RouterGroup) {
	userService := service.NewUserService()
	userController := controller.NewUserController(userService)

	router := r.Group("/users")
	router.Use(middleware.AuthMiddleware())
	{
		router.GET("", userController.ListUsers)
	}
}
