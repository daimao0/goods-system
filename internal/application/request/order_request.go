package request

// OrderRequest 订单保存请求
type OrderRequest struct {
	//总价
	Pay string `json:"pay"`
	//会员id
	MemberId string `json:"memberId"`
	//订单
	OrderItem []OrderItemRequest `json:"orderItem"`
}

// OrderRecordRequest 订单记录请求
type OrderRecordRequest struct {
	//会员id
	MemberId string
	//开始时间
	StartTime string `json:"startTime"`
	//结束时间
	EndTime string `json:"endTime"`
}

// OrderPageRequest 订单分页参数
type OrderPageRequest struct {
	//会员id
	MemberId string `json:"memberId"`
	//开始时间
	StartTime string `json:"startTime"`
	//结束时间
	EndTime string `json:"endTime"`
	//当前页
	Page int `json:"page"`
	//页面大小
	Size int `json:"size"`
}
