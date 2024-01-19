package utils

import (
	"github.com/pfjhyyj/ether/app/user/define"
	"github.com/pfjhyyj/ether/app/user/model"
)

func ConvertTenantListPageRequestToParam(req *define.ListTenantRequest) *model.QueryTenantParams {
	queryParam := &model.QueryTenantParams{
		PageRequest: req.PageRequest,
	}
	return queryParam
}

func ConvertTenantListToPageResponse(tenants []*model.Tenant) []*define.ListTenantPageResponse {
	tenantInfo := make([]*define.ListTenantPageResponse, 0, len(tenants))
	for _, tenant := range tenants {
		tenantInfo = append(tenantInfo, &define.ListTenantPageResponse{
			TenantId:   tenant.TenantId,
			TenantName: tenant.TenantName,
			CreatedAt:  tenant.CreatedAt,
			UpdatedAt:  tenant.UpdatedAt,
		})
	}
	return tenantInfo
}

func ConvertCreateTenantRequestToModel(req *define.CreateTenantRequest) *model.Tenant {
	return &model.Tenant{
		TenantName: req.TenantName,
		TenantCode: req.TenantCode,
		Domain:     req.Domain,
	}
}

func ConvertUpdateTenantRequestToModel(req *define.UpdateTenantRequest) *model.Tenant {
	return &model.Tenant{
		TenantName: req.TenantName,
		TenantCode: req.TenantCode,
		Domain:     req.Domain,
	}
}
