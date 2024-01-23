package define

import (
	"github.com/pfjhyyj/ether/common"
	"time"
)

type CreateTenantRequest struct {
	Name   string `json:"name" binding:"required"`
	Domain string `json:"domain" binding:"required"`
}

type UpdateTenantRequest struct {
	TenantId uint   `uri:"tenant_id" binding:"required"`
	Name     string `json:"name"`
	Domain   string `json:"domain"`
}

type DeleteTenantRequest struct {
	TenantId uint `uri:"tenant_id" binding:"required"`
}

type ListTenantRequest struct {
	common.PageRequest
}

type ListTenantPageResponse struct {
	TenantId  uint
	Name      string
	Domain    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
