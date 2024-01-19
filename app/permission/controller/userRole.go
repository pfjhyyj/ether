package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pfjhyyj/ether/app/permission/service"
)

type UserRoleController struct {
	service *service.UserRoleService
}

func NewUserRoleController() *UserRoleController {
	return &UserRoleController{
		service: service.NewUserRoleService(),
	}
}

func (c *UserRoleController) CreateRole(ctx *gin.Context) {

}
