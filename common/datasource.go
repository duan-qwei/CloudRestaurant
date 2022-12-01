package common

import (
	"CloudRestaurant/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
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

		//设置sql日志级别为Info
		Logger: logger.Default.LogMode(logger.Info),

		// 使用单数表名
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic("连接数据库失败")
	}

	DB = db
}
