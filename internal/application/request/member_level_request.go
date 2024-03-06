package request

// MemberLevelSaveRequest 会员保存参数
type MemberLevelSaveRequest struct {
	//会员等级名称
	Name string `json:"name"`
	//折扣
	Discount string `json:"discount"`
}

// MemberLevelUpdateRequest 会员更新参数
type MemberLevelUpdateRequest struct {
	//主键
	Id string `json:"id"`
	//会员等级名称
	Name string `json:"name"`
	//优惠折扣
	Discount string `json:"discount"`
}
