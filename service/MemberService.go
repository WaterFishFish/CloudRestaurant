package service

import (
	"CloudRestaurant/DAO"
	"CloudRestaurant/model"
	"CloudRestaurant/param"
	"CloudRestaurant/tool"
	"encoding/json"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/gin-gonic/gin"
	"math/rand"
	"time"
)

type MemberService struct {
}

func (ms *MemberService) SmsLogin(loginparam param.SmsLoginParam) *model.Member {
	//1.获取手机号和验证码
	//2.验证手机号+验证码是否正确
	md := DAO.MemberDao{tool.DbEngine}
	sms := md.ValidateSmsCode(loginparam.Phone, loginparam.Code)
	if sms.Id == 0 {
		return nil
	}
	//3.根据手机号member表中查询记录
	member := md.QueryByPhone(loginparam.Phone)
	if member.Id != 0 {
		return member
	}

	//4.新创建一个member记录，并保存
	user := model.Member{}
	user.UserName = loginparam.Phone
	user.Mobile = loginparam.Phone
	user.RegisterTime = time.Now().Unix()
	user.Id = md.InsertMember(user)
	return &user
}

// 发送验证码
func (ms *MemberService) SendCode(phone string) bool {
	//1.产生一个验证码
	code := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(100000))
	//2.调用阿里云SDK完成发送
	config := tool.GetConfig().Sms
	client, err := dysmsapi.NewClientWithAccessKey(config.RegionId, config.AppKey, config.AppSecret)
	if err != nil {
		fmt.Printf("err:%v\n", err)
		return false
	}
	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.SignName = config.SignName
	request.TemplateCode = config.TemplateCode
	request.PhoneNumbers = phone
	par, err := json.Marshal(gin.H{
		"code": code,
	})
	request.TemplateParam = string(par)
	response, err := client.SendSms(request)
	fmt.Println(response)
	if err != nil {
		fmt.Printf("err:%v\n", err)
		return false
	}

	//3.接受返回值，并判断发送状态
	//短信验证码发送成功
	if response.Code == "OK" {
		smsCode := model.SmsCode{Phone: phone, Code: code, BizId: response.BizId, CreateTime: time.Now().Unix()}
		memberDao := DAO.MemberDao{tool.DbEngine}
		result := memberDao.InsertCode(smsCode)
		return result > 0
	}
	return false
}
