package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/pfjhyyj/ether/app/auth/controller"
	"github.com/pfjhyyj/ether/app/auth/service"
	"github.com/pfjhyyj/ether/app/user/domain"
	"github.com/pfjhyyj/ether/middleware"
)

func setLoginRouter(r *gin.RouterGroup) {
	userRepo := domain.NewUserRepository()
	loginService := service.NewLoginService(userRepo)
	loginController := controller.NewLoginController(loginService)

	r.POST("/loginByUsername", loginController.LoginByUsername)
}

func setLogoutRouter(r *gin.RouterGroup) {
	logoutService := service.NewLogoutService()
	logoutController := controller.NewLogoutController(logoutService)

	r.POST("/logout", middleware.AuthMiddleware(), logoutController.Logout)
}

func setRegisterRouter(r *gin.RouterGroup) {
	userRepo := domain.NewUserRepository()
	registerService := service.NewRegisterService(userRepo)
	registerController := controller.NewRegisterController(registerService)

	r.POST("/registerByEmail", registerController.RegisterByEmail)
}

func SetAuthRouter(r *gin.RouterGroup) {
	authRouter := r.Group("/auth")
	setLoginRouter(authRouter)
	setRegisterRouter(authRouter)
	setLogoutRouter(authRouter)
}
