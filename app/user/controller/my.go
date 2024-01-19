package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pfjhyyj/ether/app/user/define"
	"github.com/pfjhyyj/ether/app/user/service"
	"github.com/pfjhyyj/ether/app/user/utils"
	"github.com/pfjhyyj/ether/common"
	"net/http"
)

type MyController struct {
	service *service.MyService
}

func NewMyController(service *service.MyService) *MyController {
	return &MyController{service: service}
}

func (c *MyController) MyInfo(ctx *gin.Context) {
	userId := ctx.GetUint(common.CtxUserIDKey)

	user, err := c.service.GetUserById(ctx, userId)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	userInfo := utils.ConvertUserToMyInfoResponse(user)
	ctx.JSON(http.StatusOK, &common.Response{
		Code: common.Ok,
		Data: userInfo,
	})
}

func (c *MyController) UpdateMyInfo(ctx *gin.Context) {
	userId := ctx.GetUint(common.CtxUserIDKey)

	var req define.UpdateMyInfoRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		_ = ctx.Error(err)
		return
	}

	err := c.service.UpdateMyInfo(ctx, userId, &req)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, &common.Response{
		Code: common.Ok,
	})
}

func (c *MyController) UpdateMyPassword(ctx *gin.Context) {
	userId := ctx.GetUint(common.CtxUserIDKey)

	var req define.UpdateMyPasswordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		_ = ctx.Error(err)
		return
	}

	err := c.service.UpdateMyPassword(ctx, userId, &req)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, &common.Response{
		Code: common.Ok,
	})
}
