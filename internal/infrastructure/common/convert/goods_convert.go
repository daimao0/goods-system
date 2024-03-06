package convert

import (
	"goods-system/internal/application/request"
	"goods-system/internal/application/vo"
	"goods-system/internal/infrastructure/common/utils"
	"goods-system/internal/infrastructure/db/entity"
	"goods-system/internal/service/dto"
	"goods-system/internal/service/param"
	"time"
)

// ToGoodsByParam 类型转化
func ToGoodsByParam(param param.GoodsParam) entity.Goods {
	return entity.Goods{Name: param.Name, GoodsNumber: param.GoodsNumber, GoodsType: param.GoodsType}
}

// ToGoodsByPageParam 类型转化
func ToGoodsByPageParam(param param.GoodsPageParam) entity.Goods {
	return entity.Goods{
		Name:      param.Name,
		GoodsType: param.GoodsType,
	}
}

// ToGoodsPageParam 分页查询商品参数
func ToGoodsPageParam(req request.GoodsPageRequest) param.GoodsPageParam {
	return param.GoodsPageParam{
		Name:      req.Name,
		GoodsType: req.GoodsType,
		Page:      req.Page,
		Size:      req.Size,
	}
}

// ToGoodsDTO 商品dto
func ToGoodsDTO(goods entity.Goods) dto.GoodsDTO {
	return dto.GoodsDTO{
		Id:          goods.Id,
		Name:        goods.Name,
		GoodsNumber: goods.GoodsNumber,
		GoodsType:   goods.GoodsType,
		Price:       goods.Price,
		Count:       goods.Count,
		Desc:        goods.Desc,
		CreateTime:  goods.CreateTime,
		UpdateTime:  goods.UpdateTime,
	}
}

// ToGoodsVO 类型转化
func ToGoodsVO(goodsDTO dto.GoodsDTO) vo.GoodsVO {
	return vo.GoodsVO{
		Id:          utils.ToStr(goodsDTO.Id),
		Name:        goodsDTO.Name,
		GoodsNumber: goodsDTO.GoodsNumber,
		GoodsType:   goodsDTO.GoodsType,
		Price:       utils.ToStr(goodsDTO.Price),
		Count:       utils.ToStr(goodsDTO.Count),
		Desc:        goodsDTO.Desc,
		CreateTime:  utils.ToStr(goodsDTO.CreateTime),
		UpdateTime:  utils.ToStr(goodsDTO.UpdateTime),
	}
}

// ToGoodsSaveParam 类型转化
func ToGoodsSaveParam(req request.GoodsSaveRequest) param.GoodsSaveParam {
	return param.GoodsSaveParam{
		Name:        req.Name,
		GoodsType:   req.GoodsType,
		GoodsNumber: req.GoodsNumber,
		Price:       utils.ToFloat64(req.Price),
		Count:       utils.ToInt64(req.Count),
		Desc:        req.Desc,
	}
}

// ToGoodsUpdateParam 类型转化
func ToGoodsUpdateParam(req request.GoodsEditRequest) param.GoodsUpdateParam {
	return param.GoodsUpdateParam{
		Id:          utils.ToUInt64(req.Id),
		Name:        req.Name,
		GoodsType:   req.GoodsType,
		GoodsNumber: req.GoodsNumber,
		Price:       utils.ToFloat64(req.Price),
		Count:       utils.ToInt64(req.Count),
		Desc:        req.Desc,
	}
}

// ToGoodsBySaveParam 类型转化
func ToGoodsBySaveParam(param param.GoodsSaveParam) entity.Goods {
	return entity.Goods{
		Name:        param.Name,
		GoodsNumber: param.GoodsNumber,
		GoodsType:   param.GoodsType,
		Price:       param.Price,
		Count:       param.Count,
		Desc:        param.Desc,
	}
}

// ToGoodsByUpdateParam 类型转化
func ToGoodsByUpdateParam(param param.GoodsUpdateParam) entity.Goods {
	return entity.Goods{
		Id:          param.Id,
		Name:        param.Name,
		GoodsNumber: param.GoodsNumber,
		GoodsType:   param.GoodsType,
		Price:       param.Price,
		Count:       param.Count,
		Desc:        param.Desc,
		UpdateTime:  time.Now(),
	}
}
