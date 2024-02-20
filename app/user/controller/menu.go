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
	service *service.MenuService
}

func NewMenuController(service *service.MenuService) *MenuController {
	return &MenuController{
		service: service,
	}
}

// CreateMenu godoc
// @Summary Create menu
// @Description Create menu
// @Tags menu
// @Accept json
// @Produce json
// @Security Bearer
// @Param request body define.CreateMenuRequest true "CreateMenuRequest"
// @Success 200 {object} common.Response
// @Router /menus [post]
func (c *MenuController) CreateMenu(ctx *gin.Context) {
	if ok := utils2.CheckPermission(ctx, "menu", "create"); !ok {
		ctx.JSON(http.StatusOK, &common.Response{
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
	if err := c.service.CreateMenu(ctx, menu); err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, &common.Response{
		Code: common.Ok,
	})
}

// UpdateMenu godoc
// @Summary Update menu
// @Description Update menu
// @Tags menu
// @Accept json
// @Produce json
// @Security Bearer
// @Param request body define.UpdateMenuRequest true "UpdateMenuRequest"
// @Success 200 {object} common.Response
// @Router /menus/{menuId} [put]
func (c *MenuController) UpdateMenu(ctx *gin.Context) {
	if ok := utils2.CheckPermission(ctx, "menu", "update"); !ok {
		ctx.JSON(http.StatusOK, &common.Response{
			Code: common.NoPermissionError,
			Msg:  "no permission",
		})
		return
	}

	var menuIdReq define.MenuIdUri
	if err := ctx.ShouldBindUri(&menuIdReq); err != nil {
		_ = ctx.Error
		return
	}

	var req define.UpdateMenuRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		_ = ctx.Error(err)
		return
	}
	req.MenuId = menuIdReq.MenuId

	menu := utils.ConvertUpdateMenuRequestToMenu(&req)
	if err := c.service.UpdateMenu(ctx, req.MenuId, menu); err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, &common.Response{
		Code: common.Ok,
	})
}

// DeleteMenu godoc
// @Summary Delete menu
// @Description Delete menu
// @Tags menu
// @Accept json
// @Produce json
// @Security Bearer
// @Param request body define.DeleteMenuRequest true "DeleteMenuRequest"
// @Success 200 {object} common.Response
// @Router /menus/{menuId} [delete]
func (c *MenuController) DeleteMenu(ctx *gin.Context) {
	if ok := utils2.CheckPermission(ctx, "menu", "delete"); !ok {
		ctx.JSON(http.StatusOK, &common.Response{
			Code: common.NoPermissionError,
			Msg:  "no permission",
		})
		return
	}

	var menuIdReq define.MenuIdUri
	if err := ctx.ShouldBindUri(&menuIdReq); err != nil {
		_ = ctx.Error(err)
		return
	}

	if err := c.service.DeleteMenu(ctx, menuIdReq.MenuId); err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, &common.Response{
		Code: common.Ok,
	})
}

//// GetMenuTree godoc
//// @Summary Get menu tree
//// @Description Get menu tree
//// @Tags menu
//// @Accept json
//// @Produce json
//// @Security Bearer
//// @Param menu_id query int true "menu_id"
//// @Success 200 {object} common.Response{data=define.GetMenuResponse}
//// @Router /menus/{menuId}/tree [get]
//func (c *MenuController) GetMenuTree(ctx *gin.Context) {
//	if ok := utils2.CheckPermission(ctx, "menu", "get"); !ok {
//		ctx.JSON(http.StatusOK, &common.Response{
//			Code: common.NoPermissionError,
//			Msg:  "no permission",
//		})
//		return
//	}
//
//	var req define.GetMenuRequest
//	if err := ctx.ShouldBindUri(&req); err != nil {
//		_ = ctx.Error(err)
//		return
//	}
//
//	menu, err := c.service.GetMenuByMenuId(ctx, req.MenuId)
//	if err != nil {
//		_ = ctx.Error(err)
//		return
//	}
//
//	menus, err := c.service.GetMenuTreeByMenuId(ctx, req.MenuId)
//	if err != nil {
//		_ = ctx.Error(err)
//		return
//	}
//
//	menuInfo := utils.ConvertMenuToResponse(menu, menus)
//	ctx.JSON(http.StatusOK, &common.Response{
//		Code: common.Ok,
//		Data: menuInfo,
//	})
//}

// ListMenus godoc
// @Summary List menus
// @Description List menus
// @Tags menu
// @Accept json
// @Produce json
// @Security Bearer
// @Param request query define.ListMenusRequest true "ListMenusRequest"
// @Success 200 {object} common.Response{data=[]define.MenuPageResponse}
// @Router /menus [get]
func (c *MenuController) ListMenus(ctx *gin.Context) {
	if ok := utils2.CheckPermission(ctx, "menu", "list"); !ok {
		ctx.JSON(http.StatusOK, &common.Response{
			Code: common.NoPermissionError,
			Msg:  "no permission",
		})
		return
	}

	menus, err := c.service.ListAllMenus(ctx)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	menusInfo := utils.ConvertMenuListToPageResponse(menus)
	ctx.JSON(http.StatusOK, &common.Response{
		Code: common.Ok,
		Data: menusInfo,
	})
}
