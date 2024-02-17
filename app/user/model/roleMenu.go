package model

import (
	"github.com/pfjhyyj/ether/common"
	"gorm.io/gorm"
)

type RoleMenu struct {
	RoleId uint `gorm:"primaryKey;autoIncrement:false"`
	MenuId uint `gorm:"primaryKey;autoIncrement:false"`
	common.Model
}

func (rm RoleMenu) TableName() string {
	return "role_menu"
}

func CreateRoleMenu(tx *gorm.DB, roleMenu *RoleMenu) error {
	return tx.Create(roleMenu).Error
}

func CreateRoleMenuBatch(tx *gorm.DB, roleMenus []*RoleMenu) error {
	return tx.Create(roleMenus).Error
}

func DeleteRoleMenu(tx *gorm.DB, roleId uint, menuIds []uint) error {
	return tx.Unscoped().Delete(&RoleMenu{}, "role_id = ? AND menu_id IN ?", roleId, menuIds).Error
}

func ListMenuIdsByRoleId(tx *gorm.DB, roleId uint) ([]uint, error) {
	var roleMenus []*RoleMenu
	if err := tx.Select("menu_id").Where("role_id = ?", roleId).Find(&roleMenus).Error; err != nil {
		return nil, err
	}
	// parse to id array
	var menuIds []uint
	for _, roleMenu := range roleMenus {
		menuIds = append(menuIds, roleMenu.MenuId)
	}
	return menuIds, nil
}

func ListMenuIdsByRoleIds(tx *gorm.DB, roleIds []uint) ([]uint, error) {
	if len(roleIds) == 0 {
		return nil, &common.SystemError{Code: common.DbError, Msg: "role ids is empty"}
	}
	var roleMenus []*RoleMenu
	tx.Where("role_id IN ?", roleIds).Distinct("menu_id").Find(&roleMenus)

	var menuIds []uint
	for _, roleMenu := range roleMenus {
		menuIds = append(menuIds, roleMenu.MenuId)
	}
	return menuIds, nil
}
