package dto

import "time"

type GoodsTypeDTO struct {
	//主键
	Id uint64
	//商品名称
	Name string
	//创建时间
	CreateTime time.Time
	//更新时间
	UpdateTime time.Time
}
