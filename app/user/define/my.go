package define

type MyInfoResponse struct {
	UserId   uint   `json:"userId"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Mobile   string `json:"mobile"`
	Avatar   string `json:"avatar"`
}

type UpdateMyInfoRequest struct {
	Avatar string `json:"avatar"`
}

type UpdateMyPasswordRequest struct {
	OldPassword    string `json:"oldPassword" binding:"required"`
	NewPassword    string `json:"newPassword" binding:"required"`
	RepeatPassword string `json:"repeatPassword" binding:"required"`
}
