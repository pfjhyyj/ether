package define

type RegisterUserRequest struct {
	Username string `json:"username" binding:"required,min=6,max=20"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,max=20"`
}
