package common

import (
	"CloudRestaurant/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDataSource() {
	sourceConfig := config.Conf.DataSourceConfig
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local&collation=%s&%s",
		sourceConfig.User,
		sourceConfig.Password,
		sourceConfig.Host,
		sourceConfig.Port,
		sourceConfig.DB,
		sourceConfig.Charset,
		sourceConfig.Collation,
		sourceConfig.Query)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 禁用外键(指定外键时不会在mysql创建真实的外键约束)
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic("连接数据库失败")
	}

	// 开启mysql日志
	if sourceConfig.LogMode {
		db.Debug()
	}

	DB = db
}
