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

func GetTopicDetail(c *gin.Context) {
	c.String(200, "获取topic_id为%s", c.Param("topic_id"))
}
func NewTopic(c *gin.Context)  {
	c.String(200, "新增topic")
}
func DelTopic(c *gin.Context)  {
	c.String(200, "删除topic")
}

func GetTopicList() {

}

