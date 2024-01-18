package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pfjhyyj/ether/app/auth/service"
	"github.com/pfjhyyj/ether/common"
	"net/http"
)

type RegisterController struct {
	service *service.RegisterService
}

func NewRegisterController(service *service.RegisterService) *RegisterController {
	return &RegisterController{service: service}
}

type RegisterUserRequest struct {
	Username string `json:"username" binding:"required,min=6,max=20"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,max=20"`
}

func (r *RegisterController) RegisterByEmail(ctx *gin.Context) {
	var req RegisterUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		_ = ctx.Error(err)
		return
	}

	err := r.service.RegisterUserByEmail(ctx, req.Username, req.Email, req.Password)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, common.Response{
		Code: common.Ok,
	})
}
