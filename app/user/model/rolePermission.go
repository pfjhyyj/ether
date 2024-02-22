package model

import (
	"github.com/pfjhyyj/ether/app/user/define"
	"github.com/pfjhyyj/ether/common"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type RolePermission struct {
	RoleId       uint `gorm:"primaryKey;autoIncrement:false"`
	PermissionId uint `gorm:"primaryKey;autoIncrement:false"`
	common.Model

	PermissionName string
}

func (RolePermission) TableName() string {
	return "role_permission"
}

func CreateRolePermissionBatch(tx *gorm.DB, rolePermissions []*RolePermission) error {
	return tx.Clauses(clause.OnConflict{DoNothing: true}).Create(rolePermissions).Error
}

func DeleteRolePermissionBatch(tx *gorm.DB, roleId uint, permissionIds []uint) error {
	return tx.Unscoped().Delete(&RolePermission{}, "role_id = ? AND permission_id IN ?", roleId, permissionIds).Error
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

func ListRolePermissions(tx *gorm.DB, d *define.ListRolePermissionRequest) ([]*RolePermission, int64, error) {
	var rolePermissions []*RolePermission
	query := tx.Model(&RolePermission{}).Select("role_permission.permission_id, p.name as permission_name").Joins("LEFT JOIN permission p ON role_permission.permission_id = p.permission_id").Where("role_permission.role_id = ?", d.RoleId)

	var total int64
	query.Count(&total)

	if d.Current > 0 && d.PageSize > 0 {
		query = query.Offset((d.Current - 1) * d.PageSize).Limit(d.PageSize)
	}

	if err := query.Find(&rolePermissions).Error; err != nil {
		return nil, 0, err
	}

	return rolePermissions, total, nil
}
