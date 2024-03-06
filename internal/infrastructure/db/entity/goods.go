package entity

import (
	"gorm.io/gorm"
	"time"
)

type Goods struct {
	//主键
	Id uint64 `gorm:"column:id;not null;autoIncrement;primaryKey;comment:'主键id'"`
	//商品名称
	Name string `gorm:"column:name;default:'';not null;comment:'商品名称';INDEX:idx_t_goods_name"`
	//商品编号
	GoodsNumber string `gorm:"column:goods_number;default:'';not null;comment:'商品编号';INDEX:idx_t_goods_goods_number"`
	//商品类型
	GoodsType string `gorm:"column:goods_type;default:'';not null;comment:'商品类型'"`
	//商品价格
	Price float64 `gorm:"column:price;default:0;not null;comment:'商品价格'"`
	//数量
	Count int64 `gorm:"column:count;default:0;not null;comment:'商品数量'"`
	//描述
	Desc string `gorm:"column:desc;default:'';not null;comment:'商品描述'"`
	//创建时间
	CreateTime time.Time `gorm:"column:create_time;not null;comment:'创建时间';autoCreateTime;INDEX:idx_t_goods_create_time"`
	//更新时间
	UpdateTime time.Time `gorm:"column:update_time;not null;comment:'更新时间';autoUpdateTime;INDEX:idx_t_goods_update_time"`
	//逻辑删除
	DeleteAt gorm.DeletedAt `gorm:"column:delete_at;comment:'逻辑删除';INDEX:idx_t_goods_delete_at"`
}

func (Goods) TableName() string {
	return "t_goods"
}
