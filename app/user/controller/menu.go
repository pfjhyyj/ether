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

type MenuController struct {
	MenuService *service.MenuService
}

func NewMenuController(service *service.MenuService) *MenuController {
	return &MenuController{
		MenuService: service,
	}
}

func (c *MenuController) CreateMenu(ctx *gin.Context) {
	if ok := utils2.CheckPermission(ctx, "menu", "create"); !ok {
		ctx.JSON(http.StatusForbidden, &common.Response{
			Code: common.NoPermissionError,
			Msg:  "no permission",
		})
		return
	}

	var req define.CreateMenuRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		_ = ctx.Error(err)
		return
	}

	menu := utils.ConvertCreateMenuRequestToMenu(&req)
	if err := c.MenuService.CreateMenu(ctx, menu); err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, &common.Response{
		Code: common.Ok,
	})
}

func (c *MenuController) UpdateMenu(ctx *gin.Context) {
	if ok := utils2.CheckPermission(ctx, "menu", "update"); !ok {
		ctx.JSON(http.StatusForbidden, &common.Response{
			Code: common.NoPermissionError,
			Msg:  "no permission",
		})
		return
	}

	var req define.UpdateMenuRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		_ = ctx.Error(err)
		return
	}

	menu := utils.ConvertUpdateMenuRequestToMenu(&req)
	if err := c.MenuService.UpdateMenu(ctx, req.MenuId, menu); err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, &common.Response{
		Code: common.Ok,
	})
}

func (c *MenuController) DeleteMenu(ctx *gin.Context) {
	if ok := utils2.CheckPermission(ctx, "menu", "delete"); !ok {
		ctx.JSON(http.StatusForbidden, &common.Response{
			Code: common.NoPermissionError,
			Msg:  "no permission",
		})
		return
	}

	var req define.DeleteMenuRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		_ = ctx.Error(err)
		return
	}

	if err := c.MenuService.DeleteMenu(ctx, req.MenuId); err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, &common.Response{
		Code: common.Ok,
	})
}

func (c *MenuController) GetMenuTree(ctx *gin.Context) {
	if ok := utils2.CheckPermission(ctx, "menu", "get"); !ok {
		ctx.JSON(http.StatusForbidden, &common.Response{
			Code: common.NoPermissionError,
			Msg:  "no permission",
		})
		return
	}

	var req define.GetMenuRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		_ = ctx.Error(err)
		return
	}

	menus, err := c.MenuService.GetMenuTreeByMenuId(ctx, req.MenuId)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, &common.Response{
		Code: common.Ok,
		Data: utils.ConvertMenuToResponse(menus),
	})
}

func (c *MenuController) ListMenus(ctx *gin.Context) {
	if ok := utils2.CheckPermission(ctx, "menu", "list"); !ok {
		ctx.JSON(http.StatusForbidden, &common.Response{
			Code: common.NoPermissionError,
			Msg:  "no permission",
		})
		return
	}

	var req define.ListMenusRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		_ = ctx.Error(err)
		return
	}

	params := utils.ConvertListMenuRequestToParam(&req)
	menus, total, err := c.MenuService.ListMenus(ctx, params)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	menusInfo := utils.ConvertMenuListToPageResponse(menus)
	ctx.JSON(http.StatusOK, &common.Response{
		Code: common.Ok,
		Data: &common.Page{
			Current:  req.Current,
			PageSize: req.PageSize,
			Total:    total,
			List:     menusInfo,
		},
	})
}
