package impl

import (
	"goods-system/internal/infrastructure/common/api"
	"goods-system/internal/infrastructure/common/convert"
	"goods-system/internal/infrastructure/common/utils"
	"goods-system/internal/infrastructure/db/entity"
	"goods-system/internal/infrastructure/db/repositories"
	"goods-system/internal/service/dto"
	"goods-system/internal/service/param"
)

type GoodsService struct {
}

var goodsService *GoodsService
var goodsRepository = repositories.GetGoodsRepositoryInstance()

func GetGoodsServiceInstance() *GoodsService {
	return goodsService
}

// Page 分页
func (p *GoodsService) Page(param param.GoodsPageParam) api.Page[dto.GoodsDTO] {
	goods := convert.ToGoodsByPageParam(param)
	page := goodsRepository.Page(&goods, param.Page, param.Size)
	//数据库查询商品列表
	goodsList := page.Data
	//var goodsDTOList []dto.GoodsDTO
	var goodsDTOList []dto.GoodsDTO
	for _, e := range goodsList {
		goodsDTO := convert.ToGoodsDTO(e)
		goodsDTOList = append(goodsDTOList, goodsDTO)
	}
	return api.Page[dto.GoodsDTO]{
		Page:  page.Page,
		Size:  page.Size,
		Total: page.Total,
		Data:  goodsDTOList,
	}
}

// Save 保存商品
func (p *GoodsService) Save(param param.GoodsSaveParam) dto.GoodsDTO {
	entity := convert.ToGoodsBySaveParam(param)
	goodsRepository.Insert(&entity)
	return convert.ToGoodsDTO(entity)
}

// Delete 删除商品
func (p *GoodsService) Delete(id uint64) {
	goodsRepository.DeleteById(id)
}

// Update 更新商品
func (p *GoodsService) Update(param param.GoodsUpdateParam) dto.GoodsDTO {
	entity := convert.ToGoodsByUpdateParam(param)
	goodsRepository.UpdateById(&entity)
	return convert.ToGoodsDTO(entity)
}

// List 获得所有商品
func (p *GoodsService) List(param param.GoodsParam) []dto.GoodsDTO {
	var goodsDTOS []dto.GoodsDTO
	goods := convert.ToGoodsByParam(param)
	goodsList := goodsRepository.Select(goods)
	for _, item := range goodsList {
		goodsDTO := convert.ToGoodsDTO(item)
		goodsDTOS = append(goodsDTOS, goodsDTO)
	}
	return goodsDTOS
}

// GetById 根据id获得商品
func (p *GoodsService) GetById(id uint64) dto.GoodsDTO {
	goods := entity.Goods{Id: utils.ToUInt64(id)}
	goodsRepository.Get(&goods)
	return convert.ToGoodsDTO(goods)
}

// ReduceCount 减少库存
func (p *GoodsService) ReduceCount(params []param.GoodsReduceParam) {
	for _, reduce := range params {
		goods := p.GetById(reduce.Id)
		updateParam := param.GoodsUpdateParam{
			Id:          goods.Id,
			Name:        goods.Name,
			GoodsNumber: goods.GoodsNumber,
			GoodsType:   goods.GoodsType,
			Price:       goods.Price,
			Count:       goods.Count - int64(reduce.Num),
			Desc:        goods.Desc,
		}
		goodsService.Update(updateParam)
	}
}
