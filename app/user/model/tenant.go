package model

import (
	"errors"
	"github.com/pfjhyyj/ether/common"
	"gorm.io/gorm"
)

type Tenant struct {
	common.Model
	TenantId   uint   `gorm:"primaryKey"`
	TenantName string `gorm:"column:tenant_name"`
	TenantCode string `gorm:"column:tenant_code"`
	Domain     string `gorm:"column:domain"`
}

func (Tenant) TableName() string {
	return "tenant"
}

type QueryTenantParams struct {
	common.PageRequest
}

func CreateTenant(tx *gorm.DB, tenant *Tenant) error {
	return tx.Create(tenant).Error
}

func UpdateTenant(tx *gorm.DB, tenantId uint, tenant *Tenant) error {
	return tx.Where("tenant_id = ?", tenantId).Updates(tenant).Error
}

func DeleteTenant(tx *gorm.DB, tenantId uint) error {
	return tx.Delete(&Tenant{}, "tenant_id = ?", tenantId).Error
}

func GetTenantByTenantId(tx *gorm.DB, tenantId uint) (*Tenant, error) {
	var tenant Tenant
	if err := tx.First(&tenant, "tenant_id = ?", tenantId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &tenant, nil
}

func ListTenants(tx *gorm.DB, params *QueryTenantParams) ([]*Tenant, int64, error) {
	var tenants []*Tenant
	query := tx.Model(&Tenant{})

	var total int64
	query.Count(&total)

	if params.Current > 0 && params.PageSize > 0 {
		query = query.Offset((params.Current - 1) * params.PageSize).Limit(params.PageSize)
	}

	if err := query.Find(&tenants).Error; err != nil {
		return nil, 0, err
	}

	return tenants, total, nil
}
