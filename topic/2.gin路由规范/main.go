package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router:=gin.Default()

	router.GET("/v1/topic", func(c *gin.Context) {
		if c.Query("username")=="" {
			c.String(200, "获取topic列表")
		} else {
			c.String(200, "获取用户名%s的topic列表", c.Query("username"))
		}
	})
	router.GET("/v1/topic/:topic_id", func(c *gin.Context) {
		c.String(http.StatusOK, "获取topic_id为%s", c.Param("topic_id"))
	})

	router.Run() // 默认8080
}
