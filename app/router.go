package app

import (
	"github.com/gin-gonic/gin"
	"github.com/pfjhyyj/ether/app/auth"
	"github.com/pfjhyyj/ether/app/notice"
	"github.com/pfjhyyj/ether/app/tenant"
	"github.com/pfjhyyj/ether/app/user"
)

func SetRouter(r *gin.RouterGroup) {
	user.SetRouter(r)
	tenant.SetRouter(r)
	auth.SetRouter(r)
	notice.SetRouter(r)
}
