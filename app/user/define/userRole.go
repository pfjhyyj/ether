package define

type AddUserRoleRequest struct {
	UserId  uint   `json:"userId" binding:"required"`
	RoleIds []uint `json:"roleIds" binding:"required,len>0"`
}

type DeleteUserRoleRequest struct {
	UserId  uint   `json:"userId" binding:"required"`
	RoleIds []uint `json:"roleIds" binding:"required,len>0"`
}

type ListUserRoleRequest struct {
	UserId uint `form:"userId" binding:"required"`
}
