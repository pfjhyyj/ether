package service

import (
	"github.com/gin-gonic/gin"
	"github.com/pfjhyyj/ether/app/notice/model"
	"github.com/pfjhyyj/ether/clients/gorm"
	"github.com/sirupsen/logrus"
)

type MessageService struct{}

func NewMessageService() *MessageService {
	return &MessageService{}
}

func (s *MessageService) ListMyMessages(ctx *gin.Context, param *model.QueryMessageParams) ([]*model.Message, int64, error) {
	logs := logrus.WithContext(ctx)
	db := gorm.GetDB().WithContext(ctx)

	messages, total, err := model.ListMessages(db, param)
	if err != nil {
		logs.WithError(err).Error("list my messages failed")
		return nil, 0, err
	}

	return messages, total, nil
}

func (s *MessageService) ReadMessage(ctx *gin.Context, messageId uint) error {
	logs := logrus.WithContext(ctx)
	db := gorm.GetDB().WithContext(ctx)

	if err := model.SetMessageRead(db, messageId); err != nil {
		logs.WithError(err).Error("update message failed")
		return err
	}

	return nil
}

func (s *MessageService) BatchReadMessage(ctx *gin.Context, messageIds []uint) error {
	logs := logrus.WithContext(ctx)
	db := gorm.GetDB().WithContext(ctx)

	if err := model.BatchSetMessageRead(db, messageIds); err != nil {
		logs.WithError(err).Error("batch read message failed")
		return err
	}

	return nil
}
