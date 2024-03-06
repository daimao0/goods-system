package controller

import (
	"go/types"
	"goods-system/internal/application/appmanager"
	"goods-system/internal/application/request"
	"goods-system/internal/application/vo"
	"goods-system/internal/infrastructure/common/api"
)

var goodsManager = appmanager.GetGoodsManagerInstance()

type GoodsController struct {
}

var goodsController *GoodsController

func GetGoodsController() *GoodsController {
	return goodsController
}

// GoodsGetById 根据id获得商品
func (p *GoodsController) GoodsGetById(id string) api.RespData[vo.GoodsVO] {
	goodsVO := goodsManager.GetById(id)
	return api.Success(goodsVO, "")
}

// GoodsAll 所有商品
func (p *GoodsController) GoodsAll() api.RespData[[]vo.GoodsVO] {
	var goodsVOS []vo.GoodsVO
	goodsVOS = goodsManager.All()
	return api.Success(goodsVOS, "")
}

// GoodsPage 分页
func (p *GoodsController) GoodsPage(req request.GoodsPageRequest) api.RespData[api.Page[vo.GoodsVO]] {
	page := goodsManager.Page(req)
	return api.Success[api.Page[vo.GoodsVO]](page, "")
}

// GoodsSave 保存商品
func (p *GoodsController) GoodsSave(req request.GoodsSaveRequest) api.RespData[types.Nil] {
	if req.Name == "" || req.GoodsType == "" || req.Price == "" || req.Count == "" {
		return api.Fail[types.Nil]("参数不全，检查商品名称、类型、价格、数量是否为空")
	}
	goodsManager.Save(req)
	return api.Success[types.Nil](types.Nil{}, "")
}

// GoodsDelete 删除商品
func (p *GoodsController) GoodsDelete(id string) api.RespData[types.Nil] {
	if id == "" {
		return api.Fail[types.Nil]("id不能为空")
	}
	goodsManager.Delete(id)
	return api.Success[types.Nil](types.Nil{}, "")
}

// GoodsEdit 编辑商品
func (p *GoodsController) GoodsEdit(req request.GoodsEditRequest) api.RespData[types.Nil] {
	if req.Id == "" {
		return api.Fail[types.Nil]("id不能为空，联系管理员")
	}
	if req.Name == "" || req.GoodsType == "" || req.Price == "" || req.Count == "" {
		return api.Fail[types.Nil]("参数不全，检查商品名称、类型、价格、数量是否为空")
	}
	goodsManager.Edit(req)
	return api.Success[types.Nil](types.Nil{}, "")
}
