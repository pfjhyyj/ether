package model

import (
	"github.com/pfjhyyj/ether/common"
	"gorm.io/gorm"
)

type MessageText struct {
	MessageTextId uint   `gorm:"primaryKey"`
	Category      uint   `gorm:"column:category"`
	Title         string `gorm:"column:title"`
	Content       string `gorm:"column:content"`
	common.Model
}

func (MessageText) TableName() string {
	return "message_text"
}

type QueryMessageTextParams struct {
	common.PageRequest
}

func CreateMessageText(tx *gorm.DB, messageText *MessageText) error {
	return tx.Create(messageText).Error
}

func UpdateMessageText(tx *gorm.DB, messageTextId uint, messageText *MessageText) error {
	return tx.Where("message_text_id = ?", messageTextId).Updates(messageText).Error
}

func DeleteMessageText(tx *gorm.DB, messageTextId uint) error {
	return tx.Delete(&MessageText{}, "message_text_id = ?", messageTextId).Error
}

func GetMessageTextByMessageTextId(tx *gorm.DB, messageTextId uint) (*MessageText, error) {
	var messageText MessageText
	if err := tx.Where("message_text_id = ?", messageTextId).First(&messageText).Error; err != nil {
		return nil, err
	}
	return &messageText, nil
}

func GetMessageTextByMessageTextIds(tx *gorm.DB, messageTextIds []uint) ([]*MessageText, error) {
	if len(messageTextIds) == 0 {
		return nil, &common.SystemError{Code: common.DbError, Msg: "message text ids is empty"}
	}
	var messageTexts []*MessageText
	tx.Where("message_text_id IN ?", messageTextIds).Find(&messageTexts)
	return messageTexts, nil
}

func ListMessageTexts(tx *gorm.DB, params *QueryMessageTextParams) ([]*MessageText, int64, error) {
	var messageTexts []*MessageText
	query := tx.Model(&MessageText{})

	var total int64
	query.Count(&total)

	if params.Current > 0 && params.PageSize > 0 {
		query.Offset((params.Current - 1) * params.PageSize).Limit(params.PageSize)
	}

	if err := query.Find(&messageTexts).Error; err != nil {
		return nil, 0, err
	}

	return messageTexts, total, nil
}
