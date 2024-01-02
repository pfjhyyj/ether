package controller

import "context"

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken            string `json:"access_token"`
	ExpireTime             int64  `json:"expire_time"`
	RefreshToken           string `json:"refresh_token"`
	RefreshTokenExpireTime int64  `json:"refresh_token_expire_time"`
}

func Login(ctx context.Context, req LoginRequest) (LoginResponse, error) {
	return LoginResponse{}, nil
}
