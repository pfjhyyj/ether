package define

import "github.com/pfjhyyj/ether/common"

type RoleIdUri struct {
	RoleId uint `uri:"roleId" binding:"required"`
}

type CreateRoleRequest struct {
	RoleName    string `json:"roleName" binding:"required"`
	RoleCode    string `json:"roleCode" binding:"required"`
	Description string `json:"description"`
}

type UpdateRoleRequest struct {
	RoleId      uint   `uri:"roleId" binding:"required"`
	RoleName    string `json:"roleName"`
	RoleCode    string `json:"roleCode"`
	Description string `json:"description"`
}

type DeleteRoleRequest struct {
	RoleId uint `uri:"roleId" binding:"required"`
}

type RoleResponse struct {
	RoleId      uint   `json:"roleId"`
	RoleName    string `json:"roleName"`
	RoleCode    string `json:"roleCode"`
	Description string `json:"description"`
}

type ListRoleRequest struct {
	common.PageRequest
	RoleName string `form:"roleName"`
}

type RolePageResponse struct {
	RoleId      uint   `json:"roleId"`
	RoleName    string `json:"roleName"`
	RoleCode    string `json:"roleCode"`
	Description string `json:"description"`
}
