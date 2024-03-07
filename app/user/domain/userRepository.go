package domain

import (
	"context"
	"errors"
	"github.com/mitchellh/mapstructure"
	"github.com/pfjhyyj/ether/app/user/model"
	"github.com/pfjhyyj/ether/client/gorm"
	"github.com/pfjhyyj/ether/domain/user"
)

type UserRepository struct {
	user.Repository
}

func GetUserRepository() *UserRepository {
	Init()
	return userRepository
}

func (r *UserRepository) CreateUser(ctx context.Context, user *user.User) error {
	var userModel model.User
	err := mapstructure.Decode(user, &userModel)
	if err != nil {
		return err
	}

	db := gorm.GetDB().WithContext(ctx)
	err = model.CreateUser(db, &userModel)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetUserByUsername(ctx context.Context, username string) (*user.User, error) {
	db := gorm.GetDB().WithContext(ctx)
	userModel, err := model.GetUserByUsername(db, username)
	if err != nil {
		return nil, err
	}
	if userModel == nil {
		return nil, errors.New("user not found")
	}

	var u user.User
	err = mapstructure.Decode(userModel, &u)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*user.User, error) {
	db := gorm.GetDB().WithContext(ctx)
	userModel, err := model.GetUserByEmail(db, email)
	if err != nil {
		return nil, err
	}
	if userModel == nil {
		return nil, errors.New("user not found")
	}

	var u user.User
	err = mapstructure.Decode(userModel, &u)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *UserRepository) GetUserByUserId(ctx context.Context, userId uint) (*user.User, error) {
	db := gorm.GetDB().WithContext(ctx)
	userModel, err := model.GetUserByUserId(db, userId)
	if err != nil {
		return nil, err
	}
	if userModel == nil {
		return nil, errors.New("user not found")
	}

	var u user.User
	err = mapstructure.Decode(userModel, &u)
	if err != nil {
		return nil, err
	}
	return &u, nil
}
