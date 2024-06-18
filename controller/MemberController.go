package controller

import (
	"CloudRestaurant/param"
	"CloudRestaurant/service"
	"CloudRestaurant/tool"
	"fmt"
	"github.com/gin-gonic/gin"
)

type MemberController struct {
}

func (mc *MemberController) Router(engin *gin.Engine) {
	engin.GET("/api/sendcode", mc.sendSmsCode)
	engin.POST("/api/login_sms", mc.smsLogin)
	engin.GET("/api/captcha", mc.captcha)
	engin.POST("/api/vertifycha", mc.vertifyCaptcha)
}

// 生成验证码
func (mc *MemberController) captcha(context *gin.Context) {
	tool.GenerateCaptcha(context)
}

// 验证验证码是否正确
func (mc *MemberController) vertifyCaptcha(context *gin.Context) {
	var captcha tool.CaptchaResult
	err := tool.Decode(context.Request.Body, &captcha)
	if err != nil {
		tool.Failed(context, "参数解析失败")
		return
	}
	result := tool.VertifyCaptcha(captcha.Id, captcha.VertifyValue)
	if result {
		fmt.Println("验证成功")
	} else {
		fmt.Println("验证失败")
	}
}

// http://localhost:8090/api/sendcode?phone=12345678910
func (mc *MemberController) sendSmsCode(context *gin.Context) {
	//发送验证码
	phone, err := context.GetQuery("phone")
	if !err {
		tool.Failed(context, "参数解析失败")

		return
	}
	ms := service.MemberService{}
	isSend := ms.SendCode(phone)
	if isSend {
		tool.Success(context, "发送成功")

		return
	}
	tool.Failed(context, "发送失败")
}

// 手机号+短信 登录的方式
func (mc *MemberController) smsLogin(context *gin.Context) {
	var smsLoginParam param.SmsLoginParam
	err := tool.Decode(context.Request.Body, &smsLoginParam)
	if err != nil {

		tool.Failed(context, "参数解析失败")
		return
	}
	//完成登录手机+验证码登录
	us := service.MemberService{}
	member := us.SmsLogin(smsLoginParam)
	if member != nil {
		tool.Success(context, member)
		return
	}
	tool.Failed(context, "登录失败")
}
