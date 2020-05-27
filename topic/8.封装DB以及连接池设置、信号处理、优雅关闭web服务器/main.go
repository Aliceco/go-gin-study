package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"topic.jtthink.com/src"
)

func main2()  {
	count:=0
	go func() {
		for {
			fmt.Println("执行", count)
			count++
			time.Sleep(time.Second*1)
		}
	}()
	c:=make(chan os.Signal)
	go func() {
		ctx, _:=context.WithTimeout(context.Background(), time.Second*5)
		select {
		case <-ctx.Done():
			c<-os.Interrupt
		}
	}()
	signal.Notify(c)
	s:=<-c
	fmt.Println(s)
}
func main() {
	// 三层架构
	router:=gin.Default()
	// 作用域
	v1:=router.Group("/v1/topics") // 单挑帖子
	{
		v1.GET("", src.GetTopicList)
		v1.Use(src.MustLogin())
		{

			v1.POST("/add", src.NewTopic)
			v1.POST("/dle/:topic_id", src.DelTopic)
		}
	}
	v2:=router.Group("/v2/topics") // 多条帖子
	{
		v2.Use(src.MustLogin())
		{
			v2.POST("/add", src.NewTopics)
		}
	}
	//router.Run() // 默认8080
	server:=&http.Server{
		Addr:":8080",
		Handler:router,
	}
	go (func() { // 启动web服务器
		err:=server.ListenAndServe()
		if err!=nil{
			log.Fatal("服务器启动失败")
		}
	})()
	go (func() {
		src.InitDB()
	})()
	src.ServerNotify()
	// 在这里还可以做一些 释放连接或者善后工作，暂时略
	ctx, cancel:=context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	err:=server.Shutdown(ctx)
	if err!=nil{
		log.Fatalln("服务器关闭")
	}
	log.Fatalln("服务器优雅退出")
}
