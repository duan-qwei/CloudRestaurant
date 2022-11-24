package datasource

import (
	"CloudRestaurant/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDataSource() (db *gorm.DB) {
	sourceConfig := config.Conf.DataSourceConfig
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		sourceConfig.User,
		sourceConfig.Password,
		sourceConfig.Host,
		sourceConfig.Port,
		sourceConfig.DB)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败")
	}

	return db
}
