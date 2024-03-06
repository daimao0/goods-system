package appmanager

import (
	"goods-system/internal/application/request"
	"goods-system/internal/application/vo"
	"goods-system/internal/infrastructure/common/api"
	"goods-system/internal/infrastructure/common/convert"
	"goods-system/internal/infrastructure/common/utils"
	"goods-system/internal/service/impl"
	"goods-system/internal/service/param"
)

var goodsService = impl.GetGoodsServiceInstance()

type GoodsManager struct {
}

var goodsManager *GoodsManager

func GetGoodsManagerInstance() *GoodsManager {
	return goodsManager
}

// Page 分页
func (p *GoodsManager) Page(request request.GoodsPageRequest) api.Page[vo.GoodsVO] {
	param := convert.ToGoodsPageParam(request)
	//商品服务分页查询
	page := goodsService.Page(param)
	data := page.Data
	var GoodsVOList []vo.GoodsVO
	for _, datum := range data {
		goodsVO := convert.ToGoodsVO(datum)
		GoodsVOList = append(GoodsVOList, goodsVO)
	}
	return api.Page[vo.GoodsVO]{
		Page:  page.Page,
		Size:  page.Size,
		Total: page.Total,
		Data:  GoodsVOList,
	}
}

// Save 保存商品
func (p *GoodsManager) Save(req request.GoodsSaveRequest) {
	param := convert.ToGoodsSaveParam(req)
	goodsService.Save(param)
}

// Delete 删除商品
func (p *GoodsManager) Delete(id string) {
	goodsService.Delete(utils.ToUInt64(id))
}

// Edit 编辑商品
func (p *GoodsManager) Edit(req request.GoodsEditRequest) {
	param := convert.ToGoodsUpdateParam(req)
	goodsService.Update(param)
}

// All 获得所有商品
func (p *GoodsManager) All() []vo.GoodsVO {
	var goodsVOS []vo.GoodsVO
	goodsDTOS := goodsService.List(param.GoodsParam{})
	for _, item := range goodsDTOS {
		goodsVO := convert.ToGoodsVO(item)
		goodsVOS = append(goodsVOS, goodsVO)
	}
	return goodsVOS
}

// GetById 根据id获得商品
func (p *GoodsManager) GetById(id string) vo.GoodsVO {
	goodsDTO := goodsService.GetById(utils.ToUInt64(id))
	return convert.ToGoodsVO(goodsDTO)
}
