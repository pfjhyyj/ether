package model

import (
	"github.com/pfjhyyj/ether/common"
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
