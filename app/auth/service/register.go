package service

import (
	"github.com/gin-gonic/gin"
	"github.com/pfjhyyj/ether/app/auth/utils"
	"github.com/pfjhyyj/ether/common"
	"github.com/pfjhyyj/ether/domain/user"
)

type RegisterService struct {
	userRepo user.Repository
}

func NewRegisterService(userRepo user.Repository) *RegisterService {
	return &RegisterService{userRepo: userRepo}
}

func (s *RegisterService) RegisterUserByEmail(ctx *gin.Context, username string, email string, password string) error {
	hashPassword, err := utils.HashPassword(password)
	if err != nil {
		return &common.SystemError{Code: common.UnknownError, Msg: "hash password fail", Err: err}
	}

	newUser := &user.User{
		Username: username,
		Email:    email,
		Password: hashPassword,
	}

	if err := s.userRepo.CreateUser(ctx, newUser); err != nil {
		return &common.SystemError{Code: common.DbError, Msg: "create newUser fail", Err: err}
	}

	return nil
}
