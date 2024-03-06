package impl

import (
	"goods-system/internal/infrastructure/common/convert"
	"goods-system/internal/infrastructure/db/repositories"
	"goods-system/internal/service/dto"
	"goods-system/internal/service/param"
)

type GoodsTypeService struct {
}

var goodsTypeService *GoodsTypeService
var goodsTypeRepository = repositories.GetGoodsTypeRepositoryInstance()

func GetGoodsTypeService() *GoodsTypeService {
	return goodsTypeService
}

// Save 保存商品类型
func (p *GoodsTypeService) Save(param param.GoodsTypeSaveParam) {
	goodsType := convert.ToGoodsTypeBySaveParam(param)
	goodsTypeRepository.Insert(&goodsType)
}

// SelectAll 查询所有商品类型
func (p *GoodsTypeService) SelectAll() []dto.GoodsTypeDTO {
	var goodsTypeDTOS []dto.GoodsTypeDTO
	all := goodsTypeRepository.SelectAll()
	for _, goodsType := range all {
		goodsTypeDTO := convert.ToGoodsTypeDTO(goodsType)
		goodsTypeDTOS = append(goodsTypeDTOS, goodsTypeDTO)
	}
	return goodsTypeDTOS
}

// UpdateById 通过id更新商品类型
func (p *GoodsTypeService) UpdateById(param param.GoodsTypeUpdateParam) {
	goodsType := convert.ToGoodsTypeByUpdateParam(param)
	goodsTypeRepository.UpdateById(&goodsType)
}

// DeleteById 删除商品类型
func (p *GoodsTypeService) DeleteById(id uint64) {
	goodsTypeRepository.DeleteById(id)
}

// DeleteByName 根据商品类型名称删除
func (p *GoodsTypeService) DeleteByName(name string) {
	goodsTypeRepository.DeleteByName(name)
}

// QueryByName 通过名称查询类型
func (p *GoodsTypeService) QueryByName(name string) dto.GoodsTypeDTO {
	goodsType := goodsTypeRepository.SelectByName(name)
	return convert.ToGoodsTypeDTO(goodsType)
}
