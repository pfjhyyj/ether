package domain

import (
	"context"
	"github.com/pfjhyyj/ether/app/notice/constants"
	"github.com/pfjhyyj/ether/app/notice/model"
	"github.com/pfjhyyj/ether/client/gorm"
	"github.com/pfjhyyj/ether/domain/notice"
	"github.com/sirupsen/logrus"
)

type NoticeRepository struct {
	notice.Repository
}

func GetNoticeRepository() *NoticeRepository {
	Init()
	return noticeRepository
}

func (r *NoticeRepository) NotifyUsers(ctx context.Context, message *notice.Message, userIds []uint) error {
	logs := logrus.WithContext(ctx)
	db := gorm.GetDB().WithContext(ctx)

	messageText := &model.MessageText{
		Category: message.Category,
		Title:    message.Title,
		Content:  message.Content,
	}

	if err := model.CreateMessageText(db, messageText); err != nil {
		logs.WithError(err).Error("create message text failed")
		return err
	}

	messageTextId := messageText.MessageTextId
	var messages []*model.Message
	for _, userId := range userIds {
		messages = append(messages, &model.Message{
			UserId:        userId,
			IsRead:        constants.MessageUnread,
			MessageTextId: messageTextId,
		})
	}

	if err := model.CreateMessages(db, messages); err != nil {
		logs.WithError(err).Error("create messages failed")
		return err
	}

	return nil
}

func (r *NoticeRepository) GetUnreadMessageCount(ctx context.Context, userId uint) (int64, error) {
	logs := logrus.WithContext(ctx)
	db := gorm.GetDB().WithContext(ctx)

	count, err := model.GetUnreadMessageCount(db, userId)
	if err != nil {
		logs.WithError(err).Error("get unread message count failed")
		return 0, err
	}
	return count, nil
}
