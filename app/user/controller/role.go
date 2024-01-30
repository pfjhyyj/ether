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

type RoleController struct {
	service *service.RoleService
}

func NewRoleController(service *service.RoleService) *RoleController {
	return &RoleController{
		service: service,
	}
}

// CreateRole godoc
// @Summary Create role
// @Description Create role
// @Tags role
// @Accept json
// @Produce json
// @Security Bearer
// @Param request body define.CreateRoleRequest true "CreateRoleRequest"
// @Success 200 {object} common.Response
// @Router /roles [post]
func (c *RoleController) CreateRole(ctx *gin.Context) {
	if ok := utils2.CheckPermission(ctx, "role", "create"); !ok {
		ctx.JSON(http.StatusForbidden, &common.Response{
			Code: common.NoPermissionError,
			Msg:  "no permission",
		})
		return
	}

	var req define.CreateRoleRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		_ = ctx.Error(err)
		return
	}

	role := utils.ConvertCreateRoleRequestToRole(&req)
	if err := c.service.CreateRole(ctx, role); err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, &common.Response{
		Code: common.Ok,
	})
}

// UpdateRole godoc
// @Summary Update role
// @Description Update role
// @Tags role
// @Accept json
// @Produce json
// @Security Bearer
// @Param request body define.UpdateRoleRequest true "UpdateRoleRequest"
// @Param role_id path int true "role_id"
// @Success 200 {object} common.Response
// @Router /roles/{role_id} [put]
func (c *RoleController) UpdateRole(ctx *gin.Context) {
	if ok := utils2.CheckPermission(ctx, "role", "update"); !ok {
		ctx.JSON(http.StatusForbidden, &common.Response{
			Code: common.NoPermissionError,
			Msg:  "no permission",
		})
		return
	}

	var req define.UpdateRoleRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		_ = ctx.Error(err)
		return
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		_ = ctx.Error(err)
		return
	}

	role := utils.ConvertUpdateRoleRequestToRole(&req)
	if err := c.service.UpdateRole(ctx, role); err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, &common.Response{
		Code: common.Ok,
	})
}

// DeleteRole godoc
// @Summary Delete role
// @Description Delete role
// @Tags role
// @Accept json
// @Produce json
// @Security Bearer
// @Param role_id path int true "role_id"
// @Success 200 {object} common.Response
// @Router /roles/{role_id} [delete]
func (c *RoleController) DeleteRole(ctx *gin.Context) {
	if ok := utils2.CheckPermission(ctx, "role", "delete"); !ok {
		ctx.JSON(http.StatusForbidden, &common.Response{
			Code: common.NoPermissionError,
			Msg:  "no permission",
		})
		return
	}

	var req define.DeleteRoleRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		_ = ctx.Error(err)
		return
	}

	if err := c.service.DeleteRole(ctx, req.RoleId); err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, &common.Response{
		Code: common.Ok,
	})
}

// ListRoles godoc
// @Summary List roles
// @Description List roles
// @Tags role
// @Accept json
// @Produce json
// @Security Bearer
// @Param page query int false "page"
// @Param page_size query int false "page_size"
// @Success 200 {object} common.Response{data=common.Page{list=[]define.RolePageResponse}}
// @Router /roles [get]
func (c *RoleController) ListRoles(ctx *gin.Context) {
	if ok := utils2.CheckPermission(ctx, "role", "list"); !ok {
		ctx.JSON(http.StatusForbidden, &common.Response{
			Code: common.NoPermissionError,
			Msg:  "no permission",
		})
		return
	}

	var req define.ListRoleRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		_ = ctx.Error(err)
		return
	}

	queryParam := utils.ConvertRoleListPageRequestToParam(&req)

	roles, total, err := c.service.ListRoles(ctx, queryParam)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	roleInfo := utils.ConvertRoleListToPageResponse(roles)

	ctx.JSON(http.StatusOK, &common.Response{
		Code: common.Ok,
		Data: &common.Page{
			Total:    total,
			PageSize: req.PageSize,
			Current:  req.Current,
			List:     roleInfo,
		},
	})
}
