package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pfjhyyj/ether/app/tenant/define"
	"github.com/pfjhyyj/ether/app/tenant/service"
	"github.com/pfjhyyj/ether/app/tenant/utils"
	"github.com/pfjhyyj/ether/common"
	utils2 "github.com/pfjhyyj/ether/utils"
	"net/http"
)

type TenantController struct {
	service *service.TenantService
}

func NewTenantController(service *service.TenantService) *TenantController {
	return &TenantController{service: service}
}

func (r *TenantController) CreateTenant(ctx *gin.Context) {
	if ok := utils2.CheckPermission(ctx, "tenant", "create"); !ok {
		ctx.JSON(http.StatusForbidden, &common.Response{
			Code: common.NoPermissionError,
			Msg:  "no permission",
		})
		return
	}

	var req define.CreateTenantRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		_ = ctx.Error(err)
		return
	}

	err := r.service.CreateTenant(ctx, &req)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, &common.Response{
		Code: common.Ok,
	})
}

func (r *TenantController) UpdateTenant(ctx *gin.Context) {
	if ok := utils2.CheckPermission(ctx, "tenant", "update"); !ok {
		ctx.JSON(http.StatusForbidden, &common.Response{
			Code: common.NoPermissionError,
			Msg:  "no permission",
		})
		return
	}

	var req define.UpdateTenantRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		_ = ctx.Error(err)
		return
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		_ = ctx.Error(err)
		return
	}

	err := r.service.UpdateTenant(ctx, &req)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(200, &common.Response{
		Code: common.Ok,
	})
}

func (r *TenantController) DeleteTenant(ctx *gin.Context) {
	if ok := utils2.CheckPermission(ctx, "tenant", "delete"); !ok {
		ctx.JSON(http.StatusForbidden, &common.Response{
			Code: common.NoPermissionError,
			Msg:  "no permission",
		})
		return
	}

	var req define.DeleteTenantRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		_ = ctx.Error(err)
		return
	}

	err := r.service.DeleteTenant(ctx, &req)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, &common.Response{
		Code: common.Ok,
	})
}

func (r *TenantController) ListTenants(ctx *gin.Context) {
	if ok := utils2.CheckPermission(ctx, "tenant", "list"); !ok {
		ctx.JSON(http.StatusForbidden, &common.Response{
			Code: common.NoPermissionError,
			Msg:  "no permission",
		})
		return
	}

	var req define.ListTenantRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		_ = ctx.Error(err)
		return
	}

	queryParam := utils.ConvertTenantListPageRequestToParam(&req)

	tenants, total, err := r.service.ListTenants(ctx, queryParam)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	tenantInfo := utils.ConvertTenantListToPageResponse(tenants)
	ctx.JSON(http.StatusOK, &common.Response{
		Code: common.Ok,
		Data: &common.Page{
			Total:    total,
			PageSize: req.PageSize,
			Current:  req.Current,
			List:     tenantInfo,
		},
	})
}
