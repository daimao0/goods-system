package controller

import (
	"go/types"
	"goods-system/internal/application/appmanager"
	"goods-system/internal/application/request"
	"goods-system/internal/application/vo"
	"goods-system/internal/infrastructure/common/api"
)

var memberManger = appmanager.GetMemberManager()

type MemberController struct {
}

var memberController *MemberController

func GetMemberController() *MemberController {
	return memberController
}

// All 所有会员
func (p *MemberController) All() api.RespData[[]vo.MemberVO] {
	memberVOS := memberManger.All()
	return api.Success(memberVOS, "")
}

// GetByPhone 通过手机号查询
func (p *MemberController) GetByPhone(phone string) api.RespData[vo.MemberVO] {
	memberVO := memberManger.GetByPhone(phone)
	return api.Success(memberVO, "")
}

// GetById 根据id查询会员
func (p *MemberController) GetById(id string) api.RespData[vo.MemberVO] {
	memberVO := memberManger.GetById(id)
	return api.Success(memberVO, "")
}

// Page 分页
func (p *MemberController) Page(request request.MemberPageRequest) api.RespData[api.Page[vo.MemberVO]] {
	page := memberManger.Page(request)
	return api.Success[api.Page[vo.MemberVO]](page, "")
}

// Save 保存
func (p *MemberController) Save(request request.MemberSaveRequest) api.RespData[types.Nil] {
	//判断会员手机号是否已注册
	memberDTO := memberManger.GetByPhone(request.Phone)
	if memberDTO.Id != "" {
		return api.Fail[types.Nil]("手机号已注册")
	}
	memberManger.Save(request)
	return api.Success[types.Nil](types.Nil{}, "")
}

// Update 更新
func (p *MemberController) Update(request request.MemberUpdateRequest) api.RespData[types.Nil] {
	memberManger.Update(request)
	return api.Success[types.Nil](types.Nil{}, "")
}

// DeleteById 删除
func (p *MemberController) DeleteById(id string) api.RespData[types.Nil] {
	memberManger.DeleteById(id)
	return api.Success[types.Nil](types.Nil{}, "")
}
