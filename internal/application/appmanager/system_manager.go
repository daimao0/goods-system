package appmanager

type SystemManager struct {
}

var systemManager *SystemManager

func GetSystemManager() *SystemManager {
	return systemManager
}

// Backup 备份
func (p *SystemManager) Backup() {

}
