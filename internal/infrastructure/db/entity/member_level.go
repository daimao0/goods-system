package entity

import (
	"gorm.io/gorm"
	"time"
)

// MemberLevel 会员等级
type MemberLevel struct {
	//主键
	Id uint64 `gorm:"column:id;not null;autoIncrement;primaryKey;comment:'主键id'"`
	//会员等级名称
	Name string `gorm:"column:name;default:'';not null;comment:'会员名称';INDEX:idx_t_member_name"`
	//优惠折扣
	Discount float64 `gorm:"column:discount;default:0;not null;comment:'优惠折扣'"`
	//创建时间
	CreateTime time.Time `gorm:"column:create_time;not null;comment:'创建时间';autoCreateTime;INDEX:idx_t_member_create_time"`
	//更新时间
	UpdateTime time.Time `gorm:"column:update_time;not null;comment:'更新时间';autoUpdateTime;INDEX:idx_t_member_update_time"`
	//逻辑删除
	DeleteAt gorm.DeletedAt `gorm:"column:delete_at;comment:'逻辑删除';INDEX:idx_t_member_delete_at"`
}

func (MemberLevel) TableName() string {
	return "t_member_level"
}
