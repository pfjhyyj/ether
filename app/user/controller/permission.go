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

type PermissionController struct {
	service *service.PermissionService
}

func NewPermissionController(service *service.PermissionService) *PermissionController {
	return &PermissionController{
		service: service,
	}
}

// CreatePermission godoc
// @Summary Create permission
// @Description Create permission
// @Tags permission
// @Accept json
// @Produce json
// @Security Bearer
// @Param request body define.CreatePermissionRequest true "CreatePermissionRequest"
// @Success 200 {object} common.Response
// @Router /permissions [post]
func (c *PermissionController) CreatePermission(ctx *gin.Context) {
	if ok := utils2.CheckPermission(ctx, "permission", "create"); !ok {
		ctx.JSON(http.StatusForbidden, &common.Response{
			Code: common.NoPermissionError,
			Msg:  "no permission",
		})
		return
	}

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

// UpdatePermission godoc
// @Summary Update permission
// @Description Update permission
// @Tags permission
// @Accept json
// @Produce json
// @Security Bearer
// @Param permission_id path int true "permission_id"
// @Param request body define.UpdatePermissionRequest true "UpdatePermissionRequest"
// @Success 200 {object} common.Response
// @Router /permissions/{permission_id} [put]
func (c *PermissionController) UpdatePermission(ctx *gin.Context) {
	if ok := utils2.CheckPermission(ctx, "permission", "update"); !ok {
		ctx.JSON(http.StatusForbidden, &common.Response{
			Code: common.NoPermissionError,
			Msg:  "no permission",
		})
		return
	}

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

// DeletePermission godoc
// @Summary Delete permission
// @Description Delete permission
// @Tags permission
// @Accept json
// @Produce json
// @Security Bearer
// @Param permission_id path int true "permission_id"
// @Success 200 {object} common.Response
// @Router /permissions/{permission_id} [delete]
func (c *PermissionController) DeletePermission(ctx *gin.Context) {
	if ok := utils2.CheckPermission(ctx, "permission", "delete"); !ok {
		ctx.JSON(http.StatusForbidden, &common.Response{
			Code: common.NoPermissionError,
			Msg:  "no permission",
		})
		return
	}

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

// ListPermissions godoc
// @Summary List permissions
// @Description List permissions
// @Tags permission
// @Accept json
// @Produce json
// @Security Bearer
// @Param current query int false "current"
// @Param page_size query int false "page_size"
// @Param target query string false "name"
// @Success 200 {object} common.Response{data=common.Page{list=[]define.PermissionPageResponse}}
// @Router /permissions [get]
func (c *PermissionController) ListPermissions(ctx *gin.Context) {
	if ok := utils2.CheckPermission(ctx, "permission", "list"); !ok {
		ctx.JSON(http.StatusForbidden, &common.Response{
			Code: common.NoPermissionError,
			Msg:  "no permission",
		})
		return
	}

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
