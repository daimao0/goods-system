package dto

import (
	"time"
)

type OrderDTO struct {
	//id
	Id uint64
	//会员
	Member MemberDTO
	//订单商品
	OrderItems []OrderItemDTO
	//实际支付价格
	Pay float64
	//创建时间
	CreateTime time.Time
}

// OrderStatisticsDTO 订单统计
type OrderStatisticsDTO struct {
	//订单数
	OrderCount uint64
	//商品总价
	GoodsPrice float64
	//用户支付总价
	MemberPay float64
	//用户优惠价格
	PreferentialPrice float64
}
