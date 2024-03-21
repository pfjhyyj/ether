package aliyun

import (
	"sync"

	"github.com/spf13/viper"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
)

var (
	smsClient *dysmsapi20170525.Client

	smsInitOnce sync.Once
)

func InitSMS() {
	smsInitOnce.Do(func() {
		id := viper.GetString("aliyun.access_key.id")
		secret := viper.GetString("aliyun.access_key.secret")
		endpoint := viper.GetString("aliyun.sms.endpoint")
		config := &openapi.Config{
			AccessKeyId:     &id,
			AccessKeySecret: &secret,
		}
		// 访问的域名
		config.Endpoint = &endpoint
		_result, err := dysmsapi20170525.NewClient(config)
		if err != nil {
			panic(err)
		}
		smsClient = _result
	})
}

func GetSMSClient() *dysmsapi20170525.Client {
	InitSMS()
	return smsClient
}
