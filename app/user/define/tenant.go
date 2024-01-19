package define

import (
	"github.com/pfjhyyj/ether/common"
	"time"
)

type CreateTenantRequest struct {
	TenantName string `json:"tenant_name" binding:"required"`
	TenantCode string `json:"tenant_code" binding:"required"`
	Domain     string `json:"domain" binding:"required"`
}

type UpdateTenantRequest struct {
	TenantId   uint   `uri:"tenant_id" binding:"required"`
	TenantName string `json:"tenant_name"`
	TenantCode string `json:"tenant_code"`
	Domain     string `json:"domain"`
}

type DeleteTenantRequest struct {
	TenantId uint `uri:"tenant_id" binding:"required"`
}

type ListTenantRequest struct {
	common.PageRequest
}

type ListTenantPageResponse struct {
	TenantId   uint
	TenantName string
	TenantCode string
	Domain     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
