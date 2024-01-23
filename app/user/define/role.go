package define

import "github.com/pfjhyyj/ether/common"

type CreateRoleRequest struct {
	TenantId    uint   `json:"tenantId" binding:"required"`
	RoleName    string `json:"roleName" binding:"required"`
	Description string `json:"description"`
}

type UpdateRoleRequest struct {
	RoleId      uint   `uri:"roleId" binding:"required"`
	TenantId    uint   `json:"tenantId"`
	RoleName    string `json:"roleName"`
	Description string `json:"description"`
}

type DeleteRoleRequest struct {
	RoleId uint `uri:"roleId" binding:"required"`
}

type ListRoleRequest struct {
	common.PageRequest
	TenantId uint `form:"tenantId"`
}

type RolePageResponse struct {
	RoleId      uint   `json:"roleId"`
	TenantId    uint   `json:"tenantId"`
	RoleName    string `json:"roleName"`
	Description string `json:"description"`
}
