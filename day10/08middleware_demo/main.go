package main

import (
	"net/http"
	"time"
	"fmt"
	"github.com/gin-gonic/gin"
)

//路由中间件

func castTime(c *gin.Context){
	start:=time.Now()
	c.Set("key","value")
	c.Next()  //运行下一个Handler函数	
	//统计耗时
	cast:=time.Since(start)
	fmt.Println("cast:",cast) 
}

func shopIndexHandler(c *gin.Context) {
	time.Sleep(3*time.Second)
	fmt.Println(c.MustGet("key").(string))
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "shopping/index",
	})
}

func shopHomeHandler(c *gin.Context) {
	time.Sleep(2*time.Second)
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "shopping/home",
	})
}

func main() {
	r := gin.Default() //Defult：默认使用了2个中间件 1.记日志 2.recover
	r.Use(castTime)
	//根据URL分执行的函数
	//http://127.0.0.1:8080/shopping/index
	shoppingGroup := r.Group("/shopping")
	{
		shoppingGroup.GET("/index",shopIndexHandler)
		shoppingGroup.GET("/home", shopHomeHandler)
	}
	r.Run()
}
