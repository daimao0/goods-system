package impl

import (
	"encoding/json"
	"goods-system/internal/infrastructure/common/utils"
	"goods-system/internal/infrastructure/db/entity"
	"goods-system/internal/infrastructure/db/query"
)

type SystemService struct {
}

var systemService *SystemService

func GetSystemService() *SystemService {
	return systemService
}

// Backup 备份
func (p *SystemService) Backup(path string) {
	goods := goodsRepository.Select(entity.Goods{})
	goodsTypes := goodsTypeRepository.SelectAll()
	members := memberRepository.List(entity.Member{})
	memberLevels := memberLevelRepository.SelectAll()
	orders := orderRepository.List(query.OrderQuery{})
	orderItems := orderItemRepository.SelectAll()
	result := make(map[string]interface{})
	result["goods"] = goods
	result["goodsTypes"] = goodsTypes
	result["members"] = members
	result["memberLevels"] = memberLevels
	result["orders"] = orders
	result["orderItems"] = orderItems
	data, _ := json.Marshal(result)
	dataStr := string(data)
	path = path + "/goods_backup.json"
	utils.WriteFile(dataStr, path)
}

func (p *SystemService) ImportBackup(fileData map[string]interface{}) {
	var goods []entity.Goods
	var goodsType []entity.GoodsType
	var member []entity.Member
	var memberLevel []entity.MemberLevel
	var order []entity.Order
	var orderItem []entity.OrderItem
	goodsMarshal, _ := json.Marshal(fileData["goods"])
	goodsTypeMarshal, _ := json.Marshal(fileData["goodsTypes"])
	memberMarshal, _ := json.Marshal(fileData["members"])
	memberLevelMarshal, _ := json.Marshal(fileData["memberLevels"])
	orderMarshal, _ := json.Marshal(fileData["orders"])
	orderItemMarshal, _ := json.Marshal(fileData["orderItems"])
	_ = json.Unmarshal(goodsMarshal, &goods)
	_ = json.Unmarshal(goodsTypeMarshal, &goodsType)
	_ = json.Unmarshal(memberMarshal, &member)
	_ = json.Unmarshal(memberLevelMarshal, &memberLevel)
	_ = json.Unmarshal(orderMarshal, &order)
	_ = json.Unmarshal(orderItemMarshal, &orderItem)

	goodsRepository.BatchSave(goods)
	goodsTypeRepository.BatchSave(goodsType)
	memberRepository.BatchSave(member)
	memberLevelRepository.BatchSave(memberLevel)
	orderRepository.BatchSave(order)
	orderItemRepository.BatchSave(orderItem)
}
