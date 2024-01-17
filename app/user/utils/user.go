package utils

import (
	"github.com/pfjhyyj/ether/app/user/define"
	"github.com/pfjhyyj/ether/app/user/model"
	"github.com/pfjhyyj/ether/common"
)

func ConvertUserListPageRequestToParam(req *define.ListUserRequest) *model.QueryUserParams {
	params := &model.QueryUserParams{
		PageRequest: common.PageRequest{
			PageSize: req.PageSize,
			Current:  req.Current,
		},
	}

	if params.PageSize == 0 {
		params.PageSize = 10
	}

	if params.Current == 0 {
		params.Current = 1
	}

	return params
}

func ConvertUserListToPageResponse(users []*model.User) []*define.ListUserPageResponse {
	userInfo := make([]*define.ListUserPageResponse, 0, len(users))
	for _, user := range users {
		userInfo = append(userInfo, &define.ListUserPageResponse{
			UserId:    user.UserId,
			Username:  user.Username,
			Email:     user.Email,
			Mobile:    user.Mobile,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}
	return userInfo
}
