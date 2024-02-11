package model

import (
	"github.com/pfjhyyj/ether/common"
	"gorm.io/gorm"
)

type RolePermission struct {
	RoleId       uint `gorm:"primaryKey;autoIncrement:false"`
	PermissionId uint `gorm:"primaryKey;autoIncrement:false"`
	common.Model
}

func (RolePermission) TableName() string {
	return "role_permission"
}

func CreateRolePermissionBatch(tx *gorm.DB, rolePermissions []*RolePermission) error {
	return tx.Create(rolePermissions).Error
}

func DeleteRolePermissionBatch(tx *gorm.DB, roleId uint, permissionIds []uint) error {
	return tx.Delete(&RolePermission{}, "role_id = ? AND permission_id IN ?", roleId, permissionIds).Error
}

func ListPermissionIdsByRoleId(tx *gorm.DB, roleId uint) ([]uint, error) {
	var rolePermissions []*RolePermission
	if err := tx.Select("permission_id").Where("role_id = ?", roleId).Find(&rolePermissions).Error; err != nil {
		return nil, err
	}
	// parse to id array
	var permissionIds []uint
	for _, rolePermission := range rolePermissions {
		permissionIds = append(permissionIds, rolePermission.PermissionId)
	}
	return permissionIds, nil
}
