package vo

// OrderItemVO 订单子项VO
type OrderItemVO struct {

	//id
	Id string `json:"id"`
	//商品信息
	Goods GoodsVO `json:"goods"`
	//实际支付的单价
	RealPrice string `json:"realPrice"`
	//数量
	Num int `json:"num"`
	//创建时间
	CreateTime string `json:"createTime"`
}
