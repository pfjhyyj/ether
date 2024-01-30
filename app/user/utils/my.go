package utils

import (
	"github.com/pfjhyyj/ether/app/user/define"
	"github.com/pfjhyyj/ether/app/user/model"
)

func ConvertUserToMyInfoResponse(user *model.User) *define.MyInfoResponse {
	return &define.MyInfoResponse{
		UserId:   user.UserId,
		Username: user.Username,
		Email:    user.Email,
		Mobile:   user.Mobile,
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
