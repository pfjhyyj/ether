package model

import (
	"errors"

	"github.com/pfjhyyj/ether/common"
	"gorm.io/gorm"
)

type User struct {
	common.Model
	UserId   uint   `gorm:"primaryKey"`
	Username string `gorm:"column:username;unique"`
	Password string `gorm:"column:password"`
	Email    string `gorm:"column:email;unique"`
	Mobile   string `gorm:"column:mobile;unique"`
	Avatar   string `gorm:"column:avatar"`
}

func (User) TableName() string {
	return "user"
}

type QueryUserParams struct {
	common.PageRequest
}

func CreateUser(tx *gorm.DB, user *User) error {
	return tx.Create(user).Error
}

func UpdateUser(tx *gorm.DB, userId uint, user *User) error {
	return tx.Where("user_id = ?", userId).Updates(user).Error
}

func DeleteUser(tx *gorm.DB, userId uint) error {
	return tx.Delete(&User{}, "user_id = ?", userId).Error
}

func GetUserByUsername(tx *gorm.DB, username string) (*User, error) {
	var user User
	if err := tx.Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func GetUserByEmail(tx *gorm.DB, email string) (*User, error) {
	var user User
	if err := tx.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func GetUserByUserId(tx *gorm.DB, userId uint) (*User, error) {
	var user User
	if err := tx.First(&user, "user_id = ?", userId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func ListUsers(tx *gorm.DB, params *QueryUserParams) ([]*User, int64, error) {
	var users []*User
	query := tx.Model(&User{})

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
