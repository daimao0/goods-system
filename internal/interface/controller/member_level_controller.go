package controller

import (
	"go/types"
	"goods-system/internal/application/appmanager"
	"goods-system/internal/application/request"
	"goods-system/internal/application/vo"
	"goods-system/internal/infrastructure/common/api"
	"goods-system/internal/infrastructure/common/convert"
	"goods-system/internal/infrastructure/common/utils"
)

var memberLevelManger = appmanager.GetMemberLevelManager()

type MemberLevelController struct {
}

var memberLevelController *MemberLevelController

func GetMemberLevelController() *MemberLevelController {
	return memberLevelController
}

func (p *MemberLevelController) SelectAll() api.RespData[[]vo.MemberLevelVO] {
	var memberLevelVOS []vo.MemberLevelVO
	memberLevelDTOS := memberLevelManger.SelectAll()
	for _, memberLevelDTO := range memberLevelDTOS {
		memberLevelVO := convert.ToMemberLevelVO(memberLevelDTO)
		memberLevelVOS = append(memberLevelVOS, memberLevelVO)
	}
	return api.Success(memberLevelVOS, "")
}

// Save 保存会员等级
func (p *MemberLevelController) Save(request request.MemberLevelSaveRequest) api.RespData[types.Nil] {
	if request.Name == "" || utils.ToFloat64(request.Discount) == 0 {
		return api.Fail[types.Nil]("名称不存在或折扣不属于小数")
	}
	//查询是否已注册
	memberLevelDTO := memberLevelManger.GetByName(request.Name)
	if memberLevelDTO.Id != 0 {
		return api.Fail[types.Nil]("名称已存在")
	}
	memberLevelManger.Save(request)
	return api.Success(types.Nil{}, "")
}

// Update 会员更新
func (p *MemberLevelController) Update(request request.MemberLevelUpdateRequest) api.RespData[types.Nil] {
	memberLevelManger.Update(request)
	return api.Success(types.Nil{}, "")
}

// DeleteById 会员等级删除
func (p *MemberLevelController) DeleteById(id string) api.RespData[types.Nil] {
	memberLevelManger.DeleteById(id)
	return api.Success(types.Nil{}, "")
}
