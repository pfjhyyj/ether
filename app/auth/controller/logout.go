package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pfjhyyj/ether/app/auth/service"
	"github.com/pfjhyyj/ether/common"
)

type LogoutController struct {
	service *service.LogoutService
}

func NewLogoutController(service *service.LogoutService) *LogoutController {
	return &LogoutController{service: service}
}

// Logout godoc
// @Summary Logout
// @Description Logout
// @Tags auth
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} common.Response
// @Router /auth/logout [post]
func (r *LogoutController) Logout(ctx *gin.Context) {
	err := r.service.Logout(ctx)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(200, &common.Response{
		Code: common.Ok,
	})
}
