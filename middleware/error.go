package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/pfjhyyj/ether/common"
	"github.com/pfjhyyj/ether/utils"
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
				logs.WithError(systemErr.Err).Error(systemErr.Msg)
				c.AbortWithStatusJSON(http.StatusOK, &common.Response{
					Code: systemErr.Code,
					Msg:  systemErr.Msg,
				})
				return
			} else if errors.As(err.Err, &validator.ValidationErrors{}) {
				logs.WithError(err).Error("request error")
				validationErrs := err.Err.(validator.ValidationErrors).Translate(utils.GetValidatorTrans())
				var errMsg string
				for _, v := range validationErrs {
					errMsg += v + ";"
				}
				c.AbortWithStatusJSON(http.StatusOK, &common.Response{
					Code: common.RequestError,
					Msg:  errMsg,
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
