package dto

import "time"

type MemberDTO struct {
	//主键
	Id uint64
	//会员名称
	Name string
	//手机号
	Phone string
	//优惠折扣
	Discount float64
	//会员等级
	Level MemberLevelDTO
	//账户余额
	Account float64
	//创建时间
	CreateTime time.Time
	//更新时间
	UpdateTime time.Time
}
