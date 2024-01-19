package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pfjhyyj/ether/app/user/define"
	"github.com/pfjhyyj/ether/app/user/service"
	"github.com/pfjhyyj/ether/app/user/utils"
	"github.com/pfjhyyj/ether/common"
	"net/http"
)

type TenantController struct {
	service *service.TenantService
}

func NewTenantController(service *service.TenantService) *TenantController {
	return &TenantController{service: service}
}

func (r *TenantController) CreateTenant(ctx *gin.Context) {
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
