package repository

import (
	"errors"
	"github.com/pfjhyyj/ether/app/user/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) UpdateUser(userId uint, user *model.User) error {
	return r.db.Where("user_id = ?", userId).Updates(user).Error
}

func (r *UserRepository) DeleteUser(userId uint) error {
	return r.db.Delete(&model.User{}, "user_id = ?", userId).Error
}

func (r *UserRepository) GetUserByUserId(userId uint) (*model.User, error) {
	var user model.User
	if err := r.db.First(&user, "user_id = ?", userId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetUserByUsername(username string) (*model.User, error) {
	var user model.User
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) ListUsers(params *model.QueryUserParams) ([]*model.User, int64, error) {
	var users []*model.User
	query := r.db.Model(&model.User{})

	var total int64
	query.Count(&total)

	if params.Current > 0 && params.PageSize > 0 {
		query = query.Offset((params.Current - 1) * params.PageSize).Limit(params.PageSize)
	}

	if err := query.Find(&users).Error; err != nil {
		return nil, 0, err
	}
	return users, total, nil
}
