package param

import "time"

// OrderSearchParam 订单搜索参数
type OrderSearchParam struct {
	//会员id
	MemberId uint64
	//开始时间
	StartTime time.Time
	//结束时间
	EndTime time.Time
}

// OrderPageParam 订单分页参数
type OrderPageParam struct {
	MemberId  uint64
	StartTime time.Time
	EndTime   time.Time
	Page      int
	Size      int
}

// OrderSaveParam 订单保存参数
type OrderSaveParam struct {
	//会员id
	MemberId uint64
	//商品ids
	OrderItemIds []uint64
	//实际支付价格
	Pay float64
}
