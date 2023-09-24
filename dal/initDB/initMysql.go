package initDB

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var dsn = "root:root@tcp(127.0.0.1:3306)/chat?charset=utf8mb4&parseTime=True&loc=Local"
var MYSQL_DB *gorm.DB

func InitMysql() {
	var err error
	MYSQL_DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		Logger:                 logger.Default.LogMode(logger.Info),
		TranslateError:         true,
	})
	if err != nil {
		panic("初始化数据库连接失败！！！")
	}
}
