package model

import (
	"github.com/pfjhyyj/ether/app/user/define"
	"github.com/pfjhyyj/ether/common"
	"gorm.io/gorm"
)

type UserRole struct {
	UserId uint `gorm:"primaryKey;autoIncrement:false"`
	RoleId uint `gorm:"primaryKey;autoIncrement:false"`
	common.Model

	RoleName string
}

func (UserRole) TableName() string {
	return "user_role"
}

func CreateUserRoleBatch(tx *gorm.DB, userRoles []*UserRole) error {
	return tx.Create(userRoles).Error
}

func DeleteUserRoleBatch(tx *gorm.DB, userId uint, roleIds []uint) error {
	return tx.Unscoped().Delete(&UserRole{}, "user_id = ? AND role_id IN ?", userId, roleIds).Error
}

func ListUserRoleIdsByUserId(tx *gorm.DB, userId uint) ([]uint, error) {
	var userRoles []*UserRole
	if err := tx.Select("role_id").Where("user_id = ?", userId).Find(&userRoles).Error; err != nil {
		return nil, err
	}
	// parse user role to id array
	var roleIds []uint
	for _, userRole := range userRoles {
		roleIds = append(roleIds, userRole.RoleId)
	}
	return roleIds, nil
}

func ListUserRoles(tx *gorm.DB, d *define.ListUserRoleRequest) ([]*UserRole, int64, error) {
	var userRoles []*UserRole
	query := tx.Model(&UserRole{}).Select("user_role.role_id, role.role_name").Joins("LEFT JOIN role ON user_role.role_id = role.role_id").Where("user_id = ?", d.UserId)

	var total int64
	query.Count(&total)

	if d.Current > 0 && d.PageSize > 0 {
		query = query.Offset((d.Current - 1) * d.PageSize).Limit(d.PageSize)
	}

	if err := query.Scan(&userRoles).Error; err != nil {
		return nil, 0, err
	}
	return userRoles, total, nil
}
