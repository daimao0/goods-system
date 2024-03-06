package impl

import (
	"goods-system/internal/infrastructure/common/api"
	"goods-system/internal/infrastructure/common/utils"
	"goods-system/internal/infrastructure/db/entity"
	"goods-system/internal/infrastructure/db/query"
	"goods-system/internal/infrastructure/db/repositories"
	"goods-system/internal/service/dto"
	"goods-system/internal/service/param"
	"time"
)

type OrderService struct {
}

var orderService *OrderService

func GetOrderService() *OrderService {
	return orderService
}

var orderRepository = repositories.GetOrderRepository()

// GetById 获得订单
func (p *OrderService) GetById(id uint64) dto.OrderDTO {
	order := entity.Order{Id: id}
	orderRepository.Get(&order)
	//类型转化
	orderDTO := ToOrderDTO(order)
	return orderDTO
}

// Save 保存订单
func (p *OrderService) Save(param param.OrderSaveParam) {
	order := ToOrderBySaveParam(param)
	orderRepository.Insert(&order)
}

// List 查询订单
func (p *OrderService) List(param param.OrderSearchParam) []dto.OrderDTO {
	orderQuery := query.OrderQuery{MemberId: param.MemberId, StartTime: param.StartTime, EndTime: param.EndTime}
	orders := orderRepository.List(orderQuery)
	var orderDTOS []dto.OrderDTO
	for _, item := range orders {
		orderDTOS = append(orderDTOS, ToOrderDTO(item))
	}
	return orderDTOS
}

// Page 分页
func (p *OrderService) Page(param param.OrderPageParam) api.Page[dto.OrderDTO] {
	orderPageQuery := query.OrderPageQuery{
		MemberId:  param.MemberId,
		StartTime: param.StartTime,
		EndTime:   param.EndTime,
		Page:      param.Page,
		Size:      param.Size,
	}
	page := orderRepository.Page(&orderPageQuery)
	var orders []dto.OrderDTO
	for _, e := range page.Data {
		orderDTO := ToOrderDTO(e)
		orders = append(orders, orderDTO)
	}
	return api.Page[dto.OrderDTO]{
		Page:  page.Page,
		Size:  page.Size,
		Total: page.Total,
		Data:  orders,
	}
}

// Statistics 统计订单
func (p *OrderService) Statistics(searchParam param.OrderSearchParam) dto.OrderStatisticsDTO {
	var orderStatisticsDTO dto.OrderStatisticsDTO
	orderDTOS := p.List(searchParam)
	for _, orderDTO := range orderDTOS {
		orderStatisticsDTO.OrderCount++
		for _, orderItem := range orderDTO.OrderItems {
			orderStatisticsDTO.GoodsPrice += orderItem.GoodsDTO.Price
		}
		orderStatisticsDTO.MemberPay += orderDTO.Pay
	}
	orderStatisticsDTO.PreferentialPrice = orderStatisticsDTO.GoodsPrice - orderStatisticsDTO.MemberPay
	return orderStatisticsDTO
}

// ToOrderBySearchParam 类型转化
func ToOrderBySearchParam(param param.OrderSearchParam) entity.Order {
	return entity.Order{MemberId: param.MemberId}
}

// ToOrderDTO 类型转化
func ToOrderDTO(order entity.Order) dto.OrderDTO {
	itemIds := utils.ToUInt64Arr(order.OrderItemIds)
	orderItems := orderItemService.ListByIds(itemIds)
	memberDTO := memberService.Get(param.MemberSearchParam{Id: order.MemberId})
	return dto.OrderDTO{
		Id:         order.Id,
		Member:     memberDTO,
		OrderItems: orderItems,
		Pay:        order.Pay,
		CreateTime: order.CreateTime,
	}
}

// ToOrderBySaveParam 类型转化
func ToOrderBySaveParam(param param.OrderSaveParam) entity.Order {
	return entity.Order{
		MemberId:     param.MemberId,
		OrderItemIds: utils.ToStrByUInt64(param.OrderItemIds),
		Pay:          param.Pay,
		CreateTime:   time.Now(),
		UpdateTime:   time.Now(),
	}
}
