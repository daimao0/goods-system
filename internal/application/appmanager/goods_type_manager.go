package appmanager

import (
	"goods-system/internal/application/request"
	"goods-system/internal/application/vo"
	"goods-system/internal/infrastructure/common/convert"
	"goods-system/internal/infrastructure/common/utils"
	"goods-system/internal/service/dto"
	"goods-system/internal/service/impl"
)

var goodsTypeService = impl.GetGoodsTypeService()

type GoodsTypeManager struct {
}

var goodsTypeManager *GoodsTypeManager

func GetGoodsTypeManagerInstance() *GoodsTypeManager {
	return goodsTypeManager
}

// SelectAll 查询全部商品类型
func (p *GoodsTypeManager) SelectAll() []vo.GoodsTypeVO {
	all := goodsTypeService.SelectAll()
	var goodsTypeVOS []vo.GoodsTypeVO
	for _, goodsTypeDTO := range all {
		goodsTypeVO := convert.ToGoodsTypeVO(goodsTypeDTO)
		goodsTypeVOS = append(goodsTypeVOS, goodsTypeVO)
	}
	return goodsTypeVOS
}

// Save 保存商品类型
func (p *GoodsTypeManager) Save(req request.GoodsTypeSaveRequest) {
	param := convert.ToGoodsTypeSaveParamBySaveRequest(req)
	goodsTypeService.Save(param)
}

// DeleteById 通过id删除商品类型
func (p *GoodsTypeManager) DeleteById(id string) {
	goodsTypeService.DeleteById(utils.ToUInt64(id))
}

// DeleteByName 通过商品类型名称删除商品类型
func (p *GoodsTypeManager) DeleteByName(name string) {
	goodsTypeService.DeleteByName(name)
}

// Edit 编辑商品类型
func (p *GoodsTypeManager) Edit(req request.GoodsTypeEditRequest) {
	param := convert.ToGoodsTypeUpdateParamByEditRequest(req)
	goodsTypeService.UpdateById(param)
}

func (p *GoodsTypeManager) QueryByName(name string) dto.GoodsTypeDTO {
	// QueryByName 通过名称查询类型
	return goodsTypeService.QueryByName(name)
}
