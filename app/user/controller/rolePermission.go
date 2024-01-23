package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pfjhyyj/ether/app/user/define"
	"github.com/pfjhyyj/ether/app/user/service"
	"github.com/pfjhyyj/ether/common"
	"net/http"
)

type RolePermissionController struct {
	service *service.RolePermissionService
}

func NewRolePermissionController(service *service.RolePermissionService) *RolePermissionController {
	return &RolePermissionController{
		service: service,
	}
}

func (c *RolePermissionController) AddRolePermission(ctx *gin.Context) {
	var req define.AddRolePermissionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		_ = ctx.Error(err)
		return
	}

	if err := c.service.AddRolePermission(ctx, &req); err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, &common.Response{
		Code: common.Ok,
	})
}

func (c *RolePermissionController) DeleteRolePermission(ctx *gin.Context) {
	var req define.DeleteRolePermissionRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		_ = ctx.Error(err)
		return
	}

	if err := c.service.DeleteRolePermission(ctx, &req); err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, &common.Response{
		Code: common.Ok,
	})
}

func (c *RolePermissionController) ListRolePermission(ctx *gin.Context) {
	var req define.ListRolePermissionRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		_ = ctx.Error(err)
		return
	}

	rolePermissions, err := c.service.ListPermissionIdsByRoleId(ctx, req.RoleId)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, &common.Response{
		Code: common.Ok,
		Data: rolePermissions,
	})
}
