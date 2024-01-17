package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pfjhyyj/ether/app/auth/service"
	"github.com/pfjhyyj/ether/common"
)

type LoginController struct {
	service *service.LoginService
}

func NewLoginController(service *service.LoginService) *LoginController {
	return &LoginController{service: service}
}

type LoginByUsernameRequest struct {
	Username string `json:"username,min=6,max=20"`
	Password string `json:"password,min=8,max=20"`
}

type TokenResponse struct {
	AccessToken string `json:"accessToken"`
	ExpireTime  int64  `json:"expireTime"`
}

func (r *LoginController) LoginByUsername(ctx *gin.Context) {
	var req LoginByUsernameRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		_ = ctx.Error(err)
		return
	}

	loginToken, err := r.service.LoginByUsername(ctx, req.Username, req.Password)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	resp := TokenResponse{
		AccessToken: loginToken.Token,
		ExpireTime:  loginToken.ExpireTime,
	}
	ctx.JSON(200, &common.Response{
		Code: common.Ok,
		Data: resp,
	})
}
