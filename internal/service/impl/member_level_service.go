package impl

import (
	"goods-system/internal/infrastructure/common/convert"
	"goods-system/internal/infrastructure/db/entity"
	"goods-system/internal/infrastructure/db/repositories"
	"goods-system/internal/service/dto"
	"goods-system/internal/service/param"
)

type MemberLevelService struct {
}

var memberLevelService *MemberLevelService
var memberLevelRepository = repositories.GetMemberLevelRepository()

func GetMemberLevelService() *MemberLevelService {
	return memberLevelService
}

// GetById 根据ID查询会员等级信息
func (p *MemberLevelService) GetById(id uint64) dto.MemberLevelDTO {
	memberLevel := entity.MemberLevel{Id: id}
	memberLevelRepository.Get(&memberLevel)
	return convert.ToMemberLevelDTO(memberLevel)
}

// SelectAll 查询所有
func (p *MemberLevelService) SelectAll() []dto.MemberLevelDTO {
	var memberLevelDTOS []dto.MemberLevelDTO
	memberLevels := memberLevelRepository.SelectAll()
	for _, memberLevel := range memberLevels {
		memberLevelDTO := convert.ToMemberLevelDTO(memberLevel)
		memberLevelDTOS = append(memberLevelDTOS, memberLevelDTO)
	}
	return memberLevelDTOS
}

// Save 保存
func (p *MemberLevelService) Save(param param.MemberLevelSaveParam) dto.MemberLevelDTO {
	memberLevel := convert.ToMemberLevelBySaveParam(param)
	memberLevelRepository.Insert(&memberLevel)
	return convert.ToMemberLevelDTO(memberLevel)
}

// Update 更新
func (p *MemberLevelService) Update(param param.MemberLevelUpdateParam) {
	memberLevel := convert.ToMemberLevelByUpdateParam(param)
	memberLevelRepository.UpdateById(&memberLevel)
}

// DeleteById 根据id删除
func (p *MemberLevelService) DeleteById(id uint64) {
	memberLevelRepository.DeleteById(id)
}

// GetByName 根据名称获得会员等级
func (p *MemberLevelService) GetByName(name string) dto.MemberLevelDTO {
	memberLevel := entity.MemberLevel{Name: name}
	memberLevelRepository.Get(&memberLevel)
	return convert.ToMemberLevelDTO(memberLevel)
}
