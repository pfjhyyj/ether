package service

import (
	"context"
	"github.com/pfjhyyj/ether/common"
	"github.com/pfjhyyj/ether/utils"
	"github.com/sirupsen/logrus"
)

type SMSService struct {
}

func NewSMSService() *SMSService {
	return &SMSService{}
}

func (s *SMSService) SendSMS(ctx context.Context, phone string) error {
	// 检验是否可以发送
	ok, err := utils.CheckCanSendSMS(ctx, phone)
	if err != nil {
		return &common.SystemError{
			Code: common.UnknownError,
			Msg:  "Check sms rate limit fail",
		}
	}
	if !ok {
		return &common.SystemError{
			Code: common.RequestError,
			Msg:  "send validation code too many times",
		}
	}

	code := utils.GenerateValidationCode()
	// 设置验证码和频率限制
	err = utils.SetSMSCodeAndRateLimit(ctx, phone, code)
	if err != nil {
		logrus.WithContext(ctx).WithError(err).Errorf("fail to set sms code and rate limit")
		return &common.SystemError{
			Code: common.UnknownError,
			Msg:  "set sms code and rate limit fail",
		}
	}
	// 发送短信
	err = utils.SendSMSByAliyun(ctx, phone, code)
	if err != nil {
		logrus.WithContext(ctx).WithError(err).Errorf("fail to send sms")
		return &common.SystemError{
			Code: common.UnknownError,
			Msg:  "send sms fail",
		}
	}
	return nil
}

func (s *SMSService) ValidatePhoneCode(ctx context.Context, phone, code string) error {
	realCode, err := utils.GetSMSCode(ctx, phone)
	if err != nil {
		logrus.WithContext(ctx).WithError(err).Errorf("fail to get sms code")
		return &common.SystemError{
			Code: common.UnknownError,
			Msg:  "get sms code fail",
		}
	}

	if code == "" || realCode != code {
		// 校验次数过多则直接删除该code
		shouldDel, err := utils.CheckCodeFailTimes(ctx, phone)
		if err != nil {
			logrus.WithContext(ctx).WithError(err).Errorf("fail to check code fail times")
			return &common.SystemError{
				Code: common.UnknownError,
				Msg:  "check code fail times fail",
			}
		}
		if shouldDel {
			err = utils.RemoveCodeFailTimes(ctx, phone)
			if err != nil {
				logrus.WithContext(ctx).WithError(err).Errorf("fail to remove code fail times")
			}
		}
		return &common.SystemError{
			Code: common.RequestError,
			Msg:  "invalid code",
		}
	}
	// 删除验证码
	err = utils.RemoveSMSCodeAndRateLimit(ctx, phone)
	if err != nil {
		logrus.WithContext(ctx).WithError(err).Errorf("fail to remove sms code and rate limit")
	}
	return nil
}
