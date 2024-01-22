package tenant

import "context"

type Repository interface {
	GetTenantByDomain(ctx context.Context, domain string) (*Tenant, error)
}
