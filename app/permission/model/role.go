package model

import (
	"errors"
	"github.com/pfjhyyj/ether/common"
	"gorm.io/gorm"
)

type Role struct {
	common.Model
	RoleId      uint   `gorm:"primaryKey"`
	TenantId    uint   `gorm:"column:tenant_id"`
	RoleName    string `gorm:"column:role_name"`
	Description string `gorm:"column:description"`
}

func (Role) TableName() string {
	return "role"
}

type QueryRoleParams struct {
	common.PageRequest
	TenantId uint
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
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &role, nil
}

func ListRoles(tx *gorm.DB, params *QueryRoleParams) ([]*Role, int64, error) {
	var roles []*Role
	query := tx.Model(&Role{})

	var total int64
	query.Count(&total)

	if params.Current > 0 && params.PageSize > 0 {
		query = query.Offset((params.Current - 1) * params.PageSize).Limit(params.PageSize)
	}

	if params.TenantId > 0 {
		query = query.Where("tenant_id = ?", params.TenantId)
	}

	if err := query.Find(&roles).Error; err != nil {
		return nil, 0, err
	}

	return roles, total, nil
}
