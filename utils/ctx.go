package utils

import "context"

func GetUserIdFromCtx(ctx context.Context) (uint, bool) {
	userId, ok := ctx.Value("userId").(uint)
	if !ok {
		return 0, false
	}
	return userId, true
}

func SetUserIdToCtx(ctx context.Context, userId uint) context.Context {
	return context.WithValue(ctx, "userId", userId)
}
