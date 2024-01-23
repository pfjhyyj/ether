package tenant

import (
	"github.com/gin-gonic/gin"
	controller2 "github.com/pfjhyyj/ether/app/tenant/controller"
	service2 "github.com/pfjhyyj/ether/app/tenant/service"
	"github.com/pfjhyyj/ether/middleware"
)

func SetTenantRouter(r *gin.RouterGroup) {
	tenantService := service2.NewTenantService()
	tenantController := controller2.NewTenantController(tenantService)

	router := r.Group("/tenants")
	router.Use(middleware.AuthMiddleware())
	{
		router.POST("", tenantController.CreateTenant)
		router.PUT("/:tenantId", tenantController.UpdateTenant)
		router.DELETE("/:tenantId", tenantController.DeleteTenant)
		router.GET("", tenantController.ListTenants)
	}
}

func SetRouter(r *gin.RouterGroup) {
	SetTenantRouter(r)
}
