package define

import (
	"github.com/pfjhyyj/ether/common"
	"time"
)

type ListMyMessagesRequest struct {
	common.PageRequest
	Category uint `form:"category"`
	IsRead   uint `form:"isRead"`
}

type ListMyMessagesResponse struct {
	MessageId uint      `json:"messageId"`
	IsRead    uint      `json:"isRead"`
	CreatedAt time.Time `json:"createdAt"`
	Category  uint      `json:"category"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
}

type ReadMessageRequest struct {
	MessageId uint `uri:"messageId" binding:"required"`
}

type BatchReadMessageRequest struct {
	MessageIds []uint `json:"messageIds"`
}
