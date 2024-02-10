package utils

import (
	"github.com/pfjhyyj/ether/app/notice/define"
	"github.com/pfjhyyj/ether/app/notice/model"
	"github.com/pfjhyyj/ether/common"
)

func ConvertListMyMessagesRequestToQueryParam(req *define.ListMyMessagesRequest) *model.QueryMessageParams {
	return &model.QueryMessageParams{
		PageRequest: common.PageRequest{
			Current:  req.Current,
			PageSize: req.PageSize,
		},
		IsRead: req.IsRead,
	}
}

func ConvertMessageListToListMyMessageResponse(messages []*model.Message) []*define.ListMyMessagesResponse {
	var messagesInfo []*define.ListMyMessagesResponse
	for _, message := range messages {
		messagesInfo = append(messagesInfo, &define.ListMyMessagesResponse{
			MessageId: message.MessageId,
			IsRead:    message.IsRead,
			CreatedAt: message.CreatedAt,
			Category:  message.MessageText.Category,
			Title:     message.MessageText.Title,
			Content:   message.MessageText.Content,
		})
	}
	return messagesInfo
}
