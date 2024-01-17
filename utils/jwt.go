package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pfjhyyj/ether/common"
	"github.com/spf13/viper"
	"strconv"
	"time"
)

type TokenPayload struct {
	UserId uint
}

type TokenClaims struct {
	UserId uint
	jwt.RegisteredClaims
}

func GenerateAccessToken(userId uint) (string, *time.Time, error) {
	secret := viper.GetString("jwt.secret")

	expiredAt := time.Now().Add(common.TokenExpireTime * time.Minute)
	claims := TokenClaims{
		userId,
		jwt.RegisteredClaims{
			Subject:   strconv.FormatUint(uint64(userId), 10),
			ExpiresAt: jwt.NewNumericDate(expiredAt),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	ss, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", nil, err
	}
	return ss, &expiredAt, nil
}

func ParseToken(tokenString string) (*TokenPayload, error) {
	secret := viper.GetString("jwt.secret")
	token, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*TokenClaims)
	if !ok {
		return nil, errors.New("invalid token")
	}

	return &TokenPayload{
		UserId: claims.UserId,
	}, nil
}
