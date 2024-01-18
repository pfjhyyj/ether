package service

import (
	"github.com/gin-gonic/gin"
	"github.com/pfjhyyj/ether/clients/redis"
	"github.com/pfjhyyj/ether/common"
	"strconv"
)

type LogoutService struct {
}

func NewLogoutService() *LogoutService {
	return &LogoutService{}
}

func (s LogoutService) Logout(ctx *gin.Context) error {
	userId := ctx.GetUint(common.CtxUserIDKey)

	redisClient := redis.GetRedisClient()
	key := common.TokenPrefix + strconv.FormatUint(uint64(userId), 10)
	_, _ = redisClient.Del(ctx, key).Result()
	return nil
}
