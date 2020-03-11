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

//编辑书籍
func editBookHandler(c *gin.Context) {
	//1.取到用户编辑的是谁？从querystring取到数的id值
	idstr := c.Query("id")
	if len(idstr) == 0 {
		//请求中没有携带id参数，请求无效
		c.String(http.StatusBadRequest, "请求无效")
	}
	//HTTP请求传过来的参数通常都是string类型，转换成对应的数据类型
	bookID, err := strconv.ParseInt(idstr, 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "请求无效")
		return
	}
	if c.Request.Method == "POST" {
		//1.获取用户提交的数据
		titleVal := c.PostForm("title")
		priceStr := c.PostForm("price")
		priceVal, err := strconv.ParseFloat(priceStr, 64)
		if err != nil {
			c.String(http.StatusBadRequest, "无效的价格信息")
		}
		//2.去数据库更新对应的书籍
		err = editBook(titleVal, priceVal, bookID)
		if err != nil {
			c.String(http.StatusInternalServerError, "更新数据失败")
		}
		//3.跳转回/book/list页面
		//相同网站跳转：写相对路径
		//不一样的网站跳转：写绝对路径
		c.Redirect(http.StatusMovedPermanently, "/book/list")

	} else {
		//需要给模板渲染上原来的旧数据
		//2.根据id取到书籍的信息
		bookBoj, err := queryBookByID(bookID)
		if err != nil {
			c.String(http.StatusBadRequest, "无效的书籍ID")
			return
		}
		//3.把书籍数据渲染到页面
		c.HTML(http.StatusOK, "book/edit_book.html", bookBoj)
	}
}

//获取书籍详情
func getBookDetailHandler(c *gin.Context) {
	//1.获取书籍ID
	tmpBookID := c.Param("id") //stirng类型的数据
	if len(tmpBookID) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"err": "invalid path param",
		})
		return
	}
	bookID, err := strconv.ParseInt(tmpBookID, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"err": "invalid path param",
		})
		return
	}
	//2.去数据库拿到具体的书籍信息
	bookObj, err := queryBookByID(bookID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"err": err.Error(),
		})
		return
	}
	//3.返回JSON数据
	c.JSON(http.StatusOK, bookObj)
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
	//编辑书籍
	r.Any("/book/edit", editBookHandler)

	//组合查询
	r.GET("/book/detail/:id", getBookDetailHandler)
	r.Run()
}
