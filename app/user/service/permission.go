package service

import (
	"github.com/gin-gonic/gin"
	"github.com/pfjhyyj/ether/app/user/model"
	"github.com/pfjhyyj/ether/clients/gorm"
	"github.com/pfjhyyj/ether/common"
	"github.com/sirupsen/logrus"
)

type PermissionService struct {
}

func NewPermissionService() *PermissionService {
	return &PermissionService{}
}

func (s *PermissionService) CreatePermission(ctx *gin.Context, permission *model.Permission) error {
	logs := logrus.WithContext(ctx)
	db := gorm.GetDB().WithContext(ctx)

	if err := model.CreatePermission(db, permission); err != nil {
		logs.WithError(err).Error("create permission failed")
		return &common.SystemError{Code: common.DbError, Msg: "create permission failed", Err: err}
	}

	return nil
}

func (s *PermissionService) UpdatePermission(ctx *gin.Context, permission *model.Permission) error {
	logs := logrus.WithContext(ctx)
	db := gorm.GetDB().WithContext(ctx)

	if err := model.UpdatePermission(db, permission.PermissionId, permission); err != nil {
		logs.WithError(err).Error("update permission failed")
		return &common.SystemError{Code: common.DbError, Msg: "update permission failed", Err: err}
	}

	return nil
}

func (s *PermissionService) DeletePermission(ctx *gin.Context, permissionId uint) error {
	logs := logrus.WithContext(ctx)
	db := gorm.GetDB().WithContext(ctx)

	if err := model.DeletePermission(db, permissionId); err != nil {
		logs.WithError(err).Error("delete permission failed")
		return &common.SystemError{Code: common.DbError, Msg: "delete permission failed", Err: err}
	}

	return nil
}

func (s *PermissionService) GetPermissionByPermissionId(ctx *gin.Context, permissionId uint) (*model.Permission, error) {
	logs := logrus.WithContext(ctx)
	db := gorm.GetDB().WithContext(ctx)

	permission, err := model.GetPermissionByPermissionId(db, permissionId)
	if err != nil {
		logs.WithError(err).Error("get permission by permission id failed")
		return nil, &common.SystemError{Code: common.DbError, Msg: "get permission by permission id failed", Err: err}
	}
	return permission, nil
}

func (s *PermissionService) ListPermissions(ctx *gin.Context, params *model.QueryPermissionParams) ([]*model.Permission, int64, error) {
	logs := logrus.WithContext(ctx)
	db := gorm.GetDB().WithContext(ctx)

	permissions, total, err := model.ListPermissions(db, params)
	if err != nil {
		logs.WithError(err).Error("list permissions failed")
		return nil, 0, &common.SystemError{Code: common.DbError, Msg: "list permissions failed", Err: err}
	}
	return permissions, total, nil
}
