package vo

type GoodsVO struct {
	//主键
	Id string `json:"id"`
	//商品名称
	Name string `json:"name"`
	//商品编号
	GoodsNumber string `json:"goodsNumber"`
	//商品类型
	GoodsType string `json:"goodsType"`
	//商品价格
	Price string `json:"price"`
	//数量
	Count string `json:"count"`
	//描述
	Desc string `json:"desc"`
	//创建时间
	CreateTime string `json:"createTime"`
	//更新时间
	UpdateTime string `json:"updateTime"`
}
