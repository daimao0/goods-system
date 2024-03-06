package request

// OrderItemRequest 订单单项
type OrderItemRequest struct {
	//商品id
	GoodsId string `json:"goodsId"`
	//数量
	Num int `json:"num"`
	//实际单价
	RealPrice string `json:"realPrice"`
}
