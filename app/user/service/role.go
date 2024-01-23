package service

import (
	"github.com/gin-gonic/gin"
	"github.com/pfjhyyj/ether/app/user/model"
	"github.com/pfjhyyj/ether/clients/gorm"
	"github.com/pfjhyyj/ether/common"
	"github.com/sirupsen/logrus"
)

type RoleService struct{}

func NewRoleService() *RoleService {
	return &RoleService{}
}

func (s *RoleService) CreateRole(ctx *gin.Context, role *model.Role) error {
	logs := logrus.WithContext(ctx)
	db := gorm.GetDB().WithContext(ctx)

	if err := model.CreateRole(db, role); err != nil {
		logs.WithError(err).Error("create role failed")
		return &common.SystemError{Code: common.DbError, Msg: "create role failed", Err: err}
	}

	return nil
}

func (s *RoleService) UpdateRole(ctx *gin.Context, role *model.Role) error {
	logs := logrus.WithContext(ctx)
	db := gorm.GetDB().WithContext(ctx)

	if err := model.UpdateRole(db, role.RoleId, role); err != nil {
		logs.WithError(err).Error("update role failed")
		return &common.SystemError{Code: common.DbError, Msg: "update role failed", Err: err}
	}

	return nil
}

func (s *RoleService) DeleteRole(ctx *gin.Context, roleId uint) error {
	logs := logrus.WithContext(ctx)
	db := gorm.GetDB().WithContext(ctx)

	if err := model.DeleteRole(db, roleId); err != nil {
		logs.WithError(err).Error("delete role failed")
		return &common.SystemError{Code: common.DbError, Msg: "delete role failed", Err: err}
	}

	return nil
}

func (s *RoleService) GetRoleByRoleId(ctx *gin.Context, roleId uint) (*model.Role, error) {
	logs := logrus.WithContext(ctx)
	db := gorm.GetDB().WithContext(ctx)

	role, err := model.GetRoleByRoleId(db, roleId)
	if err != nil {
		logs.WithError(err).Error("get role by role id failed")
		return nil, &common.SystemError{Code: common.DbError, Msg: "get role by role id failed", Err: err}
	}

	return role, nil
}

func (s *RoleService) ListRoles(ctx *gin.Context, params *model.QueryRoleParams) ([]*model.Role, int64, error) {
	logs := logrus.WithContext(ctx)
	db := gorm.GetDB().WithContext(ctx)

	roles, total, err := model.ListRoles(db, params)
	if err != nil {
		logs.WithError(err).Error("list roles failed")
		return nil, 0, &common.SystemError{Code: common.DbError, Msg: "list roles failed", Err: err}
	}

	return roles, total, nil
}
