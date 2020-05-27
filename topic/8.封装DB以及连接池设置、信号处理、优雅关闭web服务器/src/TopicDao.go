package src

import (
	"github.com/gin-gonic/gin"
)

//func GetTest() gin.HandlerFunc  {
//	return func (c *gin.Context) {
//		c.String(200, "获取topic_id为%s", c.Param("topic_id"))
//	}
//}

func MustLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, status:=c.GetQuery("token"); !status {
			c.String(401, "缺少token参数")
		}
	}
}

//func GetTopicDetail(c *gin.Context) {
//	c.JSON(200, CreateTopic(1, "标题"))
//	//c.String(200, "获取topic_id为%s", c.Param("topic_id"))
//}

// 单
func NewTopic(c *gin.Context)  {
	topic:=Topic{}
	err:=c.BindJSON(&topic)
	if err!=nil {
		c.String(400, "参数错误：%s", err.Error())
	} else {
		c.JSON(200, topic)
	}
}
// 多
func NewTopics(c *gin.Context)  {
	topic:=Topics{}
	err:=c.BindJSON(&topic)
	if err!=nil {
		c.String(400, "参数错误：%s", err.Error())
	} else {
		c.JSON(200, topic)
	}
}

func DelTopic(c *gin.Context)  {
	c.String(200, "删除topic")
}

func GetTopicList(c *gin.Context) {
	query:=TopicQuery{}
	err:=c.BindQuery(&query)
	if err!=nil {
		c.String(400, "参数错误：%s", err.Error())
	} else {
		c.JSON(200, query)
	}
}

