package vo

type MemberVO struct {
	//主键
	Id string `json:"id"`
	//会员名称
	Name string `json:"name"`
	//手机号
	Phone string `json:"phone"`
	//会员等级
	Level MemberLevelVO `json:"level"`
	//优惠折扣
	Discount string `json:"discount"`
	//账户余额
	Account string `json:"account"`
	//创建时间
	CreateTime string `json:"createTime"`
	//更新时间
	UpdateTime string `json:"updateTime"`
}
