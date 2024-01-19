package service

import (
	"context"
	utils2 "github.com/pfjhyyj/ether/app/auth/utils"
	"github.com/pfjhyyj/ether/app/user/define"
	"github.com/pfjhyyj/ether/app/user/model"
	"github.com/pfjhyyj/ether/app/user/utils"
	"github.com/pfjhyyj/ether/clients/gorm"
	"github.com/pfjhyyj/ether/common"
	"github.com/sirupsen/logrus"
)

type MyService struct {
}

func NewMyService() *MyService {
	return &MyService{}
}

func (s *MyService) GetUserById(ctx context.Context, userId uint) (*model.User, error) {
	logs := logrus.WithContext(ctx)
	db := gorm.GetDB().WithContext(ctx)

	user, err := model.GetUserByUserId(db, userId)
	if err != nil {
		logs.WithError(err).Error("get user by user id failed")
		return nil, &common.SystemError{Code: common.DbError, Message: "get user by user id failed", Err: err}
	}
	return user, nil
}

func (s *MyService) UpdateMyInfo(ctx context.Context, userId uint, d *define.UpdateMyInfoRequest) error {
	logs := logrus.WithContext(ctx)
	db := gorm.GetDB().WithContext(ctx)

	user, err := model.GetUserByUserId(db, userId)
	if err != nil {
		logs.WithError(err).Error("get user by user id failed")
		return &common.SystemError{Code: common.DbError, Message: "get user by user id failed", Err: err}
	}
	if user == nil {
		logs.Error("user not found")
		return &common.SystemError{Code: common.RequestError, Message: "user not found"}
	}

	newUser := utils.ConvertUpdateMyInfoRequestToModel(d)

	if err := model.UpdateUser(db, userId, newUser); err != nil {
		logs.WithError(err).Error("update user failed")
		return &common.SystemError{Code: common.DbError, Message: "update user failed", Err: err}
	}

	return nil
}

func (s *MyService) UpdateMyPassword(ctx context.Context, userId uint, d *define.UpdateMyPasswordRequest) error {
	logs := logrus.WithContext(ctx)
	db := gorm.GetDB().WithContext(ctx)

	user, err := model.GetUserByUserId(db, userId)
	if err != nil {
		logs.WithError(err).Error("get user by user id failed")
		return &common.SystemError{Code: common.DbError, Message: "get user by user id failed", Err: err}
	}
	if user == nil {
		logs.Error("user not found")
		return &common.SystemError{Code: common.RequestError, Message: "user not found"}
	}

	err = utils2.CompareHashAndPassword(user.Password, d.OldPassword)
	if err != nil {
		logs.Error("old password not match")
		return &common.SystemError{Code: common.RequestError, Message: "old password not match"}
	}

	newUser := utils.ConvertUpdateMyPasswordRequestToModel(d)

	if err := model.UpdateUser(db, userId, newUser); err != nil {
		logs.WithError(err).Error("update user failed")
		return &common.SystemError{Code: common.DbError, Message: "update user failed", Err: err}
	}

	return nil
}
