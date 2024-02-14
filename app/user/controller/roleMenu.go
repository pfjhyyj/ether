package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pfjhyyj/ether/app/user/define"
	"github.com/pfjhyyj/ether/app/user/service"
	"github.com/pfjhyyj/ether/common"
	utils2 "github.com/pfjhyyj/ether/utils"
	"net/http"
)

type RoleMenuController struct {
	service *service.RoleMenuService
}

func NewRoleMenuController(service *service.RoleMenuService) *RoleMenuController {
	return &RoleMenuController{
		service: service,
	}
}

// AddRoleMenu godoc
// @Summary Add role menu
// @Description Add role menu
// @Tags role
// @Accept json
// @Produce json
// @Security Bearer
// @Param role_id path int true "role_id"
// @Param request body define.AddRoleMenuRequest true "AddRoleMenuRequest"
// @Success 200 {object} common.Response
// @Router /roles/{roleId}/menus/add [post]
func (c *RoleMenuController) AddRoleMenu(ctx *gin.Context) {
	if ok := utils2.CheckPermission(ctx, "role_menu", "create"); !ok {
		ctx.JSON(http.StatusForbidden, &common.Response{
			Code: common.NoPermissionError,
			Msg:  "no permission",
		})
		return
	}

	var req define.AddRoleMenuRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		_ = ctx.Error(err)
		return
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		_ = ctx.Error(err)
		return
	}

	if err := c.service.AddRoleMenu(ctx, &req); err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, &common.Response{
		Code: common.Ok,
	})
}

// DeleteRoleMenu godoc
// @Summary Delete role menu
// @Description Delete role menu
// @Tags role
// @Accept json
// @Produce json
// @Security Bearer
// @Param role_id path int true "role_id"
// @Param request body define.DeleteRoleMenuRequest true "DeleteRoleMenuRequest"
// @Success 200 {object} common.Response
// @Router /roles/{roleId}/menus/delete [post]
func (c *RoleMenuController) DeleteRoleMenu(ctx *gin.Context) {
	if ok := utils2.CheckPermission(ctx, "role_menu", "delete"); !ok {
		ctx.JSON(http.StatusForbidden, &common.Response{
			Code: common.NoPermissionError,
			Msg:  "no permission",
		})
		return
	}

	var req define.DeleteRoleMenuRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		_ = ctx.Error(err)
		return
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		_ = ctx.Error(err)
		return
	}

	if err := c.service.DeleteRoleMenu(ctx, &req); err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, &common.Response{
		Code: common.Ok,
	})
}

// ListRoleMenu godoc
// @Summary List role menu
// @Description List role menu
// @Tags role
// @Accept json
// @Produce json
// @Security Bearer
// @Param role_id path int true "role_id"
// @Success 200 {object} common.Response{data=define.ListRoleMenuResponse}
// @Router /roles/{roleId}/menus [get]
func (c *RoleMenuController) ListRoleMenu(ctx *gin.Context) {
	if ok := utils2.CheckPermission(ctx, "role_menu", "read"); !ok {
		ctx.JSON(http.StatusForbidden, &common.Response{
			Code: common.NoPermissionError,
			Msg:  "no permission",
		})
		return
	}

	var req define.ListRoleMenuRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		_ = ctx.Error(err)
		return
	}

	roleMenus, err := c.service.ListMenuIdsByRoleId(ctx, req.RoleId)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, &common.Response{
		Code: common.Ok,
		Data: roleMenus,
	})
}
