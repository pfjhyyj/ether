package utils

import (
	"github.com/pfjhyyj/ether/app/user/define"
	"github.com/pfjhyyj/ether/app/user/model"
	"github.com/pfjhyyj/ether/common"
)

func ConvertCreateMenuRequestToMenu(req *define.CreateMenuRequest) *model.Menu {
	return &model.Menu{
		MenuType: req.MenuType,
		ParentId: req.ParentId,
		Name:     req.Name,
		Path:     req.Path,
		Locale:   req.Locale,
		Icon:     req.Icon,
		Weight:   req.Weight,
	}
}

func ConvertUpdateMenuRequestToMenu(req *define.UpdateMenuRequest) *model.Menu {
	return &model.Menu{
		MenuId:   req.MenuId,
		MenuType: req.MenuType,
		ParentId: req.ParentId,
		Name:     req.Name,
		Path:     req.Path,
		Locale:   req.Locale,
		Icon:     req.Icon,
		Weight:   req.Weight,
	}
}

func ConvertListMenuRequestToParam(req *define.ListMenusRequest) *model.QueryMenuParams {
	return &model.QueryMenuParams{
		PageRequest: common.PageRequest{
			Current:  req.Current,
			PageSize: req.PageSize,
		},
	}

}

func ConvertMenuListToPageResponse(menus []*model.Menu) []*define.MenuPageResponse {
	var menuInfos []*define.MenuPageResponse
	for _, menu := range menus {
		menuInfos = append(menuInfos, &define.MenuPageResponse{
			MenuId:   menu.MenuId,
			MenuType: menu.MenuType,
			ParentId: menu.ParentId,
			Name:     menu.Name,
			Path:     menu.Path,
			Locale:   menu.Locale,
			Icon:     menu.Icon,
			Weight:   menu.Weight,
		})
	}
	return menuInfos
}

func ConvertMenuToResponse(m *model.Menu, menus []*model.Menu) *define.GetMenuResponse {
	var menuInfos []*define.Menu
	for _, menu := range menus {
		menuInfos = append(menuInfos, &define.Menu{
			MenuId:   menu.MenuId,
			MenuType: menu.MenuType,
			ParentId: menu.ParentId,
			Name:     menu.Name,
			Path:     menu.Path,
			Locale:   menu.Locale,
			Icon:     menu.Icon,
			Weight:   menu.Weight,
		})
	}
	return &define.GetMenuResponse{
		MenuId:   m.MenuId,
		MenuType: m.MenuType,
		ParentId: m.ParentId,
		Name:     m.Name,
		Path:     m.Path,
		Locale:   m.Locale,
		Icon:     m.Icon,
		Weight:   m.Weight,
		Menus:    menuInfos,
	}
}
