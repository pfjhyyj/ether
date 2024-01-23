package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pfjhyyj/ether/app/permission/define"
	"github.com/pfjhyyj/ether/app/permission/service"
	"github.com/pfjhyyj/ether/app/permission/utils"
	"github.com/pfjhyyj/ether/common"
	"net/http"
)

type PermissionController struct {
	service *service.PermissionService
}

func NewPermissionController(service *service.PermissionService) *PermissionController {
	return &PermissionController{}
}

func (c *PermissionController) CreatePermission(ctx *gin.Context) {
	var req define.CreatePermissionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		_ = ctx.Error(err)
		return
	}

	perm := utils.ConvertCreatePermissionRequestToPermission(&req)
	if err := c.service.CreatePermission(ctx, perm); err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, &common.Response{
		Code: common.Ok,
	})
}

func (c *PermissionController) UpdatePermission(ctx *gin.Context) {
	var req define.UpdatePermissionRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		_ = ctx.Error(err)
		return
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		_ = ctx.Error(err)
		return
	}

	perm := utils.ConvertUpdatePermissionRequestToPermission(&req)
	if err := c.service.UpdatePermission(ctx, perm); err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, &common.Response{
		Code: common.Ok,
	})
}

func (c *PermissionController) DeletePermission(ctx *gin.Context) {
	var req define.DeletePermissionRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		_ = ctx.Error(err)
		return
	}

	if err := c.service.DeletePermission(ctx, req.PermissionId); err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, &common.Response{
		Code: common.Ok,
	})
}

func (c *PermissionController) ListPermissions(ctx *gin.Context) {
	var req define.ListPermissionsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		_ = ctx.Error(err)
		return
	}

	queryParam := utils.ConvertListPermissionRequestToParam(&req)
	permissions, total, err := c.service.ListPermissions(ctx, queryParam)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	permsInfo := utils.ConvertPermissionListToPageResponse(permissions)
	ctx.JSON(http.StatusOK, &common.Response{
		Code: common.Ok,
		Data: &common.Page{
			Current: req.Current,
			Total:   total,
			List:    permsInfo,
		},
	})
}
