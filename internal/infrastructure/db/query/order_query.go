package query

import "time"

// OrderQuery 订单分页参数
type OrderQuery struct {
	// 用户id
	MemberId uint64
	//createTime的起始时间
	StartTime, EndTime time.Time
}

// OrderPageQuery 订单分页参数
type OrderPageQuery struct {
	// 用户id
	MemberId uint64
	//createTime的起始时间
	StartTime, EndTime time.Time
	//分页大小
	Page, Size int
}
