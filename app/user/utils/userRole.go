package utils

import (
	"github.com/pfjhyyj/ether/app/user/define"
	"github.com/pfjhyyj/ether/app/user/model"
)

func ConvertUserRoleListToResponse(roles []*model.UserRole) []*define.ListUserRoleResponse {
	var roleInfos []*define.ListUserRoleResponse
	for _, role := range roles {
		roleInfos = append(roleInfos, &define.ListUserRoleResponse{
			RoleId:   role.RoleId,
			RoleName: role.RoleName,
		})
	}
	return roleInfos
}
