package define

type AddRolePermissionRequest struct {
	RoleId        uint   `json:"roleId" binding:"required"`
	PermissionIds []uint `json:"permissionIds" binding:"required,len>0"`
}

type DeleteRolePermissionRequest struct {
	RoleId        uint   `json:"roleId" binding:"required"`
	PermissionIds []uint `json:"permissionIds" binding:"required,len>0"`
}

type ListRolePermissionRequest struct {
	RoleId uint `form:"roleId" binding:"required"`
}
