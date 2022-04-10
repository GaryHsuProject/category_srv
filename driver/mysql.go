package driver

import (
	"fmt"
	"shop/models"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var once sync.Once
var DB *gorm.DB

func InitGorm() *gorm.DB {

	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		GlobalConfig.MysqlConfig.Username,
		GlobalConfig.MysqlConfig.Password,
		GlobalConfig.MysqlConfig.Host,
		GlobalConfig.MysqlConfig.Port,
		GlobalConfig.MysqlConfig.Name,
	)
	once.Do(func() {
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}
	})
	err = DB.AutoMigrate(&models.Category{})
	if err != nil {
		panic(err)
	}
	return DB
}
