package middleware

import (
	"github.com/gin-gonic/gin"
	redisClient "github.com/pfjhyyj/ether/client/redis"
	"github.com/pfjhyyj/ether/common"
	"github.com/pfjhyyj/ether/utils"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		logs := logrus.WithContext(c)
		// check if authorization header exists
		authorization := c.GetHeader("Authorization")
		if authorization == "" {
			logs.Errorf("no authorization header")
			c.AbortWithStatusJSON(http.StatusOK, &common.Response{
				Code: common.AuthError,
				Msg:  "Please login first.",
			})
			return
		}

		// check if authorization header is jwt token and valid
		token := strings.TrimPrefix(authorization, "Bearer ")
		tokenPayload, err := utils.ParseToken(token)
		if err != nil {
			logs.WithError(err).Errorf("fail to parse token %s", token)
			c.AbortWithStatusJSON(http.StatusOK, &common.Response{
				Code: common.AuthError,
				Msg:  "unauthorized request",
			})
			return
		}

		// check if token is valid (in redis)
		redisConn := redisClient.GetRedisClient()
		key := common.GetTokenKey(tokenPayload.UserId)
		_, err = redisConn.Get(c, key).Result()
		if err != nil {
			logs.WithError(err).Errorf("fail to check token %s in redis", token)
			c.AbortWithStatusJSON(http.StatusOK, &common.Response{
				Code: common.AuthError,
				Msg:  "unauthorized request",
			})
			return
		}

		c.Set(common.CtxUserIDKey, tokenPayload.UserId)
		c.Next()
	}
}
