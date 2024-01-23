package utils

import (
	"github.com/pfjhyyj/ether/app/user/define"
	"github.com/pfjhyyj/ether/app/user/model"
	"github.com/pfjhyyj/ether/common"
)

func ConvertCreateRoleRequestToRole(req *define.CreateRoleRequest) *model.Role {
	return &model.Role{
		TenantId: req.TenantId,
		RoleName: req.RoleName,
	}
}

func ConvertUpdateRoleRequestToRole(req *define.UpdateRoleRequest) *model.Role {
	return &model.Role{
		RoleId:      req.RoleId,
		TenantId:    req.TenantId,
		RoleName:    req.RoleName,
		Description: req.Description,
	}
}

func ConvertRoleListPageRequestToParam(req *define.ListRoleRequest) *model.QueryRoleParams {
	return &model.QueryRoleParams{
		PageRequest: common.PageRequest{
			Current:  req.Current,
			PageSize: req.PageSize,
		},
		TenantId: req.TenantId,
	}
}

func ConvertRoleListToPageResponse(roles []*model.Role) []*define.RolePageResponse {
	var roleInfos []*define.RolePageResponse
	for _, role := range roles {
		roleInfos = append(roleInfos, &define.RolePageResponse{
			RoleId:      role.RoleId,
			TenantId:    role.TenantId,
			RoleName:    role.RoleName,
			Description: role.Description,
		})
	}
	return roleInfos
}
