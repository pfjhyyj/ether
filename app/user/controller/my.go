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

// MyInfo godoc
// @Summary Get my info
// @Description Get my info
// @Tags my
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} common.Response{data=define.MyInfoResponse}
// @Router /my [get]
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

// UpdateMyInfo godoc
// @Summary Update my info
// @Description Update my info
// @Tags my
// @Accept json
// @Produce json
// @Security Bearer
// @Param request body define.UpdateMyInfoRequest true "UpdateMyInfoRequest"
// @Success 200 {object} common.Response
// @Router /my [put]
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

// UpdateMyPassword godoc
// @Summary Update my password
// @Description Update my password
// @Tags my
// @Accept json
// @Produce json
// @Security Bearer
// @Param request body define.UpdateMyPasswordRequest true "UpdateMyPasswordRequest"
// @Success 200 {object} common.Response
// @Router /my/password [put]
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

// GetMyMenu godoc
// @Summary Get my menu
// @Description Get my menu
// @Tags my
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} common.Response{data=define.GetMenuResponse}
// @Router /my/menus [get]
func (c *MyController) GetMyMenu(ctx *gin.Context) {
	userId := ctx.GetUint(common.CtxUserIDKey)

	menus, err := c.service.GetUserMenu(ctx, userId)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	menusInfo := utils.ConvertMyMenuToResponse(menus)
	ctx.JSON(http.StatusOK, &common.Response{
		Code: common.Ok,
		Data: menusInfo,
	})
}
