package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pfjhyyj/ether/app/user/define"
	"github.com/pfjhyyj/ether/app/user/service"
	"github.com/pfjhyyj/ether/app/user/utils"
	"github.com/pfjhyyj/ether/common"
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

func (c *RoleController) CreateRole(ctx *gin.Context) {
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

func (c *RoleController) UpdateRole(ctx *gin.Context) {
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

func (c *RoleController) DeleteRole(ctx *gin.Context) {
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

func (c *RoleController) ListRoles(ctx *gin.Context) {
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
