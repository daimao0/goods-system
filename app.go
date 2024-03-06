package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"go/types"
	"goods-system/internal/application/request"
	"goods-system/internal/application/vo"
	"goods-system/internal/infrastructure/common/api"
	"goods-system/internal/interface/controller"
)

var goodsController = controller.GetGoodsController()
var goodsTypeController = controller.GetGoodsTypeController()
var memberController = controller.GetMemberController()
var memberLevelController = controller.GetMemberLevelController()
var orderController = controller.GetOrderController()
var systemController = controller.GetSystemController()

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) GoodsGetById(id string) api.RespData[vo.GoodsVO] {
	return goodsController.GoodsGetById(id)
}

// GoodsAll 所有商品
func (a *App) GoodsAll() api.RespData[[]vo.GoodsVO] {
	return goodsController.GoodsAll()
}

// GoodsPage 商品分页
func (a *App) GoodsPage(req request.GoodsPageRequest) api.RespData[api.Page[vo.GoodsVO]] {
	return goodsController.GoodsPage(req)
}

// GoodsSave 保存商品
func (a *App) GoodsSave(req request.GoodsSaveRequest) api.RespData[types.Nil] {
	return goodsController.GoodsSave(req)
}

// GoodsDelete 删除商品
func (a *App) GoodsDelete(id string) api.RespData[types.Nil] {
	return goodsController.GoodsDelete(id)
}

// GoodsEdit 编辑商品
func (a *App) GoodsEdit(req request.GoodsEditRequest) api.RespData[types.Nil] {
	return goodsController.GoodsEdit(req)
}

// GoodsTypeSelectAll 查询所有商品类型
func (a *App) GoodsTypeSelectAll() api.RespData[[]vo.GoodsTypeVO] {
	return goodsTypeController.SelectAll()
}

// GoodsTypeSave 新增商品类型
func (a *App) GoodsTypeSave(request request.GoodsTypeSaveRequest) api.RespData[types.Nil] {
	return goodsTypeController.Save(request)
}

// GoodsTypeUpdate 更新商品类型
func (a *App) GoodsTypeUpdate(request request.GoodsTypeEditRequest) api.RespData[types.Nil] {
	return goodsTypeController.Update(request)
}

// GoodsTypeDeleteByName 删除商品类型
func (a *App) GoodsTypeDeleteByName(request request.GoodsTypeDeleteRequest) api.RespData[types.Nil] {
	return goodsTypeController.DeleteByName(request)
}

// MemberAll 获得所有会员
func (a *App) MemberAll() api.RespData[[]vo.MemberVO] {
	return memberController.All()
}

// MemberGetById 通过id查询会员
func (a *App) MemberGetById(id string) api.RespData[vo.MemberVO] {
	return memberController.GetById(id)
}

// MemberGetByPhone 通过手机号获得会员
func (a *App) MemberGetByPhone(phone string) api.RespData[vo.MemberVO] {
	return memberController.GetByPhone(phone)
}

// MemberPage 会员分页查询
func (a *App) MemberPage(request request.MemberPageRequest) api.RespData[api.Page[vo.MemberVO]] {
	return memberController.Page(request)
}

// MemberSave 保存会员
func (a *App) MemberSave(request request.MemberSaveRequest) api.RespData[types.Nil] {
	return memberController.Save(request)
}

// MemberUpdate 更新会员
func (a *App) MemberUpdate(request request.MemberUpdateRequest) api.RespData[types.Nil] {
	return memberController.Update(request)
}

// MemberDeleteById 更新会员
func (a *App) MemberDeleteById(id string) api.RespData[types.Nil] {
	return memberController.DeleteById(id)
}

// MemberLevelSelectAll 查询所有会员等级
func (a *App) MemberLevelSelectAll() api.RespData[[]vo.MemberLevelVO] {
	return memberLevelController.SelectAll()
}

// MemberLevelSave 查询所有会员等级
func (a *App) MemberLevelSave(request request.MemberLevelSaveRequest) api.RespData[types.Nil] {
	return memberLevelController.Save(request)
}

// MemberLevelUpdate 查询所有会员等级
func (a *App) MemberLevelUpdate(request request.MemberLevelUpdateRequest) api.RespData[types.Nil] {
	return memberLevelController.Update(request)
}

// MemberLevelDeleteById 会员等级根据id删除
func (a *App) MemberLevelDeleteById(id string) api.RespData[types.Nil] {
	return memberLevelController.DeleteById(id)
}

// Order 下订单
func (a *App) Order(request request.OrderRequest) api.RespData[types.Nil] {
	return orderController.Order(request)
}

// OrderRecords 订单记录
func (a *App) OrderRecords(request request.OrderRecordRequest) api.RespData[[]vo.OrderVO] {
	return orderController.Records(request)
}

// OrderPage 订单分页
func (a *App) OrderPage(request request.OrderPageRequest) api.RespData[api.Page[vo.OrderVO]] {
	return orderController.Page(request)
}

// OrderStatistics 订单统计
func (a *App) OrderStatistics(request request.OrderRecordRequest) api.RespData[vo.OrderStatisticsVO] {
	return orderController.Statistics(request)
}

// SystemDesktopPath 桌面path
func (a *App) SystemDesktopPath() api.RespData[vo.SystemDesktopPathVO] {
	return systemController.DesktopPath()
}

// SystemBackup 备份
func (a *App) SystemBackup(path string) api.RespData[types.Nil] {
	systemController.Backup(path)
	return api.Success(types.Nil{}, "")
}

// ImportBackup 备份
func (a *App) ImportBackup(base64File string) api.RespData[any] {
	// 解码Base64字符串
	fileBytes, err := base64.StdEncoding.DecodeString(base64File)
	var jsonData map[string]interface{}
	err = json.Unmarshal(fileBytes, &jsonData)
	if err != nil {
		return api.Fail[any]("文件解析失败")
	}
	systemController.ImportBackup(jsonData)
	return api.Success[any](types.Nil{}, "")
}
