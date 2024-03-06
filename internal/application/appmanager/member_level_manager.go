package appmanager

import (
	"goods-system/internal/application/request"
	"goods-system/internal/infrastructure/common/convert"
	"goods-system/internal/infrastructure/common/utils"
	"goods-system/internal/service/dto"
	"goods-system/internal/service/impl"
)

var memberLevelService = impl.GetMemberLevelService()

type MemberLevelManager struct {
}

var memberLevelManager *MemberLevelManager

func GetMemberLevelManager() *MemberLevelManager {
	return memberLevelManager
}

// SelectAll 查询所有会员等级
func (p *MemberLevelManager) SelectAll() []dto.MemberLevelDTO {
	return memberLevelService.SelectAll()
}

// Save 保存会员等级
func (p *MemberLevelManager) Save(request request.MemberLevelSaveRequest) {
	param := convert.ToMemberLevelSaveParamByRequest(request)
	memberLevelService.Save(param)
}

// Update 更新会员等级
func (p *MemberLevelManager) Update(request request.MemberLevelUpdateRequest) {
	param := convert.ToMemberLevelUpdateParamByRequest(request)
	memberLevelService.Update(param)
}

// DeleteById 删除会员等级
func (p *MemberLevelManager) DeleteById(id string) {
	memberLevelService.DeleteById(utils.ToUInt64(id))
}

// GetByName 根据名称查询会员等级
func (p *MemberLevelManager) GetByName(name string) dto.MemberLevelDTO {
	return memberLevelService.GetByName(name)
}
