package permission

import (
	"github.com/gin-gonic/gin"
	"github.com/pfjhyyj/ether/app/permission/controller"
	"github.com/pfjhyyj/ether/app/permission/service"
	"github.com/pfjhyyj/ether/middleware"
)

func SetRoleRouter(r *gin.RouterGroup) {
	roleService := service.NewRoleService()
	roleController := controller.NewRoleController(roleService)

	roleRouter := r.Group("/roles")
	roleRouter.Use(middleware.AuthMiddleware())
	{
		roleRouter.POST("", roleController.CreateRole)
		roleRouter.PUT("/:roleId", roleController.UpdateRole)
		roleRouter.DELETE("/:roleId", roleController.DeleteRole)
		roleRouter.GET("", roleController.ListRoles)
	}
}

func SetUserRoleRouter(r *gin.RouterGroup) {
	userRoleService := service.NewUserRoleService()
	userRoleController := controller.NewUserRoleController(userRoleService)

	userRoleRouter := r.Group("/userRoles")
	userRoleRouter.Use(middleware.AuthMiddleware())
	{
		userRoleRouter.POST("", userRoleController.AddUserRole)
		userRoleRouter.DELETE("", userRoleController.DeleteUserRole)
		userRoleRouter.GET("", userRoleController.ListUserRole)
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

func SetRolePermissionRouter(r *gin.RouterGroup) {
	rolePermissionService := service.NewRolePermissionService()
	rolePermissionController := controller.NewRolePermissionController(rolePermissionService)

	rolePermissionRouter := r.Group("/rolePermissions")
	rolePermissionRouter.Use(middleware.AuthMiddleware())
	{
		rolePermissionRouter.POST("", rolePermissionController.AddRolePermission)
		rolePermissionRouter.DELETE("", rolePermissionController.DeleteRolePermission)
		rolePermissionRouter.GET("", rolePermissionController.ListRolePermission)
	}
}

func SetRouter(r *gin.RouterGroup) {
	SetRoleRouter(r)
	SetUserRoleRouter(r)
	SetPermissionRouter(r)
	SetRolePermissionRouter(r)
}
