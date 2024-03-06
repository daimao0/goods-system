package repositories

import (
	"fmt"
	"goods-system/internal/infrastructure/common/api"
	"goods-system/internal/infrastructure/common/global"
	"goods-system/internal/infrastructure/common/utils"
	"goods-system/internal/infrastructure/db/entity"
	"goods-system/internal/infrastructure/db/query"
)

type OrderRepository struct {
}

var orderRepository *OrderRepository

func GetOrderRepository() *OrderRepository {
	return orderRepository
}

// Get 查询单个订单
func (p *OrderRepository) Get(order *entity.Order) {
	global.DB.First(order, order.Id)
}

// Page 分页
func (p *OrderRepository) Page(query *query.OrderPageQuery) api.Page[entity.Order] {
	var total int64
	var orders []entity.Order
	tx := global.DB.Model(&entity.Order{})
	if query.MemberId != 0 {
		tx.Where("member_id = ?", query.MemberId)
	}
	if !(query.StartTime.IsZero() && query.EndTime.IsZero()) {
		tx.Where("create_time between ? and ?", utils.ToStr(query.StartTime), utils.ToStr(query.EndTime))
	}
	tx.Order("create_time desc")
	tx.Count(&total).
		Limit(query.Size).
		Offset((query.Page - 1) * query.Size).
		Find(&orders)
	return api.Page[entity.Order]{
		Page:  query.Page,
		Size:  query.Size,
		Total: total,
		Data:  orders,
	}
}

// List 获得多个订单
func (p *OrderRepository) List(query query.OrderQuery) []entity.Order {
	var orders []entity.Order
	tx := global.DB.Model(&entity.Order{})
	if query.MemberId != 0 {
		tx.Where("member_id = ?", query.MemberId)
	}
	if !(query.StartTime.IsZero() && query.EndTime.IsZero()) {
		fmt.Println("true")
		tx.Where("create_time between ? and ?", utils.ToStr(query.StartTime), utils.ToStr(query.EndTime))
	}
	tx.Find(&orders)
	return orders
}

// Insert 插入
func (p *OrderRepository) Insert(order *entity.Order) {
	fmt.Println(order)
	global.DB.Save(order)
}

// BatchSave 批量保存
func (p *OrderRepository) BatchSave(order []entity.Order) {
	for _, item := range order {
		global.DB.Save(&item)
	}
}
