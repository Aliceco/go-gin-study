package src

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)
var DBHelper *gorm.DB
var err error
func InitDB() {
	DBHelper,err=gorm.Open("mysql", "root:root@(localhost)/go_base?charset=utf8&parseTime=True&loc=Local")
	if err!=nil {
		fmt.Println("DB初始化失败：", err)
		ShutDownServer(err)
		return
	}
	DBHelper.LogMode(true) // 显示sql日志
	DBHelper.DB().SetMaxOpenConns(10)
	DBHelper.DB().SetMaxIdleConns(100)
	DBHelper.DB().SetConnMaxLifetime(time.Hour)
}
