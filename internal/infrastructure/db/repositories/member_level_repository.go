package repositories

import (
	"goods-system/internal/infrastructure/common/global"
	"goods-system/internal/infrastructure/db/entity"
	"time"
)

type MemberLevelRepository struct{}

var memberLevelRepository *MemberLevelRepository

func GetMemberLevelRepository() *MemberLevelRepository {
	return memberLevelRepository
}

// Insert 保存会员等级
func (p *MemberLevelRepository) Insert(memberLevel *entity.MemberLevel) {
	global.DB.Create(memberLevel)
}

// SelectById 根据id查询会员等级
func (p *MemberLevelRepository) SelectById(id uint64) entity.MemberLevel {
	memberLevel := entity.MemberLevel{}
	global.DB.First(&memberLevel, id)
	return memberLevel
}

// SelectAll 查询全部会员等级
func (p *MemberLevelRepository) SelectAll() []entity.MemberLevel {
	var memberLevels []entity.MemberLevel
	global.DB.Find(&memberLevels)
	return memberLevels
}

// UpdateById 更新会员等级
func (p *MemberLevelRepository) UpdateById(memberLevel *entity.MemberLevel) {
	flag := false
	old := p.SelectById(memberLevel.Id)
	// 创建一个只包含有变化字段的 map
	updates := make(map[string]interface{})
	if old.Name != memberLevel.Name {
		updates["name"] = memberLevel.Name
		flag = true
	}
	if old.Discount != memberLevel.Discount {
		updates["discount"] = memberLevel.Discount
		flag = true
	}
	if flag {
		updates["update_time"] = time.Now()
		global.DB.Model(memberLevel).Updates(updates)
	}
}

// DeleteById 删除会员等级
func (p *MemberLevelRepository) DeleteById(id uint64) {
	var memberLevel entity.MemberLevel
	global.DB.Delete(&memberLevel, id)
}

// Get 查询单个
func (p *MemberLevelRepository) Get(memberLevel *entity.MemberLevel) {
	tx := global.DB.Model(memberLevel)
	if memberLevel.Id != 0 {
		tx.First(memberLevel)
		return
	}
	if memberLevel.Name != "" {
		tx.Where("name = ?", memberLevel.Name).First(memberLevel)
		return
	}
}

// BatchSave 批量保存
func (p *MemberLevelRepository) BatchSave(level []entity.MemberLevel) {
	for _, item := range level {
		global.DB.Save(&item)
	}
}
