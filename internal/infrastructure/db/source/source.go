package source

import (
	"goods-system/internal/infrastructure/db/entity"
	"gorm.io/driver/sqlite" //驱动库
	"gorm.io/gorm"
	"log"
	"sync"
)

var once sync.Once
var db *gorm.DB

// Db 获得db
func Db() *gorm.DB {
	if db != nil {
		return db
	} else {
		once.Do(func() {
			var err error
			db, err = gorm.Open(sqlite.Open("goods-system.db"))
			if err != nil {
				log.Fatal("数据库连接失败", err.Error())
			}
			//创建表
			err = db.AutoMigrate(&entity.Goods{})
			err = db.AutoMigrate(&entity.GoodsType{})
			err = db.AutoMigrate(&entity.Member{})
			err = db.AutoMigrate(&entity.MemberLevel{})
			err = db.AutoMigrate(&entity.TopUp{})
			err = db.AutoMigrate(&entity.Order{})
			err = db.AutoMigrate(&entity.OrderItem{})

			if err != nil {
				log.Fatal("自动迁移时发生错误", err.Error())
			}
		})
	}
	return db
}
