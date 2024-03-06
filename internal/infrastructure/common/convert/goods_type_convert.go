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

// ToGoodsTypeSaveParamBySaveRequest 类型转化
func ToGoodsTypeSaveParamBySaveRequest(request request.GoodsTypeSaveRequest) param.GoodsTypeSaveParam {
	return param.GoodsTypeSaveParam{
		Name: request.Name,
	}
}

func ToGoodsTypeUpdateParamByEditRequest(request request.GoodsTypeEditRequest) param.GoodsTypeUpdateParam {
	return param.GoodsTypeUpdateParam{
		Id:   utils.ToUInt64(request.Id),
		Name: request.Name,
	}
}

// ToGoodsTypeBySaveParam 类型转化
func ToGoodsTypeBySaveParam(param param.GoodsTypeSaveParam) entity.GoodsType {
	return entity.GoodsType{
		Name: param.Name,
	}
}

// ToGoodsTypeByUpdateParam 类型转化
func ToGoodsTypeByUpdateParam(param param.GoodsTypeUpdateParam) entity.GoodsType {
	return entity.GoodsType{
		Id:         param.Id,
		Name:       param.Name,
		UpdateTime: time.Now(),
	}
}

// ToGoodsTypeDTO 类型转化
func ToGoodsTypeDTO(goodsType entity.GoodsType) dto.GoodsTypeDTO {
	return dto.GoodsTypeDTO{
		Id:         goodsType.Id,
		Name:       goodsType.Name,
		CreateTime: goodsType.CreateTime,
		UpdateTime: goodsType.UpdateTime,
	}
}

// ToGoodsTypeVO 类型转化
func ToGoodsTypeVO(goodsTypeDTO dto.GoodsTypeDTO) vo.GoodsTypeVO {
	return vo.GoodsTypeVO{
		Id:         utils.ToStr(goodsTypeDTO.Id),
		Name:       goodsTypeDTO.Name,
		CreateTime: utils.ToStr(goodsTypeDTO.CreateTime),
		UpdateTime: utils.ToStr(goodsTypeDTO.UpdateTime),
	}
}
