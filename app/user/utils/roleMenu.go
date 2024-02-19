package utils

import (
	"github.com/pfjhyyj/ether/app/user/define"
)

func ConvertRoleMenuListToResponse(menuIds []uint) define.ListRoleMenuResponse {
	return define.ListRoleMenuResponse{
		MenuIds: menuIds,
	}
}
