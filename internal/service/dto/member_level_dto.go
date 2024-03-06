package dto

import "time"

type MemberLevelDTO struct {
	//主键
	Id uint64
	//会员等级名称
	Name string
	//优惠折扣
	Discount float64
	//创建时间
	CreateTime time.Time
}
