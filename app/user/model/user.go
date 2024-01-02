package model

import (
	"github.com/pfjhyyj/ether/common"
)

type User struct {
	common.Model
	UserId   uint   `gorm:"primaryKey"`
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
	Salt     string `gorm:"column:salt"`
	Email    string `gorm:"column:email"`
	Mobile   string `gorm:"column:mobile"`
}

func (User) TableName() string {
	return "user"
}

type QueryUserParams struct {
	common.PageRequest
}
