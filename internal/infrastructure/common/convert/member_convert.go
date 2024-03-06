package convert

import (
	"goods-system/internal/application/request"
	"goods-system/internal/application/vo"
	"goods-system/internal/infrastructure/common/utils"
	"goods-system/internal/infrastructure/db/entity"
	"goods-system/internal/service/dto"
	"goods-system/internal/service/param"
)

// ToMemberBySearchParam 类型转化
func ToMemberBySearchParam(param param.MemberSearchParam) entity.Member {
	return entity.Member{
		Id:      param.Id,
		Name:    param.Name,
		Phone:   param.Phone,
		LevelId: param.LevelId,
	}
}
func ToMemberSearchParamByRequest(request request.MemberSearchRequest) param.MemberSearchParam {
	return param.MemberSearchParam{
		Id:      utils.ToUInt64(request.Id),
		Name:    request.Name,
		Phone:   request.Phone,
		LevelId: utils.ToUInt64(request.LevelId),
	}
}

// ToMemberPageParamByRequest 类型转化
func ToMemberPageParamByRequest(request request.MemberPageRequest) param.MemberPageParam {
	return param.MemberPageParam{
		Name:    request.Name,
		Phone:   request.Phone,
		LevelId: utils.ToUInt64(request.LevelId),
		Account: utils.ToFloat64(request.Account),
		Page:    request.Page,
		Size:    request.Size,
	}
}

// ToMemberByPageParam 类型转化
func ToMemberByPageParam(param param.MemberPageParam) entity.Member {
	return entity.Member{
		Name:    param.Name,
		LevelId: param.LevelId,
		Phone:   param.Phone,
		Account: param.Account,
	}
}

// ToMemberBySaveParam 类型转化
func ToMemberBySaveParam(param param.MemberSaveParam) entity.Member {
	return entity.Member{
		Name:    param.Name,
		Phone:   param.Phone,
		LevelId: param.LevelId,
		Account: param.Account,
	}
}

// ToMemberByUpdateParam 类型转化
func ToMemberByUpdateParam(param param.MemberUpdateParam) entity.Member {
	return entity.Member{
		Id:      param.Id,
		Name:    param.Name,
		Phone:   param.Phone,
		LevelId: param.LevelId,
		Account: param.Account,
	}
}

// ToMemberDTOByEntity 类型转化
func ToMemberDTOByEntity(member entity.Member) dto.MemberDTO {
	return dto.MemberDTO{
		Id:         member.Id,
		Name:       member.Name,
		Phone:      member.Phone,
		Account:    member.Account,
		CreateTime: member.CreateTime,
		UpdateTime: member.UpdateTime,
	}
}

// ToMemberVOByDTO 类型转化
func ToMemberVOByDTO(dto dto.MemberDTO) vo.MemberVO {
	return vo.MemberVO{
		Id:         utils.ToStr(dto.Id),
		Name:       dto.Name,
		Phone:      dto.Phone,
		Level:      ToMemberLevelVO(dto.Level),
		Discount:   utils.ToStr(dto.Discount),
		Account:    utils.ToStr(dto.Account),
		CreateTime: utils.ToStr(dto.CreateTime),
		UpdateTime: utils.ToStr(dto.UpdateTime),
	}
}

// ToMemberSaveParamByRequest 类型转化
func ToMemberSaveParamByRequest(request request.MemberSaveRequest) param.MemberSaveParam {
	return param.MemberSaveParam{
		Name:    request.Name,
		Phone:   request.Phone,
		Account: utils.ToFloat64(request.Account),
		LevelId: utils.ToUInt64(request.LevelId),
	}
}

// ToMemberUpdateParamByRequest 类型转化
func ToMemberUpdateParamByRequest(request request.MemberUpdateRequest) param.MemberUpdateParam {
	return param.MemberUpdateParam{
		Id:      utils.ToUInt64(request.Id),
		Name:    request.Name,
		Phone:   request.Phone,
		LevelId: utils.ToUInt64(request.LevelId),
		Account: utils.ToFloat64(request.Account),
	}
}
