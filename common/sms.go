package common

import (
	"fmt"
	"time"
)

const (
	ValidationCodeMaxPerDay       = 10
	PhoneSMSCodeKeyPrefix         = "phone_code:"
	PhoneSMSValidationTimesPrefix = "phone_sms_validation_times:"
	PhoneSMSSendTimesKeyPrefix    = "phone_sms_send_times:"
	PhoneSMSSendRateKeyPrefix     = "phone_sms_send_rate:"

	MaxCodeFailTimes         = 5 // 最大允许的code校验错误次数
	ValidationCodeExpireTime = time.Minute * 5
	ValidationCodeRate       = time.Second * 90
)

func GetPhoneSMSCodeKey(phone string) string {
	return fmt.Sprintf("%s%s", PhoneSMSCodeKeyPrefix, phone)
}

func GetPhoneSMSValidationTimesKey(phone string) string {
	return fmt.Sprintf("%s%s", PhoneSMSValidationTimesPrefix, phone)
}

func GetPhoneSMSSendTimesKey(phone string) string {
	// 获取当天的时间
	today := time.Now().Format("20060102")
	return fmt.Sprintf("%s%s_%s", PhoneSMSSendTimesKeyPrefix, today, phone)
}

func GetPhoneSMSSendRateKey(phone string) string {
	return fmt.Sprintf("%s%s", PhoneSMSSendRateKeyPrefix, phone)
}
