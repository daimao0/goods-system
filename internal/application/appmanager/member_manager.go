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

var memberService = impl.GetMemberService()

type MemberManager struct {
}

var memberManager *MemberManager

func GetMemberManager() *MemberManager {
	return memberManager
}

// GetByPhone 根据手机号查询会员
func (p *MemberManager) GetByPhone(phone string) vo.MemberVO {
	memberDTO := memberService.Get(param.MemberSearchParam{Phone: phone})
	return convert.ToMemberVOByDTO(memberDTO)
}

// Page 分页查询会员信息
func (p *MemberManager) Page(request request.MemberPageRequest) api.Page[vo.MemberVO] {
	var memberVOS []vo.MemberVO
	pageParam := convert.ToMemberPageParamByRequest(request)
	page := memberService.Page(pageParam)
	for _, memberDTO := range page.Data {
		memberVO := convert.ToMemberVOByDTO(memberDTO)
		memberVOS = append(memberVOS, memberVO)
	}
	return api.Page[vo.MemberVO]{
		Page:  page.Page,
		Size:  page.Size,
		Total: page.Total,
		Data:  memberVOS,
	}
}

// Save 保存会员
func (p *MemberManager) Save(request request.MemberSaveRequest) {
	saveParam := convert.ToMemberSaveParamByRequest(request)
	memberService.Save(saveParam)
}

// Update 更新会员
func (p *MemberManager) Update(request request.MemberUpdateRequest) {
	updateParam := convert.ToMemberUpdateParamByRequest(request)
	memberService.Update(updateParam)
}

// DeleteById 删除会员
func (p *MemberManager) DeleteById(id string) {
	memberService.DeleteById(utils.ToUInt64(id))
}

// All 获得所有会员
func (p *MemberManager) All() []vo.MemberVO {
	var memberVOS []vo.MemberVO
	memberDTOS := memberService.All()
	for _, memberDTO := range memberDTOS {
		memberVO := convert.ToMemberVOByDTO(memberDTO)
		memberVOS = append(memberVOS, memberVO)
	}
	return memberVOS
}

// GetById 根据id获得会员
func (p *MemberManager) GetById(id string) vo.MemberVO {
	memberDTO := memberService.Get(param.MemberSearchParam{Id: utils.ToUInt64(id)})
	return convert.ToMemberVOByDTO(memberDTO)
}
