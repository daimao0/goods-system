package repositories

import (
	"goods-system/internal/infrastructure/common/global"
	"goods-system/internal/infrastructure/db/entity"
)

type OrderItemRepository struct {
}

var orderItemRepository *OrderItemRepository

func GetOrderItemRepository() *OrderItemRepository {
	return orderItemRepository
}

// Get 查询单个订单
func (p *OrderItemRepository) Get(OrderItem *entity.OrderItem) {
	global.DB.First(OrderItem, OrderItem.Id)
}

// ListByIds 根据id获得订单
func (p *OrderItemRepository) ListByIds(ids []uint64) []entity.OrderItem {
	var orderItems []entity.OrderItem
	global.DB.Find(&orderItems, ids)
	return orderItems
}

// Insert 插入
func (p *OrderItemRepository) Insert(OrderItem *entity.OrderItem) {
	global.DB.Save(&OrderItem)
}

// BatchSave 批量保存
func (p *OrderItemRepository) BatchSave(orderItems []entity.OrderItem) {
	for _, item := range orderItems {
		global.DB.Save(&item)
	}
}

// SelectAll 查询订单子项
func (p *OrderItemRepository) SelectAll() []entity.OrderItem {
	var orderItems []entity.OrderItem
	global.DB.Find(&orderItems)
	return orderItems
}
