package param

// OrderItemSaveParam 订单子项保存参数
type OrderItemSaveParam struct {
	//商品id
	GoodsId uint64
	//实际单价
	RealPrice float64
	//数量
	Num int
}
