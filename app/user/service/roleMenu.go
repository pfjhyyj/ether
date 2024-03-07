package service

import (
	"github.com/gin-gonic/gin"
	"github.com/pfjhyyj/ether/app/user/define"
	"github.com/pfjhyyj/ether/app/user/model"
	"github.com/pfjhyyj/ether/client/gorm"
	"github.com/pfjhyyj/ether/common"
	"github.com/sirupsen/logrus"
	gorm2 "gorm.io/gorm"
)

type RoleMenuService struct{}

func NewRoleMenuService() *RoleMenuService {
	return &RoleMenuService{}
}

func (s *RoleMenuService) AddRoleMenu(ctx *gin.Context, d *define.AddRoleMenuRequest) error {
	logs := logrus.WithContext(ctx)
	db := gorm.GetDB().WithContext(ctx)

	var roleMenus []*model.RoleMenu
	for i := 0; i < len(d.MenuIds); i++ {
		roleMenus = append(roleMenus, &model.RoleMenu{
			RoleId: d.RoleId,
			MenuId: d.MenuIds[i],
		})
	}

	if err := model.CreateRoleMenuBatch(db, roleMenus); err != nil {
		logs.WithError(err).Error("create role menu failed")
		return &common.SystemError{Code: common.DbError, Msg: "create role menu failed", Err: err}
	}

	return nil
}

func (s *RoleMenuService) DeleteRoleMenu(ctx *gin.Context, d *define.DeleteRoleMenuRequest) error {
	logs := logrus.WithContext(ctx)
	db := gorm.GetDB().WithContext(ctx)

	if err := model.DeleteRoleMenu(db, d.RoleId, d.MenuIds); err != nil {
		logs.WithError(err).Error("delete role menu failed")
		return &common.SystemError{Code: common.DbError, Msg: "delete role menu failed", Err: err}
	}

	return nil
}

func (s *RoleMenuService) SetRoleMenu(ctx *gin.Context, d *define.SetRoleMenuRequest) error {
	logs := logrus.WithContext(ctx)
	db := gorm.GetDB().WithContext(ctx)

	err := db.Transaction(func(tx *gorm2.DB) error {
		// delete all role menu
		if err := model.DeleteRoleMenu(db, d.RoleId, nil); err != nil {
			logs.WithError(err).Error("delete role menu failed")
			return &common.SystemError{Code: common.DbError, Msg: "delete role menu failed", Err: err}
		}

		// create role menu
		var roleMenus []*model.RoleMenu
		for i := 0; i < len(d.MenuIds); i++ {
			roleMenus = append(roleMenus, &model.RoleMenu{
				RoleId: d.RoleId,
				MenuId: d.MenuIds[i],
			})
		}
		if err := model.CreateRoleMenuBatch(db, roleMenus); err != nil {
			logs.WithError(err).Error("create role menu failed")
			return &common.SystemError{Code: common.DbError, Msg: "create role menu failed", Err: err}
		}
		return nil
	})

	return err
}

func (s *RoleMenuService) ListMenuIdsByRoleId(ctx *gin.Context, roleId uint) ([]uint, error) {
	logs := logrus.WithContext(ctx)
	db := gorm.GetDB().WithContext(ctx)

	menuIds, err := model.ListMenuIdsByRoleId(db, roleId)
	if err != nil {
		logs.WithError(err).Error("list menu ids by role id failed")
		return nil, &common.SystemError{Code: common.DbError, Msg: "list menu ids by role id failed", Err: err}
	}

	return menuIds, nil
}
