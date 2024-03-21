package utils

import (
	"context"
	smsclient "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/pfjhyyj/ether/client/aliyun"
	"github.com/pfjhyyj/ether/client/redis"
	"github.com/pfjhyyj/ether/common"
	"github.com/spf13/viper"
	"math/rand"
	"strconv"
	"time"
)

const (
	ValidationCodeLength = 6
	MaxCodeFailTimes     = 5
)

func GenerateValidationCode() string {
	var code string
	for i := 0; i < ValidationCodeLength; i++ {
		tmp := rand.Intn(10)
		code += strconv.Itoa(tmp)
	}
	return code
}

func CheckCanSendSMS(ctx context.Context, phone string) (bool, error) {
	// 校验当天总数是否超限
	timesKey := common.GetPhoneSMSSendTimesKey(phone)
	client := redis.GetRedisClient()
	existed, err := client.Exists(ctx, timesKey).Result()
	if err != nil {
		return false, err
	}
	// 增加当天发送次数并设置过期时间
	if existed == 0 {
		pipe := client.TxPipeline()
		pipe.Incr(ctx, timesKey)
		pipe.Expire(ctx, timesKey, time.Hour*24)
		cmder, err := pipe.Exec(ctx)
		if err != nil {
			return false, err
		}
		for _, cmd := range cmder {
			if cmd.Err() != nil {
				return false, err
			}
		}
	} else {
		// 增加当天发送次数
		times, err := client.Incr(ctx, timesKey).Result()
		if err != nil {
			return false, err
		}
		// 校验发送次数是否超限
		if times >= common.ValidationCodeMaxPerDay {
			return false, nil
		}
	}
	// 校验发送频率是否超限
	rateKey := common.GetPhoneSMSSendRateKey(phone)
	rate, err := client.Exists(ctx, rateKey).Result()
	if err != nil {
		return false, err
	}
	if rate > 0 {
		return false, nil
	}
	return true, nil
}

func SetSMSCodeAndRateLimit(ctx context.Context, phone string, code string) error {
	key := common.GetPhoneSMSCodeKey(phone)
	rateKey := common.GetPhoneSMSSendRateKey(phone)
	client := redis.GetRedisClient()
	pipe := client.Pipeline()
	pipe.SetEx(ctx, key, code, common.ValidationCodeExpireTime)
	pipe.SetEx(ctx, rateKey, 1, common.ValidationCodeRate)
	cmder, err := pipe.Exec(ctx)
	if err != nil {
		return err
	}
	for _, cmd := range cmder {
		if cmd.Err() != nil {
			return err
		}
	}
	return nil
}

func GetSMSCode(ctx context.Context, phone string) (string, error) {
	key := common.GetPhoneSMSCodeKey(phone)
	client := redis.GetRedisClient()
	code, err := client.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return code, nil
}

func RemoveSMSCodeAndRateLimit(ctx context.Context, phone string) error {
	key := common.GetPhoneSMSCodeKey(phone)
	rateKey := common.GetPhoneSMSSendRateKey(phone)
	client := redis.GetRedisClient()
	pipe := client.Pipeline()
	pipe.Del(ctx, key)
	pipe.Del(ctx, rateKey)
	cmder, err := pipe.Exec(ctx)
	if err != nil {
		return err
	}
	for _, cmd := range cmder {
		if cmd.Err() != nil {
			return err
		}
	}
	return nil
}

func SendSMSByAliyun(ctx context.Context, phone string, code string) error {
	sign := viper.GetString("aliyun.sms.sign")
	tplCode := viper.GetString("aliyun.sms.template")
	_, err := aliyun.GetSMSClient().SendSms(&smsclient.SendSmsRequest{
		SignName:      &sign,
		TemplateCode:  tea.String(tplCode),
		PhoneNumbers:  tea.String(phone),
		TemplateParam: tea.String("{\"code\":\"" + code + "\"}"),
	})
	if err != nil {
		return err
	}
	return nil
}

// GetCodeFailTimes 获取验证码校验失败次数并+1
func GetCodeFailTimes(ctx context.Context, phone string) (int64, error) {
	key := common.GetPhoneSMSValidationTimesKey(phone)
	client := redis.GetRedisClient()
	times, err := client.Incr(ctx, key).Result()
	if err != nil {
		return 0, err
	}
	return times, nil
}

func CheckCodeFailTimes(ctx context.Context, phone string) (bool, error) {
	times, err := GetCodeFailTimes(ctx, phone)
	if err != nil {
		return false, err
	}
	if times >= MaxCodeFailTimes {
		return true, nil
	}
	return false, nil
}

func RemoveCodeFailTimes(ctx context.Context, phone string) error {
	key := common.GetPhoneSMSValidationTimesKey(phone)
	client := redis.GetRedisClient()
	_, err := client.Del(ctx, key).Result()
	if err != nil {
		return err
	}
	return nil
}
