package service

import (
	"fmt"
	casbin2 "github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
	"github.com/pfjhyyj/ether/app/user/define"
	"github.com/pfjhyyj/ether/app/user/model"
	"github.com/pfjhyyj/ether/clients/casbin"
	"github.com/pfjhyyj/ether/clients/gorm"
	"github.com/pfjhyyj/ether/common"
	"github.com/sirupsen/logrus"
	gorm2 "gorm.io/gorm"
)

type UserRoleService struct{}

func (s *UserRoleService) AddUserRole(ctx *gin.Context, userId uint, roleIds []uint) error {
	logs := logrus.WithContext(ctx)
	db := gorm.GetDB().WithContext(ctx)
	e := casbin.GetEnforcer()

	var userRoles []*model.UserRole
	for i := 0; i < len(roleIds); i++ {
		userRoles = append(userRoles, &model.UserRole{
			UserId: userId,
			RoleId: roleIds[i],
		})
	}

	roles, err := model.GetRoleByRoleIds(db, roleIds)
	if err != nil {
		logs.WithError(err).Error("get role by role ids failed")
		return &common.SystemError{Code: common.DbError, Msg: "get role by role ids failed", Err: err}
	}

	err = db.Transaction(func(tx *gorm2.DB) error {
		if err := model.CreateUserRoleBatch(tx, userRoles); err != nil {
			logs.WithError(err).Error("add user role failed")
			return err
		}
		err := e.GetAdapter().(*gormadapter.Adapter).Transaction(e, func(e casbin2.IEnforcer) error {
			userIdStr := fmt.Sprintf("%d", userId)
			for i := 0; i < len(roles); i++ {
				role := roles[i]
				_, err := e.AddRoleForUser(userIdStr, role.RoleCode)
				if err != nil {
					return err
				}
			}
			return nil
		})

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		logs.WithError(err).Error("add user role failed")
		return &common.SystemError{Code: common.DbError, Msg: "add user role failed", Err: err}
	}

	return nil
}

func (s *UserRoleService) DeleteUserRole(ctx *gin.Context, userId uint, roleIds []uint) error {
	logs := logrus.WithContext(ctx)
	db := gorm.GetDB().WithContext(ctx)

	roles, err := model.GetRoleByRoleIds(db, roleIds)
	if err != nil {
		logs.WithError(err).Error("get role by role ids failed")
		return &common.SystemError{Code: common.DbError, Msg: "get role by role ids failed", Err: err}
	}

	err = db.Transaction(func(tx *gorm2.DB) error {
		if err := model.DeleteUserRoleBatch(tx, userId, roleIds); err != nil {
			logs.WithError(err).Error("delete user role failed")
			return err
		}
		e := casbin.GetEnforcer()
		err := e.GetAdapter().(*gormadapter.Adapter).Transaction(e, func(e casbin2.IEnforcer) error {
			userIdStr := fmt.Sprintf("%d", userId)
			for i := 0; i < len(roles); i++ {
				role := roles[i]
				_, err := e.DeleteRoleForUser(userIdStr, role.RoleCode)
				if err != nil {
					return err
				}
			}
			return nil
		})

		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		logs.WithError(err).Error("delete user role failed")
		return &common.SystemError{Code: common.DbError, Msg: "delete user role failed", Err: err}
	}

	return nil
}

func (s *UserRoleService) ListUserRoleByUserId(ctx *gin.Context, d *define.ListUserRoleRequest) ([]*model.UserRole, int64, error) {
	logs := logrus.WithContext(ctx)
	db := gorm.GetDB().WithContext(ctx)

	userRoles, total, err := model.ListUserRoles(db, d)
	if err != nil {
		logs.WithError(err).Error("list user role by user id failed")
		return nil, 0, &common.SystemError{Code: common.DbError, Msg: "list user role by user id failed", Err: err}
	}

	return userRoles, total, nil
}

func NewUserRoleService() *UserRoleService {
	return &UserRoleService{}
}
