package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pfjhyyj/ether/app/user/define"
	"github.com/pfjhyyj/ether/app/user/service"
	"github.com/pfjhyyj/ether/app/user/utils"
	"github.com/pfjhyyj/ether/common"
	utils2 "github.com/pfjhyyj/ether/utils"
	"net/http"
)

type UserController struct {
	service *service.UserService
}

func NewUserController(service *service.UserService) *UserController {
	return &UserController{service: service}
}

// ListUsers godoc
// @Summary List users
// @Description List users
// @Tags user
// @Accept json
// @Produce json
// @Security Bearer
// @Param request query define.ListUserRequest true "ListUserRequest"
// @Success 200 {object} common.Response{data=common.Page{list=[]define.ListUserPageResponse}}
// @Router /users [get]
func (r *UserController) ListUsers(ctx *gin.Context) {
	if ok := utils2.CheckPermission(ctx, "user", "list"); !ok {
		ctx.JSON(http.StatusOK, &common.Response{
			Code: common.NoPermissionError,
			Msg:  "no permission",
		})
		return
	}

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

// GetUser godoc
// @Summary Get user
// @Description Get user
// @Tags user
// @Accept json
// @Produce json
// @Security Bearer
// @Param request query define.GetUserRequest true "GetUserRequest"
// @Success 200 {object} common.Response{data=define.GetUserResponse}
// @Router /users/{userId} [get]
func (r *UserController) GetUser(ctx *gin.Context) {
	if ok := utils2.CheckPermission(ctx, "user", "get"); !ok {
		ctx.JSON(http.StatusOK, &common.Response{
			Code: common.NoPermissionError,
			Msg:  "no permission",
		})
		return
	}

	var req define.GetUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		_ = ctx.Error
		return
	}

	user, err := r.service.GetUserById(ctx, req.UserId)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	userInfo := utils.ConvertUserToGetUserResponse(user)

	ctx.JSON(http.StatusOK, &common.Response{
		Code: common.Ok,
		Data: userInfo,
	})
}
