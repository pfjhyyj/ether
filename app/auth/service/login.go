package service

import (
	"github.com/gin-gonic/gin"
	"github.com/pfjhyyj/ether/app/auth/utils"
	"github.com/pfjhyyj/ether/app/user/constants"
	"github.com/pfjhyyj/ether/client/redis"
	"github.com/pfjhyyj/ether/common"
	"github.com/pfjhyyj/ether/domain/user"
	utils2 "github.com/pfjhyyj/ether/utils"
	"time"
)

type LoginService struct {
	userRepo user.Repository
}

func NewLoginService(userRepo user.Repository) *LoginService {
	return &LoginService{userRepo: userRepo}
}

type LoginToken struct {
	Token      string
	ExpireTime int64
}

func (s *LoginService) LoginByUsername(ctx *gin.Context, username string, password string) (*LoginToken, error) {
	// get user by username
	u, err := s.userRepo.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, &common.SystemError{Code: common.RequestError, Msg: "invalid username or password", Err: err}
	}

	if u.Status != constants.UserStatusEnabled {
		return nil, &common.SystemError{Code: common.RequestError, Msg: "user is disabled", Err: nil}
	}

	// compare password
	err = utils.CompareHashAndPassword(u.Password, password)
	if err != nil {
		return nil, &common.SystemError{Code: common.RequestError, Msg: "invalid username or password", Err: err}
	}

	// generate authToken and set to redis
	redisClient := redis.GetRedisClient()
	token, expiredAt, err := utils2.GenerateAccessToken(u.UserId)
	if err != nil {
		return nil, &common.SystemError{Code: common.UnknownError, Msg: "generate access authToken fail", Err: err}
	}
	key := common.GetTokenKey(u.UserId)
	redisClient.Set(ctx, key, token, common.TokenExpireTime*time.Minute)

	authToken := &LoginToken{
		Token:      token,
		ExpireTime: expiredAt.Unix(),
	}

	return authToken, nil
}
