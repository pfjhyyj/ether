package service

import (
	"github.com/gin-gonic/gin"
	"github.com/pfjhyyj/ether/app/tenant/define"
	"github.com/pfjhyyj/ether/app/tenant/model"
	"github.com/pfjhyyj/ether/app/tenant/utils"
	"github.com/pfjhyyj/ether/client/gorm"
	"github.com/pfjhyyj/ether/common"
	"github.com/sirupsen/logrus"
)

type TenantService struct {
}

func (s TenantService) ListTenants(ctx *gin.Context, param *model.QueryTenantParams) ([]*model.Tenant, int64, error) {
	logs := logrus.WithContext(ctx)
	db := gorm.GetDB().WithContext(ctx)

	tenants, total, err := model.ListTenants(db, param)
	if err != nil {
		logs.WithError(err).Error("list tenants failed")
		return nil, 0, &common.SystemError{Code: common.DbError, Msg: "list tenants failed", Err: err}
	}

	return tenants, total, nil
}

func (s TenantService) CreateTenant(ctx *gin.Context, d *define.CreateTenantRequest) error {
	logs := logrus.WithContext(ctx)
	db := gorm.GetDB().WithContext(ctx)

	tenant := utils.ConvertCreateTenantRequestToModel(d)

	if err := model.CreateTenant(db, tenant); err != nil {
		logs.WithError(err).Error("create tenant failed")
		return &common.SystemError{Code: common.DbError, Msg: "create tenant failed", Err: err}
	}

	return nil
}

func (s TenantService) UpdateTenant(ctx *gin.Context, d *define.UpdateTenantRequest) error {
	logs := logrus.WithContext(ctx)
	db := gorm.GetDB().WithContext(ctx)

	tenant := utils.ConvertUpdateTenantRequestToModel(d)

	if err := model.UpdateTenant(db, d.TenantId, tenant); err != nil {
		logs.WithError(err).Error("update tenant failed")
		return &common.SystemError{Code: common.DbError, Msg: "update tenant failed", Err: err}
	}

	return nil
}

func (s TenantService) DeleteTenant(ctx *gin.Context, d *define.DeleteTenantRequest) error {
	logs := logrus.WithContext(ctx)
	db := gorm.GetDB().WithContext(ctx)

	if err := model.DeleteTenant(db, d.TenantId); err != nil {
		logs.WithError(err).Error("delete tenant failed")
		return &common.SystemError{Code: common.DbError, Msg: "delete tenant failed", Err: err}
	}

	return nil
}

func NewTenantService() *TenantService {
	return &TenantService{}
}
