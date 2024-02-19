package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pfjhyyj/ether/app/user/define"
	"github.com/pfjhyyj/ether/app/user/service"
	"github.com/pfjhyyj/ether/app/user/utils"
	"github.com/pfjhyyj/ether/common"
	utils2 "github.com/pfjhyyj/ether/utils"
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

// AddRolePermission godoc
// @Summary Add role permission
// @Description Add role permission
// @Tags role
// @Accept json
// @Produce json
// @Security Bearer
// @Param role_id path int true "role_id"
// @Param request body define.AddRolePermissionRequest true "AddRolePermissionRequest"
// @Success 200 {object} common.Response
// @Router /roles/{roleId}/permissions/add [post]
func (c *RolePermissionController) AddRolePermission(ctx *gin.Context) {
	if ok := utils2.CheckPermission(ctx, "role_permission", "create"); !ok {
		ctx.JSON(http.StatusOK, &common.Response{
			Code: common.NoPermissionError,
			Msg:  "no permission",
		})
		return
	}

	var roleIdReq define.RoleIdUri
	if err := ctx.ShouldBindUri(&roleIdReq); err != nil {
		_ = ctx.Error(err)
		return
	}

	var req define.AddRolePermissionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		_ = ctx.Error(err)
		return
	}
	req.RoleId = roleIdReq.RoleId

	if err := c.service.AddRolePermission(ctx, &req); err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, &common.Response{
		Code: common.Ok,
	})
}

// DeleteRolePermission godoc
// @Summary Delete role permission
// @Description Delete role permission
// @Tags role
// @Accept json
// @Produce json
// @Security Bearer
// @Param role_id path int true "role_id"
// @Success 200 {object} common.Response
// @Router /roles/{roleId}/permissions/delete [post]
func (c *RolePermissionController) DeleteRolePermission(ctx *gin.Context) {
	if ok := utils2.CheckPermission(ctx, "role_permission", "delete"); !ok {
		ctx.JSON(http.StatusOK, &common.Response{
			Code: common.NoPermissionError,
			Msg:  "no permission",
		})
		return
	}

	var roleIdReq define.RoleIdUri
	if err := ctx.ShouldBindUri(&roleIdReq); err != nil {
		_ = ctx.Error(err)
		return
	}

	var req define.DeleteRolePermissionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		_ = ctx.Error(err)
		return
	}
	req.RoleId = roleIdReq.RoleId

	if err := c.service.DeleteRolePermission(ctx, &req); err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, &common.Response{
		Code: common.Ok,
	})
}

// ListRolePermission godoc
// @Summary List role permission
// @Description List role permission
// @Tags role
// @Accept json
// @Produce json
// @Security Bearer
// @Param role_id path int true "role_id"
// @Success 200 {object} common.Response
// @Router /roles/{roleId}/permissions [get]
func (c *RolePermissionController) ListRolePermission(ctx *gin.Context) {
	if ok := utils2.CheckPermission(ctx, "role_permission", "list"); !ok {
		ctx.JSON(http.StatusOK, &common.Response{
			Code: common.NoPermissionError,
			Msg:  "no permission",
		})
		return
	}

	var roleIdReq define.RoleIdUri
	if err := ctx.ShouldBindUri(&roleIdReq); err != nil {
		_ = ctx.Error
		return
	}

	var req define.ListRolePermissionRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		_ = ctx.Error(err)
		return
	}
	req.RoleId = roleIdReq.RoleId

	rolePermissions, total, err := c.service.ListRolePermissions(ctx, &req)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	list := utils.ConvertRolePermissionListToResponse(rolePermissions)

	ctx.JSON(http.StatusOK, &common.Response{
		Code: common.Ok,
		Data: &common.Page{
			Current:  req.Current,
			PageSize: req.PageSize,
			Total:    total,
			List:     list,
		},
	})
}
