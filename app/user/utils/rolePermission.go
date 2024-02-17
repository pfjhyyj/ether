package utils

import (
	"github.com/pfjhyyj/ether/app/user/define"
	"github.com/pfjhyyj/ether/app/user/model"
)

func ConvertRolePermissionListToResponse(rolePermissions []*model.RolePermission) []*define.RolePermissionResponse {
	var list []*define.RolePermissionResponse
	for _, rolePermission := range rolePermissions {
		list = append(list, &define.RolePermissionResponse{
			PermissionId:   rolePermission.PermissionId,
			PermissionName: rolePermission.PermissionName,
		})
	}
	return list
}
