package service

import (
	"github.com/gin-gonic/gin"
	"github.com/pfjhyyj/ether/clients/redis"
	"github.com/pfjhyyj/ether/common"
	"github.com/pfjhyyj/ether/utils"
	"strconv"
)

type LogoutService struct {
}

func NewLogoutService() *LogoutService {
	return &LogoutService{}
}

func (s LogoutService) Logout(ctx *gin.Context) error {
	userId, ok := utils.GetUserIdFromCtx(ctx.Request.Context())
	if !ok {
		return nil
	}

	redisClient := redis.GetRedisClient()
	key := common.TokenPrefix + strconv.FormatUint(uint64(userId), 10)
	_, _ = redisClient.Del(ctx, key).Result()
	return nil
}
