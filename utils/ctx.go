package utils

import (
	"context"
	"github.com/pfjhyyj/ether/common"
)

func GetUserIdFromCtx(ctx context.Context) (uint, bool) {
	userId, ok := ctx.Value(common.CtxUserIDKey).(uint)
	if !ok {
		return 0, false
	}
	return userId, true
}

func SetUserIdToCtx(ctx context.Context, userId uint) context.Context {
	return context.WithValue(ctx, common.CtxUserIDKey, userId)
}
