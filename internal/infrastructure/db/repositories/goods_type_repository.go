package repositories

import (
	"goods-system/internal/infrastructure/common/global"
	"goods-system/internal/infrastructure/db/entity"
	"time"
)

var goodsTypeInstance *GoodsTypeRepository

type GoodsTypeRepository struct {
}

func GetGoodsTypeRepositoryInstance() *GoodsTypeRepository {
	return goodsTypeInstance
}

// SelectAll 查询全部商品类型
func (p *GoodsTypeRepository) SelectAll() []entity.GoodsType {
	var goodsTypes []entity.GoodsType
	global.DB.Find(&goodsTypes)
	return goodsTypes
}

// Insert 保存
func (p *GoodsTypeRepository) Insert(goodsType *entity.GoodsType) {
	global.DB.Create(goodsType)
}

// SelectById 通过id获得商品
func (p *GoodsTypeRepository) SelectById(id uint64) entity.GoodsType {
	var goodsType entity.GoodsType
	global.DB.First(&goodsType, id)
	return goodsType
}

// DeleteById 根据id删除
func (p *GoodsTypeRepository) DeleteById(id uint64) {
	var goodsType entity.GoodsType
	global.DB.Delete(&goodsType, id)
}

// UpdateById 更新
func (p *GoodsTypeRepository) UpdateById(goodsType *entity.GoodsType) {
	flag := false
	//查询原来的数据
	old := p.SelectById(goodsType.Id)
	// 创建一个只包含有变化字段的 map
	updates := make(map[string]interface{})
	if goodsType.Name != old.Name {
		updates["name"] = goodsType.Name
		flag = true
	}
	if flag {
		updates["update_time"] = time.Now()
		global.DB.Model(goodsType).Updates(updates)
	}
}

// DeleteByName 根据商品类型名称删除商品类型
func (p *GoodsTypeRepository) DeleteByName(name string) {
	goodsType := entity.GoodsType{}
	global.DB.Where("name=?", name).Delete(&goodsType)
}

func (p *GoodsTypeRepository) SelectByName(name string) entity.GoodsType {
	goodsType := entity.GoodsType{}
	global.DB.Where("name=?", name).First(&goodsType)
	return goodsType
}

func (p *GoodsTypeRepository) BatchSave(goodsType []entity.GoodsType) {
	for _, item := range goodsType {
		global.DB.Save(&item)
	}
}
