package entity

import (
	"gorm.io/gorm"
	"time"
)

// OrderItem 账单
type OrderItem struct {
	//主键
	Id uint64 `gorm:"column:id;not null;autoIncrement;primaryKey;comment:'主键id'"`
	//商品id
	GoodsId uint64 `gorm:"column:goods_id;not null;comment:'商品id';INDEX:idx_t_order_item_goods_id"`
	//实际单价
	RealPrice float64 `gorm:"column:real_price;default:0;not null;comment:'实际支付'"`
	//数量
	Num int `gorm:"column:num;default:0;not null;comment:'下单数量'"`
	//创建时间
	CreateTime time.Time `gorm:"column:create_time;not null;comment:'创建时间';autoCreateTime;INDEX:idx_t_order_item_create_time"`
	//更新时间
	UpdateTime time.Time `gorm:"column:update_time;not null;comment:'更新时间';autoUpdateTime;INDEX:idx_t_order_item_update_time"`
	//逻辑删除
	DeleteAt gorm.DeletedAt `gorm:"column:delete_at;comment:'逻辑删除';INDEX:idx_t_order_item_delete_at"`
}

func (OrderItem) TableName() string {
	return "t_order_item"
}
