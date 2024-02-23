package model

import (
	"github.com/pfjhyyj/ether/app/notice/constants"
	"github.com/pfjhyyj/ether/common"
	"gorm.io/gorm"
)

type Message struct {
	MessageId     uint `gorm:"primaryKey"`
	UserId        uint `gorm:"column:user_id"`
	IsRead        uint `gorm:"column:is_read"`
	MessageTextId uint `gorm:"column:message_text_id"`
	common.Model

	MessageText MessageText `gorm:"foreignKey:message_text_id"`
}

func (Message) TableName() string {
	return "message"
}

type QueryMessageParams struct {
	common.PageRequest

	UserId   uint
	Category uint
	IsRead   uint
	WithText bool
}

func CreateMessage(tx *gorm.DB, message *Message) error {
	return tx.Create(message).Error
}

func CreateMessages(tx *gorm.DB, messages []*Message) error {
	return tx.Create(messages).Error
}

func UpdateMessage(tx *gorm.DB, messageId uint, message *Message) error {
	return tx.Where("message_id = ?", messageId).Updates(message).Error
}

func DeleteMessage(tx *gorm.DB, messageId uint) error {
	return tx.Delete(&Message{}, "message_id = ?", messageId).Error
}

func GetMessageByMessageId(tx *gorm.DB, messageId uint) (*Message, error) {
	var message Message
	if err := tx.Where("message_id = ?", messageId).First(&message).Error; err != nil {
		return nil, err
	}
	return &message, nil
}

func GetMessageByMessageIds(tx *gorm.DB, messageIds []uint) ([]*Message, error) {
	if len(messageIds) == 0 {
		return nil, &common.SystemError{Code: common.DbError, Msg: "message ids is empty"}
	}
	var messages []*Message
	tx.Where("message_id IN ?", messageIds).Find(&messages)
	return messages, nil
}

func ListMessages(tx *gorm.DB, params *QueryMessageParams) ([]*Message, int64, error) {
	var messages []*Message
	query := tx.Model(&Message{})

	var total int64
	query.Count(&total)

	if params.Current > 0 && params.PageSize > 0 {
		query.Offset((params.Current - 1) * params.PageSize).Limit(params.PageSize)
	}

	if params.UserId > 0 {
		query = query.Where("user_id = ?", params.UserId)
	}

	if params.Category > 0 {
		query = query.Where("category = ?", params.Category)
	}

	if params.IsRead > 0 {
		query = query.Where("is_read = ?", params.IsRead)
	}

	if params.WithText {
		query = query.Preload("MessageText")
	}

	if err := query.Find(&messages).Error; err != nil {
		return nil, 0, err
	}

	return messages, total, nil
}

func SetMessageRead(tx *gorm.DB, messageId uint) error {
	return tx.Model(&Message{}).Where("message_id = ?", messageId).Update("is_read", constants.MessageRead).Error
}

func BatchSetMessageRead(tx *gorm.DB, messageIds []uint) error {
	return tx.Model(&Message{}).Where("message_id IN ?", messageIds).Update("is_read", constants.MessageRead).Error
}

func GetUnreadMessageCount(tx *gorm.DB, userId uint) (int64, error) {
	var count int64
	if err := tx.Model(&Message{}).Where("user_id = ? AND is_read = ?", userId, constants.MessageUnread).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
