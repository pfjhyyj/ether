package common

import "fmt"

const (
	TokenPrefix     = "token:"
	TokenExpireTime = 3600 * 24 * 7
)

func GetTokenKey(userId uint) string {
	return fmt.Sprintf("%s%d", TokenPrefix, userId)
}
