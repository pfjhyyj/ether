package service

import (
	"github.com/gin-gonic/gin"
	"github.com/pfjhyyj/ether/app/user/model"
	"github.com/pfjhyyj/ether/client/gorm"
	"github.com/pfjhyyj/ether/common"
	"github.com/sirupsen/logrus"
)

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (s UserService) ListUsers(ctx *gin.Context, param *model.QueryUserParams) ([]*model.User, int64, error) {
	logs := logrus.WithContext(ctx)
	db := gorm.GetDB().WithContext(ctx)

	users, total, err := model.ListUsers(db, param)
	if err != nil {
		logs.WithError(err).Error("list users failed")
		return nil, 0, &common.SystemError{Code: common.DbError, Msg: "list users failed", Err: err}
	}

	return users, total, nil
}

func (s UserService) GetUserById(ctx *gin.Context, userId uint) (*model.User, error) {
	logs := logrus.WithContext(ctx)
	db := gorm.GetDB().WithContext(ctx)

	user, err := model.GetUserByUserId(db, userId)
	if err != nil {
		logs.WithError(err).Error("get user by id failed")
		return nil, &common.SystemError{Code: common.DbError, Msg: "get user by id failed", Err: err}
	}

	return user, nil
}
