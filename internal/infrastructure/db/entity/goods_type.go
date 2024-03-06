package entity

import (
	"gorm.io/gorm"
	"time"
)

type GoodsType struct {
	//主键
	Id uint64 `gorm:"column:id;not null;autoIncrement;primaryKey;comment:'主键id'"`
	//商品名称
	Name string `gorm:"column:name;default:'';not null;comment:'商品名称';INDEX:idx_t_goods_type_name"`
	//创建时间
	CreateTime time.Time `gorm:"column:create_time;not null;comment:'创建时间';autoCreateTime;INDEX:idx_t_goods_type_create_time"`
	//更新时间
	UpdateTime time.Time `gorm:"column:update_time;not null;comment:'更新时间';autoUpdateTime;INDEX:idx_t_goods_type_update_time"`
	//逻辑删除
	DeleteAt gorm.DeletedAt `gorm:"column:delete_at;comment:'逻辑删除';INDEX:idx_t_goods_type_delete_at"`
}

func (GoodsType) TableName() string {
	return "t_goods_type"
}
