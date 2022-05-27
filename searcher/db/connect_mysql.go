package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var MysqlDB *gorm.DB
var err error

func ConnectMySql() {
	dsn := "root:root@tcp(101.200.128.148:3306)/documents?charset=utf8mb4&parseTime=True&loc=Local"
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,        // 彩色打印
		},
	)
	//全局模式
	MysqlDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("数据库连接失败:" + err.Error())
	}
	fmt.Println("数据库连接成功")

	//自动迁移 建表
	//err = DB.AutoMigrate(&Docs{})
	//if err != nil {
	//	fmt.Println("自动迁移失败....", err.Error())
	//}
}
