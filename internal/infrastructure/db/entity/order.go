package entity

import (
	"gorm.io/gorm"
	"time"
)

// Order 账单
type Order struct {
	//主键
	Id uint64 `gorm:"column:id;not null;autoIncrement;primaryKey;comment:'主键id'"`
	//会员id
	MemberId uint64 `gorm:"column:member_id;not null;default:0;comment:'会员id';INDEX:idx_t_order_member_id"`
	//商品ids
	OrderItemIds string `gorm:"column:order_item_ids;not null;comment:'商品id用,隔开';INDEX:idx_t_order_item_ids"`
	//实际支付价格
	Pay float64 `gorm:"column:pay;default:0;not null;comment:'实际支付价格'"`
	//创建时间
	CreateTime time.Time `gorm:"column:create_time;not null;comment:'创建时间';autoCreateTime;INDEX:idx_t_order_create_time"`
	//更新时间
	UpdateTime time.Time `gorm:"column:update_time;not null;comment:'更新时间';autoUpdateTime;INDEX:idx_t_order_update_time"`
	//逻辑删除
	DeleteAt gorm.DeletedAt `gorm:"column:delete_at;comment:'逻辑删除';INDEX:idx_t_order_delete_at"`
}

func (Order) TableName() string {
	return "t_order"
}
