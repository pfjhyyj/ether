package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pfjhyyj/ether/app/user/define"
	"github.com/pfjhyyj/ether/app/user/service"
	"github.com/pfjhyyj/ether/common"
	utils2 "github.com/pfjhyyj/ether/utils"
	"net/http"
)

type UserRoleController struct {
	service *service.UserRoleService
}

func NewUserRoleController(service *service.UserRoleService) *UserRoleController {
	return &UserRoleController{
		service: service,
	}
}

func (c *UserRoleController) AddUserRole(ctx *gin.Context) {
	if ok := utils2.CheckPermission(ctx, "user_role", "create"); !ok {
		ctx.JSON(http.StatusForbidden, &common.Response{
			Code: common.NoPermissionError,
			Msg:  "no permission",
		})
		return
	}

	var req define.AddUserRoleRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		_ = ctx.Error(err)
		return
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		_ = ctx.Error(err)
		return
	}

	if err := c.service.AddUserRole(ctx, &req); err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, &common.Response{
		Code: common.Ok,
	})
}

func (c *UserRoleController) DeleteUserRole(ctx *gin.Context) {
	if ok := utils2.CheckPermission(ctx, "user_role", "delete"); !ok {
		ctx.JSON(http.StatusForbidden, &common.Response{
			Code: common.NoPermissionError,
			Msg:  "no permission",
		})
		return
	}

	var req define.DeleteUserRoleRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		_ = ctx.Error(err)
		return
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		_ = ctx.Error(err)
		return
	}

	if err := c.service.DeleteUserRole(ctx, &req); err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, &common.Response{
		Code: common.Ok,
	})
}

func (c *UserRoleController) ListUserRole(ctx *gin.Context) {
	if ok := utils2.CheckPermission(ctx, "user_role", "list"); !ok {
		ctx.JSON(http.StatusForbidden, &common.Response{
			Code: common.NoPermissionError,
			Msg:  "no permission",
		})
		return
	}

	var req define.ListUserRoleRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		_ = ctx.Error(err)
		return
	}

	if err := ctx.ShouldBindQuery(&req); err != nil {
		_ = ctx.Error(err)
		return
	}

	userRoles, err := c.service.ListUserRoleByUserId(ctx, req.UserId)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, &common.Response{
		Code: common.Ok,
		Data: userRoles,
	})
}
