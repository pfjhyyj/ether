package define

type AddRoleMenuRequest struct {
	RoleId  uint
	MenuIds []uint `json:"menuIds" binding:"required,gt=0"`
}

type DeleteRoleMenuRequest struct {
	RoleId  uint
	MenuIds []uint `json:"menuIds" binding:"required,gt=0"`
}

type SetRoleMenuRequest struct {
	RoleId  uint
	MenuIds []uint `json:"menuIds" binding:"required,gt=0"`
}

type ListRoleMenuRequest struct {
	RoleId uint `uri:"roleId" binding:"required"`
}

type ListRoleMenuResponse struct {
	MenuIds []uint `json:"menuIds"`
}
