package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pfjhyyj/ether/app/auth/define"
	"github.com/pfjhyyj/ether/app/auth/service"
	"github.com/pfjhyyj/ether/common"
)

type LoginController struct {
	service *service.LoginService
}

func NewLoginController(service *service.LoginService) *LoginController {
	return &LoginController{service: service}
}

// LoginByUsername godoc
// @Summary Login by username
// @Description Login by username
// @Tags auth
// @Accept json
// @Produce json
// @Param request body define.LoginByUsernameRequest true "LoginByUsernameRequest"
// @Success 200 {object} define.TokenResponse
// @Router /auth/loginByUsername [post]
func (r *LoginController) LoginByUsername(ctx *gin.Context) {
	var req define.LoginByUsernameRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		_ = ctx.Error(err)
		return
	}

	loginToken, err := r.service.LoginByUsername(ctx, req.Username, req.Password)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	resp := &define.TokenResponse{
		AccessToken: loginToken.Token,
		ExpireTime:  loginToken.ExpireTime,
	}
	ctx.JSON(200, &common.Response{
		Code: common.Ok,
		Data: resp,
	})
}
