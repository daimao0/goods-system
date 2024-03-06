package controller

import (
	"fmt"
	"go/types"
	"goods-system/internal/application/vo"
	"goods-system/internal/infrastructure/common/api"
	"goods-system/internal/service/impl"
	"os/user"
	"path/filepath"
	"runtime"
)

type SystemController struct {
}

var systemController *SystemController

func GetSystemController() *SystemController {
	return systemController
}

var systemService = impl.GetSystemService()

// Backup 备份
func (p *SystemController) Backup(path string) api.RespData[types.Nil] {
	systemService.Backup(path)
	return api.Success(types.Nil{}, "")
}

// ImportBackup 导入备份文件
func (p *SystemController) ImportBackup(jsonData map[string]interface{}) api.RespData[types.Nil] {
	systemService.ImportBackup(jsonData)
	return api.Success(types.Nil{}, "")
}

// DesktopPath 桌面路径
func (p *SystemController) DesktopPath() api.RespData[vo.SystemDesktopPathVO] {
	// 获取当前用户信息
	u, err := user.Current()
	if err != nil {
		return api.Success(vo.SystemDesktopPathVO{Path: ""}, "")
	}

	// 根据操作系统类型构造桌面路径
	var desktopPath string
	switch runtime.GOOS {
	case "windows":
		// 在Windows下桌面路径通常是 `%USERPROFILE%\Desktop`
		desktopPath = filepath.Join(u.HomeDir, "Desktop")
	case "darwin":
		// 对于MacOS（Darwin），桌面位于 ~/Desktop
		desktopPath = filepath.Join(u.HomeDir, "Desktop")
	case "linux":
		// 大多数Linux发行版遵循 freedesktop.org 标准，桌面位于 ~/Desktop
		desktopPath = filepath.Join(u.HomeDir, "Desktop")
	default:
		desktopPath = ""
		err = fmt.Errorf("未知的操作系统类型: %s", runtime.GOOS)
	}
	return api.Success(vo.SystemDesktopPathVO{Path: desktopPath}, "")
}
