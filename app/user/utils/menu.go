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
		Order:    req.Order,
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
		Order:    req.Order,
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

func ConvertMenuListToPageResponse(menu []*model.Menu) []*define.MenuPageResponse {
	var menuInfos []*define.MenuPageResponse
	for _, m := range menu {
		menuInfos = append(menuInfos, &define.MenuPageResponse{
			MenuId: m.MenuId,
			Name:   m.Name,
		})
	}
	return menuInfos
}

func ConvertMenuToResponse(menus []*model.Menu) *define.GetMenuResponse {
	return &define.GetMenuResponse{
		Menus: parseMenuListToForest(menus, 0),
	}
}

func parseMenuListToForest(menus []*model.Menu, parentId uint) []*define.Menu {
	var menuInfos []*define.Menu
	for _, menu := range menus {
		if menu.ParentId == parentId {
			menuInfos = append(menuInfos, &define.Menu{
				MenuId:   menu.MenuId,
				MenuType: menu.MenuType,
				ParentId: menu.ParentId,
				Name:     menu.Name,
				Path:     menu.Path,
				Locale:   menu.Locale,
				Icon:     menu.Icon,
				Order:    menu.Order,
				Children: parseMenuListToForest(menus, menu.MenuId),
			})
		}
	}
	return menuInfos
}
