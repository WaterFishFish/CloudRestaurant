package DAO

import (
	"CloudRestaurant/model"
	"CloudRestaurant/tool"
	"fmt"
)

type MemberDao struct {
	*tool.Orm
}

// 验证手机号和验证码是否存在
func (md *MemberDao) ValidateSmsCode(phone string, code string) *model.SmsCode {
	var sms model.SmsCode
	if _, err := md.Where("phone = ? and code = ?", phone, code).Get(&sms); err != nil {
		fmt.Printf("err:%v\n", err)
	}
	return &sms

}
func (md *MemberDao) QueryByPhone(phone string) *model.Member {
	var member model.Member
	if _, err := md.Where("mobile = ?", phone).Get(&member); err != nil {
		fmt.Printf("err:%v\n", err)
	}
	return &member
}

// 新用户数据库插入
func (md *MemberDao) InsertMember(member model.Member) int64 {
	result, err := md.InsertOne(&member)
	if err != nil {
		fmt.Printf("err:%v\n", err)
		return 0
	}
	return result
}
func (md *MemberDao) InsertCode(sms model.SmsCode) int64 {
	result, err := md.InsertOne(&sms)
	if err != nil {
		fmt.Printf("err:%v\n", err)
	}
	return result
}
