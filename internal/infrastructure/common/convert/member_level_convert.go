package convert

import (
	"goods-system/internal/application/request"
	"goods-system/internal/application/vo"
	"goods-system/internal/infrastructure/common/utils"
	"goods-system/internal/infrastructure/db/entity"
	"goods-system/internal/service/dto"
	"goods-system/internal/service/param"
)

func ToMemberLevelBySaveParam(param param.MemberLevelSaveParam) entity.MemberLevel {
	return entity.MemberLevel{
		Name:     param.Name,
		Discount: param.Discount,
	}
}

func ToMemberLevelByUpdateParam(param param.MemberLevelUpdateParam) entity.MemberLevel {
	return entity.MemberLevel{
		Id:       param.Id,
		Name:     param.Name,
		Discount: param.Discount,
	}
}

func ToMemberLevelDTO(memberLevel entity.MemberLevel) dto.MemberLevelDTO {
	return dto.MemberLevelDTO{
		Id:       memberLevel.Id,
		Name:     memberLevel.Name,
		Discount: memberLevel.Discount,
	}
}

func ToMemberLevelSaveParamByRequest(request request.MemberLevelSaveRequest) param.MemberLevelSaveParam {
	return param.MemberLevelSaveParam{
		Name:     request.Name,
		Discount: utils.ToFloat64(request.Discount),
	}
}

func ToMemberLevelUpdateParamByRequest(request request.MemberLevelUpdateRequest) param.MemberLevelUpdateParam {
	return param.MemberLevelUpdateParam{
		Id:       utils.ToUInt64(request.Id),
		Name:     request.Name,
		Discount: utils.ToFloat64(request.Discount),
	}
}

func ToMemberLevelVO(memberLevel dto.MemberLevelDTO) vo.MemberLevelVO {
	return vo.MemberLevelVO{
		Id:       utils.ToStr(memberLevel.Id),
		Name:     memberLevel.Name,
		Discount: utils.ToStr(memberLevel.Discount),
	}
}
