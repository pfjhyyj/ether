package domain

import (
	"context"
	"github.com/pfjhyyj/ether/app/tenant/model"
	"github.com/pfjhyyj/ether/client/gorm"
	"github.com/pfjhyyj/ether/domain/tenant"
	"sync"
)

type TenantRepository struct {
	tenant.Repository
}

var (
	tenantRepository *TenantRepository

	once sync.Once
)

func Init() {
	once.Do(func() {
		tenantRepository = &TenantRepository{}
	})
}

func GetTenantRepository() *TenantRepository {
	Init()
	return tenantRepository
}

func (r *TenantRepository) GetTenantByDomain(ctx context.Context, domain string) (*tenant.Tenant, error) {
	db := gorm.GetDB().WithContext(ctx)
	tenantModel, err := model.GetTenantByDomain(db, domain)
	if err != nil {
		return nil, err
	}

	t := &tenant.Tenant{
		TenantId: tenantModel.TenantId,
		Domain:   tenantModel.Domain,
	}

	return t, nil
}
