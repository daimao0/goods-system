package param

// MemberLevelSaveParam 会员保存参数
type MemberLevelSaveParam struct {
	//会员等级
	Name string
	//折扣
	Discount float64
}

// MemberLevelUpdateParam 会员保存参数
type MemberLevelUpdateParam struct {
	//主键
	Id uint64
	//会员等级
	Name string
	//折扣
	Discount float64
}
