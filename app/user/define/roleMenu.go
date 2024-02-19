package define

type AddRoleMenuRequest struct {
	RoleId  uint   `uri:"roleId" binding:"required"`
	MenuIds []uint `json:"menuIds" binding:"required,len>0"`
}

type DeleteRoleMenuRequest struct {
	RoleId  uint   `uri:"roleId" binding:"required"`
	MenuIds []uint `json:"menuIds" binding:"required,len>0"`
}

type ListRoleMenuRequest struct {
	RoleId uint `uri:"roleId" binding:"required"`
}

type ListRoleMenuResponse struct {
	MenuIds []uint `json:"menuIds"`
}
