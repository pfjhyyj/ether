package service

import (
	"github.com/gin-gonic/gin"
	"github.com/pfjhyyj/ether/client/redis"
	"github.com/pfjhyyj/ether/common"
)

type LogoutService struct {
}

func NewLogoutService() *LogoutService {
	return &LogoutService{}
}

func (s LogoutService) Logout(ctx *gin.Context) error {
	userId := ctx.GetUint(common.CtxUserIDKey)

	redisClient := redis.GetRedisClient()
	key := common.GetTokenKey(userId)
	_, _ = redisClient.Del(ctx, key).Result()
	return nil
}
