package param

// GoodsTypeSaveParam 商品类型保存参数
type GoodsTypeSaveParam struct {
	//商品类型名称
	Name string `json:"name"`
}

// GoodsTypeUpdateParam 商品类型编辑参数
type GoodsTypeUpdateParam struct {
	//商品类型id
	Id uint64 `json:"id"`
	//商品类型名称
	Name string `json:"name"`
}
