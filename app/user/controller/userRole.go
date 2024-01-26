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

// AddUserRole godoc
// @Summary Add user role
// @Description Add user role
// @Tags user
// @Accept json
// @Produce json
// @Security Bearer
// @Param user_id path int true "user_id"
// @Param request body define.AddUserRoleRequest true "AddUserRoleRequest"
// @Success 200 {object} string
// @Router /users/{userId}/roles/add [post]
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

// DeleteUserRole godoc
// @Summary Delete user role
// @Description Delete user role
// @Tags user
// @Accept json
// @Produce json
// @Security Bearer
// @Param user_id path int true "user_id"
// @Param request body define.DeleteUserRoleRequest true "DeleteUserRoleRequest"
// @Success 200 {object} string
// @Router /users/{userId}/roles/delete [post]
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

// ListUserRole godoc
// @Summary List user role
// @Description List user role
// @Tags user
// @Accept json
// @Produce json
// @Security Bearer
// @Param user_id path int true "user_id"
// @Param request body define.ListUserRoleRequest true "ListUserRoleRequest"
// @Success 200 {object} string
// @Router /users/{userId}/roles [get]
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
