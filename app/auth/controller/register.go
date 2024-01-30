package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pfjhyyj/ether/app/auth/define"
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

// RegisterByEmail godoc
// @Summary Register by email
// @Description Register by email
// @Tags auth
// @Accept json
// @Produce json
// @Param request body define.RegisterUserRequest true "RegisterUserRequest"
// @Success 200 {object} common.Response
// @Router /auth/registerByEmail [post]
func (r *RegisterController) RegisterByEmail(ctx *gin.Context) {
	var req define.RegisterUserRequest
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
