package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pfjhyyj/ether/clients/casbin"
	"github.com/sirupsen/logrus"
)

func CheckPermission(ctx *gin.Context, ob string, act string) bool {
	logs := logrus.WithContext(ctx)
	e := casbin.GetEnforcer()

	userId := ctx.GetUint("userId")
	userIdStr := fmt.Sprintf("%d", userId)

	ok, err := e.Enforce(userIdStr, ob, act)

	if err != nil {
		logs.WithError(err).Error("check permission failed")
		return false
	}

	return ok
}
