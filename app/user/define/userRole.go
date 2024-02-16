package define

import "github.com/pfjhyyj/ether/common"

type AddUserRoleRequest struct {
	RoleIds []uint `json:"roleIds" binding:"required,gt=0"`
}

type DeleteUserRoleRequest struct {
	RoleIds []uint `json:"roleIds" binding:"required,gt=0"`
}

type ListUserRoleRequest struct {
	common.PageRequest
	UserId uint
}

type ListUserRoleResponse struct {
	RoleId   uint   `json:"roleId"`
	RoleName string `json:"roleName"`
}
