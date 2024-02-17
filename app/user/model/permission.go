package model

import (
	"github.com/pfjhyyj/ether/common"
	"gorm.io/gorm"
)

type Permission struct {
	PermissionId uint   `gorm:"primaryKey;column:permission_id"`
	Name         string `gorm:"column:name"`
	Target       string `gorm:"column:target"`
	Action       string `gorm:"column:action"`
	Description  string `gorm:"column:description"`
	common.Model
}

func (p Permission) TableName() string {
	return "permission"
}

type QueryPermissionParams struct {
	common.PageRequest
	Target string
	Name   string
}

func CreatePermission(tx *gorm.DB, perm *Permission) error {
	return tx.Create(perm).Error
}

func UpdatePermission(tx *gorm.DB, permId uint, perm *Permission) error {
	return tx.Where("permission_id = ?", permId).Updates(perm).Error
}

func DeletePermission(tx *gorm.DB, permId uint) error {
	return tx.Delete(&Permission{}, "permission_id = ?", permId).Error
}

func GetPermissionByPermissionId(tx *gorm.DB, permId uint) (*Permission, error) {
	var perm Permission
	if err := tx.Where("permission_id = ?", permId).First(&perm).Error; err != nil {
		return nil, err
	}
	return &perm, nil
}

func GetPermissionByPermissionIds(tx *gorm.DB, permIds []uint) ([]*Permission, error) {
	if (len(permIds)) == 0 {
		return nil, &common.SystemError{Code: common.DbError, Msg: "permission ids is empty"}
	}
	var perms []*Permission
	tx.Where("permission_id IN ?", permIds).Find(&perms)
	return perms, nil
}

func ListPermissions(tx *gorm.DB, params *QueryPermissionParams) ([]*Permission, int64, error) {
	var perms []*Permission
	query := tx.Model(&Permission{})

	var total int64
	query.Count(&total)

	if params.Current > 0 && params.PageSize > 0 {
		query = query.Offset((params.Current - 1) * params.PageSize).Limit(params.PageSize)
	}

	if params.Target != "" {
		query = query.Where("target = ?", params.Target)
	}

	if params.Name != "" {
		query = query.Where("name = ?", params.Name)
	}

	if err := query.Find(&perms).Error; err != nil {
		return nil, 0, err
	}

	return perms, total, nil
}
