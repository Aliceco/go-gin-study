package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"topic.jtthink.com/src"
)

func main() {
	db, _ := gorm.Open("mysql", "root:root@(localhost)/go_base?charset=utf8&parseTime=True&loc=Local")
	db.LogMode(true) // 显示sql日志
	//tc := &src.TopicClass{}
	//db.Table("topic_class").First(&tc) // 指定表名 查询单个
	//fmt.Println(tc) // &{1 技术类 12211 0}

	//var tcs [] src.TopicClass
	//db.Table("topic_class").Find(&tcs) // 指定表名 查询单个
	//fmt.Println(tcs) // [{1 技术类 12211 0} {2 产品类 asaasas 1}]

	var tcs [] src.TopicClass
	db.Where(&src.TopicClass{ClassName:"技术类"}).Find(&tcs) // TopicClass ,表名为topic_classes (注意复数，英文基础, ch sh x s 结尾时 加es变复数 )
	fmt.Println(tcs) // [{1 技术类 12211 0}]

	defer db.Close()
	//// 三层架构
	//router:=gin.Default()
	//// 作用域
	//v1:=router.Group("/v1/topics") // 单挑帖子
	//{
	//	v1.GET("", src.GetTopicList)
	//	v1.Use(src.MustLogin())
	//	{
	//
	//		v1.POST("/add", src.NewTopic)
	//		v1.POST("/dle/:topic_id", src.DelTopic)
	//	}
	//}
	//v2:=router.Group("/v2/topics") // 多条帖子
	//{
	//	v2.Use(src.MustLogin())
	//	{
	//		v2.POST("/add", src.NewTopics)
	//	}
	//}
	//router.Run() // 默认8080
}
