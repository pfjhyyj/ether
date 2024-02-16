package define

import (
	"github.com/pfjhyyj/ether/common"
	"time"
)

type ListUserRequest struct {
	common.PageRequest
}

type UserIdUri struct {
	UserId uint `uri:"userId" binding:"required"`
}

type ListUserPageResponse struct {
	UserId    uint      `json:"userId"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Mobile    string    `json:"mobile"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Status    uint      `json:"status"`
}

type GetUserRequest struct {
	UserId uint `uri:"userId" binding:"required"`
}

type GetUserResponse struct {
	UserId    uint      `json:"userId"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Mobile    string    `json:"mobile"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Status    uint      `json:"status"`
}
