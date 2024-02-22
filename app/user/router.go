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

	userRoleService := service.NewUserRoleService()
	userRoleController := controller.NewUserRoleController(userRoleService)

	router := r.Group("/users")
	router.Use(middleware.AuthMiddleware())
	{
		router.GET("", userController.ListUsers)
		router.GET("/:userId", userController.GetUser)
		router.GET("/:userId/roles", userRoleController.ListUserRole)
		router.POST("/:userId/roles/add", userRoleController.AddUserRole)
		router.POST("/:userId/roles/delete", userRoleController.DeleteUserRole)
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
		router.GET("/menus", myController.GetMyMenu)
	}
}

func SetRoleRouter(r *gin.RouterGroup) {
	roleService := service.NewRoleService()
	roleController := controller.NewRoleController(roleService)

	rolePermissionService := service.NewRolePermissionService()
	rolePermissionController := controller.NewRolePermissionController(rolePermissionService)

	roleMenuService := service.NewRoleMenuService()
	roleMenuController := controller.NewRoleMenuController(roleMenuService)

	roleRouter := r.Group("/roles")
	roleRouter.Use(middleware.AuthMiddleware())
	{
		roleRouter.POST("", roleController.CreateRole)
		roleRouter.GET("", roleController.ListRoles)
		roleRouter.GET("/:roleId", roleController.GetRole)
		roleRouter.PUT("/:roleId", roleController.UpdateRole)
		roleRouter.DELETE("/:roleId", roleController.DeleteRole)
		roleRouter.GET("/:roleId/permissions", rolePermissionController.ListRolePermission)
		roleRouter.POST("/:roleId/permissions/add", rolePermissionController.AddRolePermission)
		roleRouter.POST("/:roleId/permissions/delete", rolePermissionController.DeleteRolePermission)
		roleRouter.GET("/:roleId/menus", roleMenuController.ListRoleMenu)
		roleRouter.POST("/:roleId/menus/add", roleMenuController.AddRoleMenu)
		roleRouter.POST("/:roleId/menus/delete", roleMenuController.DeleteRoleMenu)
		roleRouter.POST("/:roleId/menus/set", roleMenuController.SetRoleMenu)
	}
}

func SetPermissionRouter(r *gin.RouterGroup) {
	permissionService := service.NewPermissionService()
	permissionController := controller.NewPermissionController(permissionService)

	permissionRouter := r.Group("/permissions")
	permissionRouter.Use(middleware.AuthMiddleware())
	{
		permissionRouter.POST("", permissionController.CreatePermission)
		permissionRouter.PUT("/:permissionId", permissionController.UpdatePermission)
		permissionRouter.DELETE("/:permissionId", permissionController.DeletePermission)
		permissionRouter.GET("", permissionController.ListPermissions)
	}
}

func SetMenuRouter(r *gin.RouterGroup) {
	menuService := service.NewMenuService()
	menuController := controller.NewMenuController(menuService)

	router := r.Group("/menus")
	router.Use(middleware.AuthMiddleware())
	{
		router.POST("", menuController.CreateMenu)
		router.GET("", menuController.ListMenus)
		router.PUT("/:menuId", menuController.UpdateMenu)
		router.DELETE("/:menuId", menuController.DeleteMenu)
	}
}

func SetRouter(r *gin.RouterGroup) {
	SetUserRouter(r)
	SetMyRouter(r)
	SetRoleRouter(r)
	SetPermissionRouter(r)
	SetMenuRouter(r)
}
