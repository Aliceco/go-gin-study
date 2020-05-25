package main

import (
	"github.com/gin-gonic/gin"
	"topic.jtthink.com/src"
)

func main() {
	// 三层架构
	router:=gin.Default()
	// 注册验证器
	//if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
	//	v.RegisterValidation("topicUrl", src.TopicUrl)
	//}
	// 作用域
	v1:=router.Group("/v1/topics") // 单挑帖子
	{
		//v1.GET("", func(c *gin.Context) {
		//	if c.Query("username")=="" {
		//		c.String(200, "获取topic列表")
		//	} else {
		//		c.String(200, "获取用户名%s的topic列表", c.Query("username"))
		//	}
		//})
		v1.GET("", src.GetTopicList)

		//v1.GET("/:topic_id", src.GetTopicDetail)

		v1.Use(src.MustLogin())
		{

			v1.POST("/add", src.NewTopic)
			v1.POST("/dle/:topic_id", src.DelTopic)
		}
		//v1.GET("/:topic_id", src.GetTest()) // 想写()
	}
	v2:=router.Group("/v2/topics") // 多条帖子
	{
		v2.Use(src.MustLogin())
		{
			v2.POST("/add", src.NewTopics)
		}
	}

	router.Run() // 默认8080
}
