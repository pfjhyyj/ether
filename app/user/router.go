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

func SetMyRouter(r *gin.RouterGroup) {
	myService := service.NewMyService()
	myController := controller.NewMyController(myService)

	router := r.Group("/my")
	router.Use(middleware.AuthMiddleware())
	{
		router.GET("", myController.MyInfo)
		router.PUT("", myController.UpdateMyInfo)
		router.PUT("/password", myController.UpdateMyPassword)
	}
}

func SetRouter(r *gin.RouterGroup) {
	SetUserRouter(r)
	SetMyRouter(r)
}
