package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pfjhyyj/ether/app/user/define"
	"github.com/pfjhyyj/ether/app/user/service"
	"github.com/pfjhyyj/ether/app/user/utils"
	"github.com/pfjhyyj/ether/common"
	"net/http"
)

type UserController struct {
	service *service.UserService
}

func NewUserController(service *service.UserService) *UserController {
	return &UserController{service: service}
}

func (r *UserController) ListUsers(ctx *gin.Context) {
	var req define.ListUserRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		_ = ctx.Error(err)
		return
	}

	queryParam := utils.ConvertUserListPageRequestToParam(&req)

	users, total, err := r.service.ListUsers(ctx, queryParam)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	userInfo := utils.ConvertUserListToPageResponse(users)

	ctx.JSON(http.StatusOK, &common.Response{
		Code: common.Ok,
		Data: &common.Page{
			Total:    total,
			PageSize: req.PageSize,
			Current:  req.Current,
			List:     userInfo,
		},
	})
}
