package dto

import "time"

type GoodsDTO struct {
	//主键
	Id uint64
	//商品名称
	Name string
	//商品编号
	GoodsNumber string
	//商品类型
	GoodsType string
	//商品价格
	Price float64
	//数量
	Count int64
	//描述
	Desc string
	//创建时间
	CreateTime time.Time
	//更新时间
	UpdateTime time.Time
}
