package define

type LoginByUsernameRequest struct {
	Username string `json:"username,min=6,max=20"`
	Password string `json:"password,min=8,max=20"`
}

type TokenResponse struct {
	AccessToken string `json:"accessToken"`
	ExpireTime  int64  `json:"expireTime"`
}
