package model

import (
	"github.com/pfjhyyj/ether/common"
	"gorm.io/gorm"
)

type UserRole struct {
	common.Model
	UserId uint `gorm:"primaryKey;autoIncrement:false"`
	RoleId uint `gorm:"primaryKey;autoIncrement:false"`

	RoleName string
}

func (UserRole) TableName() string {
	return "user_role"
}

func CreateUserRoleBatch(tx *gorm.DB, userRoles []*UserRole) error {
	return tx.Create(userRoles).Error
}

func DeleteUserRoleBatch(tx *gorm.DB, userId uint, roleIds []uint) error {
	return tx.Delete(&UserRole{}, "user_id = ? AND role_id IN ?", userId, roleIds).Error
}

func DeleteUserRolesByUserId(tx *gorm.DB, userId uint) error {
	return tx.Delete(&UserRole{}, "user_id = ?", userId).Error
}

func DeleteUserRolesByRoleId(tx *gorm.DB, roleId uint) error {
	return tx.Delete(&UserRole{}, "role_id = ?", roleId).Error
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

func ListUserRolesByUserId(tx *gorm.DB, userId uint) ([]*UserRole, error) {
	var userRoles []*UserRole
	if err := tx.Raw(""+
		"SELECT ur.role_id, r.role_name "+
		"FROM user_role ur "+
		"LEFT JOIN role r ON ur.role_id = r.role_id "+
		"WHERE ur.user_id = ?", userId,
	).Scan(&userRoles).Error; err != nil {
		return nil, err
	}
	return userRoles, nil
}
