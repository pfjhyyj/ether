package service

import (
	"context"
	"github.com/pfjhyyj/ether/app/permission/define"
	"github.com/pfjhyyj/ether/app/permission/model"
	"github.com/pfjhyyj/ether/clients/gorm"
	"github.com/pfjhyyj/ether/common"
	"github.com/sirupsen/logrus"
)

type UserRoleService struct{}

func (s *UserRoleService) AddUserRole(ctx context.Context, d *define.AddUserRoleRequest) error {
	logs := logrus.WithContext(ctx)
	db := gorm.GetDB().WithContext(ctx)

	var userRoles []*model.UserRole
	for i := 0; i < len(d.RoleIds); i++ {
		userRoles = append(userRoles, &model.UserRole{
			UserId: d.UserId,
			RoleId: d.RoleIds[i],
		})
	}
	if err := model.CreateUserRoleBatch(db, userRoles); err != nil {
		logs.WithError(err).Error("add user role failed")
		return &common.SystemError{Code: common.DbError, Msg: "add user role failed", Err: err}
	}

	return nil
}

func (s *UserRoleService) DeleteUserRole(ctx context.Context, d *define.DeleteUserRoleRequest) error {
	logs := logrus.WithContext(ctx)
	db := gorm.GetDB().WithContext(ctx)

	if err := model.DeleteUserRoleBatch(db, d.UserId, d.RoleIds); err != nil {
		logs.WithError(err).Error("delete user role failed")
		return &common.SystemError{Code: common.DbError, Msg: "delete user role failed", Err: err}
	}

	return nil
}

func (s *UserRoleService) ListUserRoleByUserId(ctx context.Context, userId uint) ([]*model.UserRole, error) {
	logs := logrus.WithContext(ctx)
	db := gorm.GetDB().WithContext(ctx)

	userRoles, err := model.ListUserRolesByUserId(db, userId)
	if err != nil {
		logs.WithError(err).Error("list user role by user id failed")
		return nil, &common.SystemError{Code: common.DbError, Msg: "list user role by user id failed", Err: err}
	}

	return userRoles, nil
}

func NewUserRoleService() *UserRoleService {
	return &UserRoleService{}
}
