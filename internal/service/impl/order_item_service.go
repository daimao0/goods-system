package impl

import (
	"goods-system/internal/infrastructure/db/entity"
	"goods-system/internal/infrastructure/db/repositories"
	"goods-system/internal/service/dto"
	"goods-system/internal/service/param"
)

type OrderItemService struct {
}

var orderItemService *OrderItemService

func GetOrderItemService() *OrderItemService {
	return orderItemService
}

var orderItemRepository = repositories.GetOrderItemRepository()

// ListByIds 根据ids查询
func (p *OrderItemService) ListByIds(ids []uint64) []dto.OrderItemDTO {
	var orderItemDTOS []dto.OrderItemDTO
	orderItems := orderItemRepository.ListByIds(ids)
	for _, orderItem := range orderItems {
		orderItemDTO := ToOrderItemDTO(orderItem)
		orderItemDTOS = append(orderItemDTOS, orderItemDTO)
	}
	return orderItemDTOS
}

// BatchSave 批量保存
func (p *OrderItemService) BatchSave(params []param.OrderItemSaveParam) []dto.OrderItemDTO {
	var orderItems []entity.OrderItem
	for _, item := range params {
		orderItem := ToOrderItemBySaveParam(item)
		orderItems = append(orderItems, orderItem)
	}
	//保存
	orderItemRepository.BatchSave(orderItems)
	//类型转化
	var orderItemsDTOS []dto.OrderItemDTO
	for _, item := range orderItems {
		orderItemsDTOS = append(orderItemsDTOS, ToOrderItemDTO(item))
	}
	return orderItemsDTOS
}

// ToOrderItemDTO 类型转化
func ToOrderItemDTO(orderItem entity.OrderItem) dto.OrderItemDTO {
	goodsDTO := goodsService.GetById(orderItem.GoodsId)
	return dto.OrderItemDTO{
		Id:         orderItem.Id,
		GoodsDTO:   goodsDTO,
		RealPrice:  orderItem.RealPrice,
		Num:        orderItem.Num,
		CreateTime: orderItem.CreateTime,
	}
}

// ToOrderItemBySaveParam 类型转化
func ToOrderItemBySaveParam(param param.OrderItemSaveParam) entity.OrderItem {
	return entity.OrderItem{
		GoodsId:   param.GoodsId,
		RealPrice: param.RealPrice,
		Num:       param.Num,
	}
}
