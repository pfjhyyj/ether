package define

import (
	"github.com/pfjhyyj/ether/common"
	"time"
)

type ListUserRequest struct {
	common.PageRequest
}

type ListUserPageResponse struct {
	UserId    uint
	Username  string
	Email     string
	Mobile    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
