package request

// GoodsPageRequest 商品请求参数
type GoodsPageRequest struct {
	//商品名称
	Name string `json:"name"`
	//商品类型
	GoodsType string `json:"goodsType"`
	//当前页
	Page int `json:"page"`
	//页面大小
	Size int `json:"size"`
}

// GoodsSaveRequest 商品请求参数
type GoodsSaveRequest struct {
	//商品名称
	Name string `json:"name"`
	//商品类型
	GoodsType string `json:"goodsType"`
	//商品编号
	GoodsNumber string `json:"goodsNumber"`
	//商品价格
	Price string `json:"price"`
	//数量
	Count string `json:"count"`
	//描述
	Desc string `json:"desc"`
}

// GoodsEditRequest 商品请求参数
type GoodsEditRequest struct {
	//主键
	Id string `json:"id"`
	//商品名称
	Name string `json:"name"`
	//商品类型
	GoodsType string `json:"goodsType"`
	//商品编号
	GoodsNumber string `json:"goodsNumber"`
	//商品价格
	Price string `json:"price"`
	//数量
	Count string `json:"count"`
	//描述
	Desc string `json:"desc"`
}
