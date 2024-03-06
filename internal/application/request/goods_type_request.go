package request

// GoodsTypeSaveRequest 商品类型保存请求参数
type GoodsTypeSaveRequest struct {
	//商品名称
	Name string `json:"name"`
}

// GoodsTypeEditRequest 商品类型编辑请求参数
type GoodsTypeEditRequest struct {
	//主键
	Id string `json:"id"`
	//商品名称
	Name string `json:"name"`
}

// GoodsTypeDeleteRequest 商品类型删除请求参数
type GoodsTypeDeleteRequest struct {
	//商品名称
	Name string `json:"name"`
}
