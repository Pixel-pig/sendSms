package util

import (
	"encoding/json"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
)

//模仿配置文件
const (
	ACCESSKEYID     = "LTAI4G5odyGRBe6ettBhroXr"
	ACCESSKEYSECRET = "vUU1ZvPyFA9CwLLz49zqS50jrlx22j"
	SIGNNAME        = "线上餐厅"
)

//模板号
const SMS_TPL_REGISTER = "SMS_205393604"

//封装动态编码
type SmsCode struct {
	Code string `json:"code"`
}

//获取返回的数据信息
type SmsResult struct {
	RequestId string
	BizId     string
	Code      string
	Message   string
}

func SendSms(phone string, code string, templateType string) (*SmsResult, error) {
	//读取配置文件

	client, err := dysmsapi.NewClientWithAccessKey("cn-hangzhou", ACCESSKEYID, ACCESSKEYSECRET)

	//创建一个发送短信的请求
	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https" //使用https协议，更安全

	//电话号码
	request.PhoneNumbers = phone
	//短信签名
	request.SignName = SIGNNAME
	//短信模板号
	request.TemplateCode = templateType

	//生成验证码
	smsCode := SmsCode{Code:code}
	codeBytes, _ := json.Marshal(smsCode)
	//指定短信模板中的动态验证码 的数据
	request.TemplateParam = string(codeBytes)

	response, err := client.SendSms(request)
	if err != nil {
		return nil, err
	}
	result := &SmsResult{
		RequestId: response.RequestId,
		BizId:     response.BizId,
		Code:      response.Code,
		Message:   response.Message,
	}
	return result, nil
}
