package vo

// OrderVO 订单VO
type OrderVO struct {
	//id
	Id string `json:"id"`
	//会员
	Member MemberVO `json:"member"`
	//订单商品
	OrderItems []OrderItemVO `json:"orderItems"`
	//实际支付价格
	Pay string `json:"pay"`
	//创建时间
	CreateTime string `json:"createTime"`
}

// OrderStatisticsVO 订单统计
type OrderStatisticsVO struct {
	//订单数
	OrderCount string `json:"orderCount"`
	//商品总价
	GoodsPrice string `json:"goodsPrice"`
	//用户支付总价
	MemberPay string `json:"memberPay"`
	//用户优惠价格
	PreferentialPrice string `json:"preferentialPrice"`
}
