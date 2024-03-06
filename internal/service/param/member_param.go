package param

// MemberSearchParam 会员分页参数
type MemberSearchParam struct {
	//会员id
	Id uint64
	//会员名称
	Name string
	//手机号
	Phone string
	//会员等级
	LevelId uint64
}

// MemberPageParam 会员分页参数
type MemberPageParam struct {
	//会员名称
	Name string
	//手机号
	Phone string
	//会员等级
	LevelId uint64
	//账户余额
	Account float64
	//当前页
	Page int
	//页面大小
	Size int
}

// MemberSaveParam 会员保存参数
type MemberSaveParam struct {
	//会员名称
	Name string
	//手机号
	Phone string
	//会员等级
	LevelId uint64
	//账户余额
	Account float64
}

// MemberUpdateParam 会员保存参数
type MemberUpdateParam struct {
	//主键
	Id uint64
	//会员名称
	Name string
	//手机号
	Phone string
	//会员等级
	LevelId uint64
	//账户余额
	Account float64
}
