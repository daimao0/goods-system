package impl

import (
	"goods-system/internal/infrastructure/common/api"
	"goods-system/internal/infrastructure/common/convert"
	"goods-system/internal/infrastructure/db/entity"
	"goods-system/internal/infrastructure/db/repositories"
	"goods-system/internal/service/dto"
	"goods-system/internal/service/param"
	"time"
)

type MemberService struct {
}

var memberService *MemberService
var memberRepository = repositories.GetMemberRepository()

func GetMemberService() *MemberService {
	return memberService
}

// Page 分页
func (p *MemberService) Page(param param.MemberPageParam) api.Page[dto.MemberDTO] {
	var memberDTOS []dto.MemberDTO
	member := convert.ToMemberByPageParam(param)
	page := memberRepository.Page(&member, param.Page, param.Size)
	for _, entity := range page.Data {
		memberDTO := convert.ToMemberDTOByEntity(entity)
		p.setMemberLevelDTO(&memberDTO, entity.LevelId)
		memberDTOS = append(memberDTOS, memberDTO)
	}
	return api.Page[dto.MemberDTO]{
		Page:  page.Page,
		Size:  page.Size,
		Total: page.Total,
		Data:  memberDTOS,
	}
}

// Save 保存
func (p *MemberService) Save(param param.MemberSaveParam) dto.MemberDTO {
	member := convert.ToMemberBySaveParam(param)
	memberRepository.Insert(&member)
	return convert.ToMemberDTOByEntity(member)
}

// Update 更新
func (p *MemberService) Update(param param.MemberUpdateParam) {
	member := convert.ToMemberByUpdateParam(param)
	memberRepository.UpdateById(&member)
}

// DeleteById 根据id删除
func (p *MemberService) DeleteById(id uint64) {
	memberRepository.DeleteById(id)
}

// Get 获得单个会员
func (p *MemberService) Get(param param.MemberSearchParam) dto.MemberDTO {
	member := convert.ToMemberBySearchParam(param)
	memberRepository.Get(&member)
	memberDTO := convert.ToMemberDTOByEntity(member)
	p.setMemberLevelDTO(&memberDTO, member.LevelId)
	return memberDTO
}

// setMemberLevelDTO memberDTO赋值level
func (p *MemberService) setMemberLevelDTO(memberDTO *dto.MemberDTO, memberLevelId uint64) {
	levelDTO := memberLevelService.GetById(memberLevelId)
	memberDTO.Level = levelDTO
}

// All 获得所有会员
func (p *MemberService) All() []dto.MemberDTO {
	var memberDTOS []dto.MemberDTO
	members := memberRepository.List(entity.Member{})
	for _, member := range members {
		memberDTO := convert.ToMemberDTOByEntity(member)
		memberDTOS = append(memberDTOS, memberDTO)
	}
	return memberDTOS
}

// DeductMoney 扣存款
func (p *MemberService) DeductMoney(id uint64, pay float64) {
	memberDTO := memberService.Get(param.MemberSearchParam{Id: id})
	if memberDTO.Account-pay < 0 {
		memberDTO.Account = 0
	} else {
		memberDTO.Account = memberDTO.Account - pay
	}
	member := ToMember(memberDTO)
	memberRepository.UpdateById(&member)
}

func ToMember(memberDTO dto.MemberDTO) entity.Member {
	return entity.Member{
		Id:         memberDTO.Id,
		Name:       memberDTO.Name,
		LevelId:    memberDTO.Level.Id,
		Phone:      memberDTO.Phone,
		Account:    memberDTO.Account,
		UpdateTime: time.Now(),
	}
}
