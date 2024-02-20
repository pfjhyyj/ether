package service

import (
	"github.com/gin-gonic/gin"
	"github.com/pfjhyyj/ether/app/user/model"
	"github.com/pfjhyyj/ether/clients/gorm"
	"github.com/pfjhyyj/ether/common"
	"github.com/sirupsen/logrus"
)

type MenuService struct{}

func NewMenuService() *MenuService {
	return &MenuService{}
}

func (s *MenuService) CreateMenu(ctx *gin.Context, menu *model.Menu) error {
	logs := logrus.WithContext(ctx)
	db := gorm.GetDB().WithContext(ctx)

	if err := model.CreateMenu(db, menu); err != nil {
		logs.WithError(err).Error("create menu failed")
		return &common.SystemError{
			Code: common.DbError,
			Msg:  "create menu failed",
			Err:  err,
		}
	}

	return nil
}

func (s *MenuService) UpdateMenu(ctx *gin.Context, menuId uint, menu *model.Menu) error {
	logs := logrus.WithContext(ctx)
	db := gorm.GetDB().WithContext(ctx)

	if err := model.UpdateMenu(db, menuId, menu); err != nil {
		logs.WithError(err).Error("update menu failed")
		return &common.SystemError{
			Code: common.DbError,
			Msg:  "update menu failed",
			Err:  err,
		}
	}

	return nil
}

func (s *MenuService) DeleteMenu(ctx *gin.Context, menuId uint) error {
	logs := logrus.WithContext(ctx)
	db := gorm.GetDB().WithContext(ctx)

	menus, err := model.ListMenuTreeByMenuId(db, menuId)
	if err != nil {
		logs.WithError(err).Error("get menu tree failed")
		return &common.SystemError{
			Code: common.DbError,
			Msg:  "get menu tree failed",
			Err:  err,
		}
	}

	menuIds := make([]uint, 0, len(menus))
	for _, menu := range menus {
		menuIds = append(menuIds, menu.MenuId)
	}
	if err := model.DeleteMenus(db, menuIds); err != nil {
		logs.WithError(err).Error("delete menu failed")
		return &common.SystemError{
			Code: common.DbError,
			Msg:  "delete menu failed",
			Err:  err,
		}
	}

	return nil
}

func (s *MenuService) GetMenuByMenuId(ctx *gin.Context, menuId uint) (*model.Menu, error) {
	logs := logrus.WithContext(ctx)
	db := gorm.GetDB().WithContext(ctx)

	menu, err := model.GetMenuByMenuId(db, menuId)
	if err != nil {
		logs.WithError(err).Error("get menu by menu id failed")
		return nil, &common.SystemError{
			Code: common.DbError,
			Msg:  "get menu by menu id failed",
			Err:  err,
		}
	}

	return menu, nil
}

func (s *MenuService) GetMenuTreeByMenuId(ctx *gin.Context, menuId uint) ([]*model.Menu, error) {
	logs := logrus.WithContext(ctx)
	db := gorm.GetDB().WithContext(ctx)

	menus, err := model.ListMenuTreeByMenuId(db, menuId)
	if err != nil {
		logs.WithError(err).Error("get menu tree failed")
		return nil, &common.SystemError{
			Code: common.DbError,
			Msg:  "get menu tree failed",
			Err:  err,
		}
	}

	return menus, nil
}

func (s *MenuService) ListMenus(ctx *gin.Context, params *model.QueryMenuParams) ([]*model.Menu, int64, error) {
	logs := logrus.WithContext(ctx)
	db := gorm.GetDB().WithContext(ctx)

	menus, total, err := model.ListMenus(db, params)
	if err != nil {
		logs.WithError(err).Error("list menus failed")
		return nil, 0, &common.SystemError{
			Code: common.DbError,
			Msg:  "list menus failed",
			Err:  err,
		}
	}

	return menus, total, nil
}

func (s *MenuService) ListAllMenus(ctx *gin.Context) ([]*model.Menu, error) {
	logs := logrus.WithContext(ctx)
	db := gorm.GetDB().WithContext(ctx)

	menus, err := model.ListAllMenus(db)
	if err != nil {
		logs.WithError(err).Error("list all menus failed")
		return nil, &common.SystemError{
			Code: common.DbError,
			Msg:  "list all menus failed",
			Err:  err,
		}
	}

	return menus, nil
}
