package controller

import (
	"go/types"
	"goods-system/internal/application/appmanager"
	"goods-system/internal/application/request"
	"goods-system/internal/application/vo"
	"goods-system/internal/infrastructure/common/api"
)

type OrderController struct {
}

var orderController *OrderController

func GetOrderController() *OrderController {
	return orderController
}

// manager
var orderManager = appmanager.GetOrderManager()

// Order 下订单
func (p *OrderController) Order(request request.OrderRequest) api.RespData[types.Nil] {
	orderManager.Order(request)
	return api.Success(types.Nil{}, "")
}

// Records 订单记录
func (p *OrderController) Records(request request.OrderRecordRequest) api.RespData[[]vo.OrderVO] {
	orderVOS := orderManager.Records(request)
	return api.Success(orderVOS, "")
}

// Page 分页
func (p *OrderController) Page(request request.OrderPageRequest) api.RespData[api.Page[vo.OrderVO]] {
	page := orderManager.Page(request)
	return api.Success(page, "")
}

// Statistics 统计
func (p *OrderController) Statistics(request request.OrderRecordRequest) api.RespData[vo.OrderStatisticsVO] {
	statisticsVO := orderManager.Statistics(request)
	return api.Success(statisticsVO, "")
}
