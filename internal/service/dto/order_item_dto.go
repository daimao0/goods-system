package dto

import "time"

type OrderItemDTO struct {
	//id
	Id uint64
	//商品信息
	GoodsDTO GoodsDTO
	//实际支付的单价
	RealPrice float64
	//数量
	Num int
	//创建时间
	CreateTime time.Time
}
