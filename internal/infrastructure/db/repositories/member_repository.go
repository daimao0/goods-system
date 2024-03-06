package repositories

import (
	"goods-system/internal/infrastructure/common/api"
	"goods-system/internal/infrastructure/common/global"
	"goods-system/internal/infrastructure/db/entity"
	"time"
)

var memberRepository *MemberRepository

type MemberRepository struct {
}

func GetMemberRepository() *MemberRepository {
	return memberRepository
}

// Page 分页查询会员
func (p *MemberRepository) Page(member *entity.Member, page, size int) api.Page[entity.Member] {
	var total int64
	var members []entity.Member
	tx := global.DB.Model(&entity.Member{})
	if member.Name != "" {
		tx.Where("name like ?", "%"+member.Name+"%")
	}
	if member.Phone != "" {
		tx.Where("phone like ?", member.Phone+"%")
	}
	if member.LevelId != 0 {
		tx.Where("level_id = ?", member.LevelId)
	}
	if member.Account != 0 {
		tx.Where("account > ?", member.Account)
	}
	tx.Count(&total).
		Limit(size).
		Offset((page - 1) * size).
		Find(&members)
	return api.Page[entity.Member]{
		Page:  page,
		Size:  size,
		Total: total,
		Data:  members,
	}
}

// Insert 保存会员
func (p *MemberRepository) Insert(member *entity.Member) {
	global.DB.Create(member)
}

// BatchSave 批量插入
func (p *MemberRepository) BatchSave(members []entity.Member) {
	for _, member := range members {
		global.DB.Save(&member)
	}
}

// SelectById 根据id查询会员
func (p *MemberRepository) SelectById(id uint64) entity.Member {
	member := entity.Member{}
	global.DB.First(&member, id)
	return member
}

// SelectOneByPhone 根据手机号查询会员
func (p *MemberRepository) SelectOneByPhone(phone string) entity.Member {
	member := entity.Member{}
	tx := global.DB.Model(&member)
	if member.Phone != "" {
		tx.Where("phone = ?", phone)
	}
	tx.First(&member)
	return member
}

// UpdateById 更新
func (p *MemberRepository) UpdateById(member *entity.Member) {
	flag := false
	//查询原来的数据
	old := p.SelectById(member.Id)
	// 创建一个只包含有变化字段的 map
	updates := make(map[string]interface{})
	if member.Name != old.Name {
		updates["name"] = member.Name
		flag = true
	}
	if member.LevelId != old.LevelId {
		updates["level_id"] = member.LevelId
		flag = true
	}
	if member.Phone != old.Phone {
		updates["phone"] = member.Phone
		flag = true
	}
	if member.Account != old.Account {
		updates["account"] = member.Account
		flag = true
	}
	if flag {
		updates["update_time"] = time.Now()
		global.DB.Model(member).Updates(updates)
	}
}

// DeleteById 根据id删除会员
func (p *MemberRepository) DeleteById(id uint64) {
	var member entity.Member
	global.DB.Delete(&member, id)
}

// Get 查询单个会员
func (p *MemberRepository) Get(member *entity.Member) {
	tx := global.DB.Model(member)
	if member.Id != 0 {
		tx.First(member, member.Id)
		return
	}
	if member.Phone != "" {
		tx.Where("phone = ?", member.Phone).First(member)
		return
	}
}

// List 获得所有会员
func (p *MemberRepository) List(member entity.Member) []entity.Member {
	var members []entity.Member
	tx := global.DB.Model(&member)
	if member.Name != "" {
		tx.Where("name = ?", member.Name)
	}
	if member.LevelId != 0 {
		tx.Where("level_id = ?", member.LevelId)
	}
	if member.Account != 0 {
		tx.Where("account = ?", member.Account)
	}
	tx.Find(&members)
	return members
}
