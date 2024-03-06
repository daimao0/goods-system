package controller

import (
	"go/types"
	"goods-system/internal/application/appmanager"
	"goods-system/internal/application/request"
	"goods-system/internal/application/vo"
	"goods-system/internal/infrastructure/common/api"
)

var goodsTypeManager = appmanager.GetGoodsTypeManagerInstance()

type GoodsTypeController struct {
}

var goodsTypeController *GoodsTypeController

func GetGoodsTypeController() *GoodsTypeController {
	return goodsTypeController
}

// SelectAll 查询全部商品类型
func (p *GoodsTypeController) SelectAll() api.RespData[[]vo.GoodsTypeVO] {
	all := goodsTypeManager.SelectAll()
	return api.Success(all, "")
}

// DeleteByName 通过商品类型名称删除
func (p *GoodsTypeController) DeleteByName(request request.GoodsTypeDeleteRequest) api.RespData[types.Nil] {
	goodsTypeManager.DeleteByName(request.Name)
	return api.Success[types.Nil](types.Nil{}, "")
}

// Save 保存商品类型参数
func (p *GoodsTypeController) Save(request request.GoodsTypeSaveRequest) api.RespData[types.Nil] {
	if request.Name == "" {
		return api.Fail[types.Nil]("类型名称不能为空")
	}
	//查询类型是否已存在
	goodsTypeDTO := goodsTypeManager.QueryByName(request.Name)
	if goodsTypeDTO.Id != 0 {
		return api.Fail[types.Nil]("商品名称" + request.Name + "已存在")
	}

	goodsTypeManager.Save(request)
	return api.Success[types.Nil](types.Nil{}, "")
}

func (p *GoodsTypeController) Update(request request.GoodsTypeEditRequest) api.RespData[types.Nil] {
	goodsTypeManager.Edit(request)
	return api.Success[types.Nil](types.Nil{}, "")
}
