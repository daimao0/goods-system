package appmanager

import (
	"goods-system/internal/application/request"
	"goods-system/internal/application/vo"
	"goods-system/internal/infrastructure/common/api"
	"goods-system/internal/infrastructure/common/convert"
	"goods-system/internal/infrastructure/common/utils"
	"goods-system/internal/service/dto"
	"goods-system/internal/service/impl"
	"goods-system/internal/service/param"
)

type OrderManager struct {
}

var orderManager *OrderManager

func GetOrderManager() *OrderManager {
	return orderManager
}

// 服务
var orderService = impl.GetOrderService()
var orderItemService = impl.GetOrderItemService()

// Order 下订单
func (p *OrderManager) Order(request request.OrderRequest) {
	var orderItemSaveParams []param.OrderItemSaveParam
	for _, item := range request.OrderItem {
		orderItemSaveParams = append(orderItemSaveParams, ToOrderItemSaveParam(item))
	}
	//保存订单每个子项
	orderItemDTOS := orderItemService.BatchSave(orderItemSaveParams)
	//保存订单
	var orderItemIds []uint64
	for _, item := range orderItemDTOS {
		orderItemIds = append(orderItemIds, item.Id)
	}
	orderService.Save(ToOrderSaveParam(request, orderItemIds))
	//扣库存
	reduceParams := ToGoodsReduceParams(request)
	goodsService.ReduceCount(reduceParams)
	//扣会员的存款
	memberService.DeductMoney(utils.ToUInt64(request.MemberId), utils.ToFloat64(request.Pay))
}

// Records 订单记录
func (p *OrderManager) Records(request request.OrderRecordRequest) []vo.OrderVO {
	var orderVOS []vo.OrderVO
	orderDTOS := orderService.List(ToOrderRecordRequest(request))
	for _, orderDTO := range orderDTOS {
		orderVOS = append(orderVOS, ToOrderVO(orderDTO))
	}
	return orderVOS
}

// Page 分页
func (p *OrderManager) Page(request request.OrderPageRequest) api.Page[vo.OrderVO] {
	page := orderService.Page(ToOrderPageParam(request))
	var orders []vo.OrderVO
	for _, datum := range page.Data {
		orders = append(orders, ToOrderVO(datum))
	}
	return api.Page[vo.OrderVO]{
		Page:  page.Page,
		Size:  page.Size,
		Total: page.Total,
		Data:  orders,
	}
}

// Statistics 统计
func (p *OrderManager) Statistics(request request.OrderRecordRequest) vo.OrderStatisticsVO {
	searchParam := param.OrderSearchParam{
		MemberId:  utils.ToUInt64(request.MemberId),
		StartTime: utils.ToTime(request.StartTime),
		EndTime:   utils.ToTime(request.EndTime),
	}
	statisticsDTO := orderService.Statistics(searchParam)
	return vo.OrderStatisticsVO{
		OrderCount:        utils.ToStr(statisticsDTO.OrderCount),
		GoodsPrice:        utils.ToStr(statisticsDTO.GoodsPrice),
		MemberPay:         utils.ToStr(statisticsDTO.MemberPay),
		PreferentialPrice: utils.ToStr(statisticsDTO.PreferentialPrice),
	}
}

// ToOrderVO 类型转化
func ToOrderVO(orderDTO dto.OrderDTO) vo.OrderVO {
	var orderItems []vo.OrderItemVO
	for _, item := range orderDTO.OrderItems {
		orderItems = append(orderItems, ToOrderItemVO(item))
	}
	return vo.OrderVO{
		Id:         utils.ToStr(orderDTO.Id),
		Pay:        utils.ToStr(orderDTO.Pay),
		OrderItems: orderItems,
		Member:     convert.ToMemberVOByDTO(orderDTO.Member),
		CreateTime: utils.ToStr(orderDTO.CreateTime),
	}
}

// ToOrderPageParam 类型转化
func ToOrderPageParam(request request.OrderPageRequest) param.OrderPageParam {
	return param.OrderPageParam{
		MemberId:  utils.ToUInt64(request.MemberId),
		StartTime: utils.ToTime(request.StartTime),
		EndTime:   utils.ToTime(request.EndTime),
		Page:      request.Page,
		Size:      request.Size,
	}
}

// ToOrderRecordRequest 类型转化
func ToOrderRecordRequest(request request.OrderRecordRequest) param.OrderSearchParam {
	return param.OrderSearchParam{MemberId: utils.ToUInt64(request.MemberId)}
}

// ToOrderSaveParam 类型转化
func ToOrderSaveParam(request request.OrderRequest, orderItemIds []uint64) param.OrderSaveParam {
	return param.OrderSaveParam{
		MemberId:     utils.ToUInt64(request.MemberId),
		OrderItemIds: orderItemIds,
		Pay:          utils.ToFloat64(request.Pay),
	}
}

// ToOrderItemSaveParam 类型转化
func ToOrderItemSaveParam(request request.OrderItemRequest) param.OrderItemSaveParam {
	return param.OrderItemSaveParam{
		GoodsId:   utils.ToUInt64(request.GoodsId),
		RealPrice: utils.ToFloat64(request.RealPrice),
		Num:       request.Num,
	}
}

// ToGoodsReduceParams 类型转化
func ToGoodsReduceParams(request request.OrderRequest) []param.GoodsReduceParam {
	var goodsReduces []param.GoodsReduceParam
	for _, item := range request.OrderItem {
		reduce := param.GoodsReduceParam{
			Id:  utils.ToUInt64(item.GoodsId),
			Num: item.Num,
		}
		goodsReduces = append(goodsReduces, reduce)
	}
	return goodsReduces
}

// ToOrderItemVO 类型转化
func ToOrderItemVO(itemDTO dto.OrderItemDTO) vo.OrderItemVO {
	return vo.OrderItemVO{
		Id:         utils.ToStr(itemDTO.Id),
		Goods:      convert.ToGoodsVO(itemDTO.GoodsDTO),
		RealPrice:  utils.ToStr(itemDTO.RealPrice),
		Num:        itemDTO.Num,
		CreateTime: utils.ToStr(itemDTO.CreateTime),
	}
}
