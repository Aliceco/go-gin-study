package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
type Topic struct {
	TId int
	TName string
}
func main() {
	//m:=make(map[string]interface{})
	//m["name"] = "an"
	//m["age"] = 12
	router:=gin.Default()
	router.GET("/", func(context *gin.Context) {
		//context.Writer.Write([]byte("hello gin"))
		//context.JSON(http.StatusOK, m)
		// gin.H{}
		//context.JSON(http.StatusOK, Topic{101, "title"})
		context.JSON(http.StatusOK, gin.H{"name": "an", "age": 12})
	})
	router.Run() // 默认8080
}
