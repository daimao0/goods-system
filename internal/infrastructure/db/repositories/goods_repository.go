package repositories

import (
	"goods-system/internal/infrastructure/common/api"
	"goods-system/internal/infrastructure/common/global"
	"goods-system/internal/infrastructure/db/entity"
	"time"
)

var goodsInstance *GoodsRepository

type GoodsRepository struct {
}

func GetGoodsRepositoryInstance() *GoodsRepository {
	return goodsInstance
}

// Page 分页
func (p *GoodsRepository) Page(goods *entity.Goods, page, size int) api.Page[entity.Goods] {
	var total int64
	var goodsList []entity.Goods
	tx := global.DB.Model(&entity.Goods{})
	if goods.GoodsType != "" {
		tx.Where("goods_type = ?", goods.GoodsType)
	}
	tx.Where("name like ?", "%"+goods.Name+"%").
		Count(&total).
		Limit(size).
		Offset((page - 1) * size).
		Find(&goodsList)

	return api.Page[entity.Goods]{
		Page:  page,
		Size:  size,
		Total: total,
		Data:  goodsList,
	}
}

// Insert 保存
func (p *GoodsRepository) Insert(goods *entity.Goods) {
	global.DB.Create(goods)
}

// SelectById 通过id获得商品
func (p *GoodsRepository) SelectById(id uint64) entity.Goods {
	var goods entity.Goods
	global.DB.First(&goods, id)
	return goods
}

// DeleteById 根据id删除
func (p *GoodsRepository) DeleteById(id uint64) {
	var goods entity.Goods
	global.DB.Delete(&goods, id)
}

// UpdateById 更新
func (p *GoodsRepository) UpdateById(goods *entity.Goods) {
	flag := false
	//查询原来的数据
	old := p.SelectById(goods.Id)
	// 创建一个只包含有变化字段的 map
	updates := make(map[string]interface{})
	if goods.Name != old.Name {
		updates["name"] = goods.Name
		flag = true
	}
	if goods.GoodsNumber != old.GoodsNumber {
		updates["goods_number"] = goods.GoodsNumber
		flag = true
	}
	if goods.GoodsType != old.GoodsType {
		updates["goods_type"] = goods.GoodsType
		flag = true
	}
	if goods.Price != old.Price {
		updates["price"] = goods.Price
		flag = true
	}
	if goods.Count != old.Count {
		updates["count"] = goods.Count
		flag = true
	}
	if goods.Desc != old.Desc {
		updates["desc"] = goods.Desc
		flag = true
	}
	if flag {
		updates["update_time"] = time.Now()
		// goods为新属性
		global.DB.Model(&goods).Updates(updates)
	}
}

// Select 获得商品
func (p *GoodsRepository) Select(goods entity.Goods) []entity.Goods {
	tx := global.DB.Model(&goods)
	var goodsList []entity.Goods
	if goods.Name != "" {
		tx.Where("name = ?", goods.Name)
	}
	if goods.GoodsNumber != "" {
		tx.Where("goods_number = ?", goods.GoodsNumber)
	}
	if goods.GoodsType != "" {
		tx.Where("goods_type = ?", goods.GoodsType)
	}
	tx.Find(&goodsList)
	return goodsList
}

// Get 获得一个商品
func (p *GoodsRepository) Get(goods *entity.Goods) {
	tx := global.DB.Model(goods)
	if goods.Id != 0 {
		tx.Where("id = ?", goods.Id)
	}
	if goods.Name != "" {
		tx.Where("name = ?", goods.Id)
	}
	if goods.GoodsNumber != "" {
		tx.Where("goods_number = ?", goods.GoodsNumber)
	}
	if goods.GoodsType != "" {
		tx.Where("goods_type = ?", goods.GoodsNumber)
	}
	if goods.Price != 0 {
		tx.Where("price =?", goods.Price)
	}
	if goods.Count != 0 {
		tx.Where("count = ?", goods.Count)
	}
	tx.First(goods)
}

// BatchSave 批量保存
func (p *GoodsRepository) BatchSave(goods []entity.Goods) {
	for _, item := range goods {
		global.DB.Save(&item)
	}
}
