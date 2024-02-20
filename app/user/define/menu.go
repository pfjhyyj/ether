package define

import "github.com/pfjhyyj/ether/common"

type MenuIdUri struct {
	MenuId uint `uri:"menuId" binding:"required"`
}

type CreateMenuRequest struct {
	MenuType uint   `json:"menuType" binding:"required"`
	ParentId uint   `json:"parentId"`
	Name     string `json:"name"`
	Path     string `json:"path"`
	Locale   string `json:"locale"`
	Icon     string `json:"icon"`
	Weight   int    `json:"weight"`
}

type UpdateMenuRequest struct {
	MenuId   uint
	MenuType uint   `json:"menuType"`
	ParentId uint   `json:"parentId"`
	Name     string `json:"name"`
	Path     string `json:"path"`
	Locale   string `json:"locale"`
	Icon     string `json:"icon"`
	Weight   int    `json:"weight"`
}

type DeleteMenuRequest struct {
	MenuId uint `uri:"menuId" binding:"required"`
}

type ListMenusRequest struct {
	common.PageRequest
}

type MenuPageResponse struct {
	MenuId   uint    `json:"menuId"`
	MenuType uint    `json:"menuType"`
	ParentId uint    `json:"parentId"`
	Name     string  `json:"name"`
	Path     string  `json:"path"`
	Locale   string  `json:"locale"`
	Icon     string  `json:"icon"`
	Weight   int     `json:"weight"`
	Children []*Menu `json:"children"`
}

type GetMenuRequest struct {
	MenuId uint `uri:"menuId" binding:"required"`
}

type GetMenuResponse struct {
	MenuId   uint    `json:"menuId"`
	MenuType uint    `json:"menuType"`
	ParentId uint    `json:"parentId"`
	Name     string  `json:"name"`
	Path     string  `json:"path"`
	Locale   string  `json:"locale"`
	Icon     string  `json:"icon"`
	Weight   int     `json:"weight"`
	Menus    []*Menu `json:"menus"`
}

type Menu struct {
	MenuId   uint   `json:"menuId"`
	MenuType uint   `json:"menuType"`
	ParentId uint   `json:"parentId"`
	Name     string `json:"name"`
	Path     string `json:"path"`
	Locale   string `json:"locale"`
	Icon     string `json:"icon"`
	Weight   int    `json:"weight"`
}
