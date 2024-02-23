package utils

import (
	"github.com/pfjhyyj/ether/app/user/define"
	"github.com/pfjhyyj/ether/app/user/model"
)

func ConvertUserToMyInfoResponse(user *model.User) *define.MyInfoResponse {
	return &define.MyInfoResponse{
		UserId:   user.UserId,
		Username: user.Username,
		Avatar:   user.Avatar,
	}
}

func ConvertUpdateMyInfoRequestToModel(req *define.UpdateMyInfoRequest) *model.User {
	return &model.User{
		Avatar: req.Avatar,
	}
}

func ConvertUpdateMyPasswordRequestToModel(req *define.UpdateMyPasswordRequest) *model.User {
	return &model.User{
		Password: req.NewPassword,
	}
}

func ConvertMyMenuToResponse(menus []*model.Menu) *define.GetMyMenuResponse {
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
	return &define.GetMyMenuResponse{
		Menus: menuInfos,
	}
}

func ConvertMyUnreadMessageCountToResponse(count int64) *define.GetMyUnreadMessageCountResponse {
	return &define.GetMyUnreadMessageCountResponse{
		UnreadMessageCount: count,
	}
}
