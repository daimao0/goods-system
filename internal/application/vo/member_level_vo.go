package vo

type MemberLevelVO struct {
	//主键
	Id string `json:"id"`
	//会员名称
	Name string `json:"name"`
	//优惠折扣
	Discount string `json:"discount"`
	//创建时间
	CreateTime string `json:"createTime"`
	//更新时间
	UpdateTime string `json:"updateTime"`
}
