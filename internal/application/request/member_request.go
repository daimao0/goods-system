package request

// MemberSearchRequest 会员分页参数
type MemberSearchRequest struct {
	//会员id
	Id string `json:"id"`
	//会员名称
	Name string `json:"name"`
	//手机号
	Phone string `json:"phone"`
	//会员等级
	LevelId string `json:"levelId"`
}

// MemberPageRequest 会员分页参数
type MemberPageRequest struct {
	//会员名称
	Name string `json:"name"`
	//手机号
	Phone string `json:"phone"`
	//会员等级
	LevelId string `json:"levelId"`
	//账户余额
	Account string `json:"account"`
	//当前页
	Page int `json:"page"`
	//页面大小
	Size int `json:"size"`
}

// MemberSaveRequest 会员保存参数
type MemberSaveRequest struct {
	//会员名称
	Name string `json:"name"`
	//手机号
	Phone string `json:"phone"`
	//会员等级
	LevelId string `json:"levelId"`
	//账户余额
	Account string `json:"account"`
}

// MemberUpdateRequest 会员更新参数
type MemberUpdateRequest struct {
	//主键
	Id string `json:"id"`
	//会员名称
	Name string `json:"name"`
	//手机号
	Phone string `json:"phone"`
	//会员等级
	LevelId string `json:"levelId"`
	//账户余额
	Account string `json:"account"`
	//优惠折扣
	Discount string `json:"discount"`
}
