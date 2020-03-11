package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//BookManagementSystem

func bookListHandler(c *gin.Context) {
	//查数据
	bookList, err := queryAllBook()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  err,
		})
	}
	//返回给浏览器
	c.HTML(http.StatusOK, "book/list_book.tmpl", gin.H{
		"code": 0,
		"data": bookList,
	})
}

func addBookHandler(c *gin.Context) {
	//给用户返回一个添加书籍页面的处理函数
	c.HTML(http.StatusOK, "book/add_book.html", nil)
}

func createBookHandler(c *gin.Context) {
	//创建书籍的处理函数
	//从form表单取数据
	var msg string
	titleVal := c.PostForm("title")
	priceVal := c.PostForm("price")
	price, err := strconv.ParseFloat(priceVal, 64)
	if err != nil {
		msg = "无效的价格参数"
		c.JSON(http.StatusOK, gin.H{
			"msg": msg,
		})
		return
	}
	fmt.Printf("%T %T\n", titleVal, price)

	//拿到数据 插入数据库
	err = insertBook(titleVal, price)
	if err != nil {
		msg = "插入数据失败"
		c.JSON(http.StatusOK, gin.H{
			"msg": msg,
		})
		return
	}
	//插入数据成功
	c.Redirect(http.StatusMovedPermanently, "/book/list")
}

//删除具体某本书
func deleteBookHandler(c *gin.Context) {
	//取query string 参数
	idStr := c.Query("id")
	idVal, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  err,
		})
		return
	}
	//数据是个正经数字
	//去数据库删除具体的记录
	err = deleteBook(idVal)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  err,
		})
		return
	}
	//删除成功还是跳转到书籍列表页
	c.Redirect(http.StatusMovedPermanently, "/book/list")

}

func main() {
	//程序启动就应该链接数据库
	err := initDB()
	if err != nil {
		panic(err)
	}
	r := gin.Default()
	r.LoadHTMLGlob("template/**/*")
	//查看所有的书籍数据
	r.GET("/book/list", bookListHandler)
	//添加书籍
	r.GET("/book/add", addBookHandler)
	r.POST("/book/add", createBookHandler)
	r.GET("/book/delete", deleteBookHandler)
	r.Run()
}
