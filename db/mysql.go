package db

import (
	. "css/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

var DB *gorm.DB

func InitMySQL() {
	fmt.Println("数据库初始化.....")
	dsn := "root:zhijing@tcp(172.16.1.24:3306)/css_test?charset=utf8mb4&parseTime=True&loc=Local"
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			LogLevel: logger.Info,
		},
	)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		fmt.Println("mysql error:", err)
	}
	err = DB.AutoMigrate(
		&Class{},
		&Course{},
		&Teacher{},
		&Student{},
		&User{},
		&Role{},
		&Permission{},
		&Admin{},
	)
	if err != nil {
		return
	}
}
