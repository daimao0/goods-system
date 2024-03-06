package vo

type GoodsTypeVO struct {
	//主键
	Id string `json:"id"`
	//商品名称
	Name string `json:"name"`
	//创建时间
	CreateTime string `json:"createTime"`
	//更新时间
	UpdateTime string `json:"updateTime"`
}
