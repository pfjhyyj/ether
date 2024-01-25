package service

import (
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

type RolePermissionService struct {
}

func NewRolePermissionService() *RolePermissionService {
	return &RolePermissionService{}
}

func (s *RolePermissionService) AddRolePermission(ctx *gin.Context, d *define.AddRolePermissionRequest) error {
	logs := logrus.WithContext(ctx)
	db := gorm.GetDB().WithContext(ctx)
	e := casbin.GetEnforcer()

	var rolePermissions []*model.RolePermission
	for i := 0; i < len(d.PermissionIds); i++ {
		rolePermissions = append(rolePermissions, &model.RolePermission{
			RoleId:       d.RoleId,
			PermissionId: d.PermissionIds[i],
		})
	}

	role, err := model.GetRoleByRoleId(db, d.RoleId)
	if err != nil {
		logs.WithError(err).Error("get role by role id failed")
		return &common.SystemError{Code: common.DbError, Msg: "get role by role id failed", Err: err}
	}

	permissions, err := model.GetPermissionByPermissionIds(db, d.PermissionIds)
	if err != nil {
		logs.WithError(err).Error("get permission by permission ids failed")
		return &common.SystemError{Code: common.DbError, Msg: "get permission by permission ids failed", Err: err}
	}

	err = db.Transaction(func(tx *gorm2.DB) error {
		if err := model.CreateRolePermissionBatch(tx, rolePermissions); err != nil {
			return err
		}

		// add role permission to casbin
		err := e.GetAdapter().(*gormadapter.Adapter).Transaction(e, func(e casbin2.IEnforcer) error {
			for i := 0; i < len(permissions); i++ {
				permission := permissions[i]
				_, err := e.AddPermissionForUser(role.RoleCode, permission.Action)
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
		logs.WithError(err).Error("add role permission failed")
		return &common.SystemError{Code: common.DbError, Msg: "add role permission failed", Err: err}
	}

	return nil
}

func (s *RolePermissionService) DeleteRolePermission(ctx *gin.Context, d *define.DeleteRolePermissionRequest) error {
	logs := logrus.WithContext(ctx)
	db := gorm.GetDB().WithContext(ctx)
	e := casbin.GetEnforcer()

	role, err := model.GetRoleByRoleId(db, d.RoleId)
	if err != nil {
		logs.WithError(err).Error("get role by role id failed")
		return &common.SystemError{Code: common.DbError, Msg: "get role by role id failed", Err: err}
	}

	permissions, err := model.GetPermissionByPermissionIds(db, d.PermissionIds)
	if err != nil {
		logs.WithError(err).Error("get permission by permission ids failed")
		return &common.SystemError{Code: common.DbError, Msg: "get permission by permission ids failed", Err: err}
	}

	err = db.Transaction(func(tx *gorm2.DB) error {
		if err := model.DeleteRolePermissionBatch(tx, d.RoleId, d.PermissionIds); err != nil {
			return err
		}

		// delete role permission from casbin
		err := e.GetAdapter().(*gormadapter.Adapter).Transaction(e, func(e casbin2.IEnforcer) error {
			for i := 0; i < len(permissions); i++ {
				permission := permissions[i]
				_, err := e.DeletePermissionForUser(role.RoleCode, permission.Action)
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
		logs.WithError(err).Error("delete role permission failed")
		return &common.SystemError{Code: common.DbError, Msg: "delete role permission failed", Err: err}
	}

	return nil
}

func (s *RolePermissionService) ListPermissionIdsByRoleId(ctx *gin.Context, roleId uint) ([]uint, error) {
	logs := logrus.WithContext(ctx)
	db := gorm.GetDB().WithContext(ctx)

	permissionIds, err := model.ListPermissionIdsByRoleId(db, roleId)
	if err != nil {
		logs.WithError(err).Error("list permission ids by role id failed")
		return nil, &common.SystemError{Code: common.DbError, Msg: "list permission ids by role id failed", Err: err}
	}

	return permissionIds, nil
}
