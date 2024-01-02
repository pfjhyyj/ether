package user

import "github.com/pfjhyyj/ether/app/user/model"

type Repository interface {
	GetUserByUsername(username string) (*model.User, error)
	GetUserByPhone(phoneNum string) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	GetUserById(userId uint) (*model.User, error)
}
