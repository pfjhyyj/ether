package repository

import (
	"errors"
	"github.com/pfjhyyj/ether/app/user/model"
	"gorm.io/gorm"
)

type TenantRepository struct {
	db *gorm.DB
}

func NewTenantRepository(db *gorm.DB) *TenantRepository {
	return &TenantRepository{db: db}
}

func (r *TenantRepository) CreateTenant(tenant *model.Tenant) error {
	return r.db.Create(tenant).Error
}

func (r *TenantRepository) UpdateTenant(tenantId uint, tenant *model.Tenant) error {
	return r.db.Where("tenant_id = ?", tenantId).Updates(tenant).Error
}

func (r *TenantRepository) DeleteTenant(tenantId uint) error {
	return r.db.Delete(&model.Tenant{}, "tenant_id = ?", tenantId).Error
}

func (r *TenantRepository) GetTenantByTenantId(tenantId uint) (*model.Tenant, error) {
	var tenant model.Tenant
	if err := r.db.First(&tenant, "tenant_id = ?", tenantId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &tenant, nil
}

func (r *TenantRepository) ListTenants(params *model.QueryTenantParams) ([]*model.Tenant, int64, error) {
	var tenants []*model.Tenant
	query := r.db.Model(&model.Tenant{})

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
