package user

import "context"

type Repository interface {
	GetUserByUsername(ctx context.Context, username string) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	GetUserByUserId(ctx context.Context, userId uint) (*User, error)
	CreateUser(ctx context.Context, user *User) error
}
