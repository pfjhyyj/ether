package model

import (
	"github.com/pfjhyyj/ether/common"
	"gorm.io/gorm"
)

type Role struct {
	common.Model
	RoleId      uint   `gorm:"primaryKey"`
	RoleCode    string `gorm:"column:role_code"`
	RoleName    string `gorm:"column:role_name"`
	Description string `gorm:"column:description"`
}

func (Role) TableName() string {
	return "role"
}

type QueryRoleParams struct {
	common.PageRequest
}

func CreateRole(tx *gorm.DB, role *Role) error {
	return tx.Create(role).Error
}

func UpdateRole(tx *gorm.DB, roleId uint, role *Role) error {
	return tx.Where("role_id = ?", roleId).Updates(role).Error
}

func DeleteRole(tx *gorm.DB, roleId uint) error {
	return tx.Delete(&Role{}, "role_id = ?", roleId).Error
}

func GetRoleByRoleId(tx *gorm.DB, roleId uint) (*Role, error) {
	var role Role
	if err := tx.Where("role_id = ?", roleId).First(&role).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func GetRoleByRoleIds(tx *gorm.DB, roleIds []uint) ([]*Role, error) {
	if (len(roleIds)) == 0 {
		return nil, &common.SystemError{Code: common.DbError, Msg: "role ids is empty"}
	}
	var roles []*Role
	tx.Where("role_id IN ?", roleIds).Find(&roles)
	return roles, nil
}

func ListRoles(tx *gorm.DB, params *QueryRoleParams) ([]*Role, int64, error) {
	var roles []*Role
	query := tx.Model(&Role{})

	var total int64
	query.Count(&total)

	if params.Current > 0 && params.PageSize > 0 {
		query = query.Offset((params.Current - 1) * params.PageSize).Limit(params.PageSize)
	}

	if err := query.Find(&roles).Error; err != nil {
		return nil, 0, err
	}

	return roles, total, nil
}
