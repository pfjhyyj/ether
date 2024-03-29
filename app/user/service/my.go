package service

import (
	"github.com/gin-gonic/gin"
	utils2 "github.com/pfjhyyj/ether/app/auth/utils"
	"github.com/pfjhyyj/ether/app/user/define"
	"github.com/pfjhyyj/ether/app/user/model"
	"github.com/pfjhyyj/ether/app/user/utils"
	"github.com/pfjhyyj/ether/client/gorm"
	"github.com/pfjhyyj/ether/common"
	"github.com/sirupsen/logrus"
)

type MyService struct {
}

func NewMyService() *MyService {
	return &MyService{}
}

func (s *MyService) GetUserById(ctx *gin.Context, userId uint) (*model.User, error) {
	logs := logrus.WithContext(ctx)
	db := gorm.GetDB().WithContext(ctx)

	user, err := model.GetUserByUserId(db, userId)
	if err != nil {
		logs.WithError(err).Error("get user by user id failed")
		return nil, &common.SystemError{Code: common.DbError, Msg: "get user by user id failed", Err: err}
	}
	return user, nil
}

func (s *MyService) UpdateMyInfo(ctx *gin.Context, userId uint, d *define.UpdateMyInfoRequest) error {
	logs := logrus.WithContext(ctx)
	db := gorm.GetDB().WithContext(ctx)

	user, err := model.GetUserByUserId(db, userId)
	if err != nil {
		logs.WithError(err).Error("get user by user id failed")
		return &common.SystemError{Code: common.DbError, Msg: "get user by user id failed", Err: err}
	}
	if user == nil {
		logs.Error("user not found")
		return &common.SystemError{Code: common.RequestError, Msg: "user not found"}
	}

	newUser := utils.ConvertUpdateMyInfoRequestToModel(d)

	if err := model.UpdateUser(db, userId, newUser); err != nil {
		logs.WithError(err).Error("update user failed")
		return &common.SystemError{Code: common.DbError, Msg: "update user failed", Err: err}
	}

	return nil
}

func (s *MyService) UpdateMyPassword(ctx *gin.Context, userId uint, d *define.UpdateMyPasswordRequest) error {
	logs := logrus.WithContext(ctx)
	db := gorm.GetDB().WithContext(ctx)

	user, err := model.GetUserByUserId(db, userId)
	if err != nil {
		logs.WithError(err).Error("get user by user id failed")
		return &common.SystemError{Code: common.DbError, Msg: "get user by user id failed", Err: err}
	}
	if user == nil {
		logs.Error("user not found")
		return &common.SystemError{Code: common.RequestError, Msg: "user not found"}
	}

	err = utils2.CompareHashAndPassword(user.Password, d.OldPassword)
	if err != nil {
		logs.Error("old password not match")
		return &common.SystemError{Code: common.RequestError, Msg: "old password not match"}
	}

	if d.NewPassword != d.RepeatPassword {
		logs.Error("new password not match")
		return &common.SystemError{Code: common.RequestError, Msg: "new password not match"}
	}

	newUser := utils.ConvertUpdateMyPasswordRequestToModel(d)

	if err := model.UpdateUser(db, userId, newUser); err != nil {
		logs.WithError(err).Error("update user failed")
		return &common.SystemError{Code: common.DbError, Msg: "update user failed", Err: err}
	}

	return nil
}

func (s *MyService) GetUserMenu(ctx *gin.Context, userId uint) ([]*model.Menu, error) {
	logs := logrus.WithContext(ctx)
	db := gorm.GetDB().WithContext(ctx)

	roleIds, err := model.ListUserRoleIdsByUserId(db, userId)
	if err != nil {
		logs.WithError(err).Error("get user roles failed")
		return nil, &common.SystemError{
			Code: common.DbError,
			Msg:  "get user roles failed",
			Err:  err,
		}
	}

	if len(roleIds) == 0 {
		return nil, nil
	}

	menuIds, err := model.ListMenuIdsByRoleIds(db, roleIds)
	if err != nil {
		logs.WithError(err).Error("get role's menus failed")
		return nil, &common.SystemError{
			Code: common.DbError,
			Msg:  "get role's menus failed",
			Err:  err,
		}
	}

	menus, err := model.ListMenuTreeFromBottomByMenuIds(db, menuIds)
	if err != nil {
		logs.WithError(err).Error("get menus failed")
		return nil, &common.SystemError{
			Code: common.DbError,
			Msg:  "get menus failed",
			Err:  err,
		}
	}

	return menus, nil
}
