package main

import (
	"github.com/gin-gonic/gin"
	"topic.jtthink.com/src"
)

func main() {
	// 三层架构
	router:=gin.Default()

	// 作用域
	topics:=router.Group("/v1/topics")
	{
		//topics.GET("", func(c *gin.Context) {
		//	if c.Query("username")=="" {
		//		c.String(200, "获取topic列表")
		//	} else {
		//		c.String(200, "获取用户名%s的topic列表", c.Query("username"))
		//	}
		//})
		topics.GET("", src.GetTopicList)

		//topics.GET("/:topic_id", src.GetTopicDetail)

		topics.Use(src.MustLogin())
		{

			topics.POST("/add", src.NewTopic)
			topics.POST("/dle/:topic_id", src.DelTopic)
		}
		//topics.GET("/:topic_id", src.GetTest()) // 想写()
	}

	router.Run() // 默认8080
}
