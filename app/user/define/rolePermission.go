package define

import "github.com/pfjhyyj/ether/common"

type AddRolePermissionRequest struct {
	RoleId        uint
	PermissionIds []uint `json:"permissionIds" binding:"required,gt=0"`
}

type DeleteRolePermissionRequest struct {
	RoleId        uint
	PermissionIds []uint `json:"permissionIds" binding:"required,gt=0"`
}

type ListRolePermissionRequest struct {
	common.PageRequest

	RoleId uint
}

type RolePermissionResponse struct {
	PermissionId   uint   `json:"permissionId"`
	PermissionName string `json:"permissionName"`
}
