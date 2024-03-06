package entity

import (
	"gorm.io/gorm"
	"time"
)

// TopUp 充值记录
type TopUp struct {
	//主键
	Id uint64 `gorm:"column:id;not null;autoIncrement;primaryKey;comment:'主键id'"`
	//会员id
	MemberId uint64 `gorm:"column:member_id;not null;comment:'会员id';INDEX:idx_t_top_up_member_id"`
	//充值金额
	TopUp float64 `gorm:"column:top_up;default:0;not null;comment:'充值金额'"`
	//用户当前余额
	CurrentAccount float64 `gorm:"column:current_account;default:0;not null;comment:'用户当前金额'"`
	//创建时间
	CreateTime time.Time `gorm:"column:create_time;not null;comment:'创建时间';autoCreateTime;INDEX:idx_t_top_up_create_time"`
	//更新时间
	UpdateTime time.Time `gorm:"column:update_time;not null;comment:'更新时间';autoUpdateTime;INDEX:idx_t_top_up_update_time"`
	//逻辑删除
	DeleteAt gorm.DeletedAt `gorm:"column:delete_at;comment:'逻辑删除';INDEX:idx_t_top_up_delete_at"`
}

func (TopUp) TableName() string {
	return "t_top_up"
}
