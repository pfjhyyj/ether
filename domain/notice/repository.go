package notice

import "context"

type Repository interface {
	NotifyUsers(ctx context.Context, message *Message, userIds []uint) error
}
