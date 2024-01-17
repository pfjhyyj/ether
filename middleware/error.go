package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/pfjhyyj/ether/common"
	"github.com/sirupsen/logrus"
	"net/http"
)

func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		logs := logrus.WithContext(c)
		c.Next()

		for _, err := range c.Errors {
			var systemErr *common.SystemError
			if errors.As(err.Err, &systemErr) {
				logs.WithError(systemErr.Err).Error(systemErr.Message)
				c.AbortWithStatusJSON(http.StatusOK, &common.Response{
					Code: systemErr.Code,
					Msg:  systemErr.Message,
				})
				return
			} else {
				logs.WithError(err).Error("unknown error")
				c.AbortWithStatusJSON(http.StatusOK, &common.Response{
					Code: common.UnknownError,
					Msg:  "unknown error",
				})
				return
			}
		}
	}
}
