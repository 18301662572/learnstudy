package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//gin demo 和gin参数绑定

//UserInfo 一个结构体
type UserInfo struct {
	UserName string `form:"username" json:"username" binding:"required"`
	PassWord string `form:"password" json:"password"`
}

func indexHandler(c *gin.Context) {
	//具体的处理请求的业务逻辑
	c.JSON(http.StatusOK, gin.H{
		"msg": "Hello",
	})
}

/*
func loginHandler(c *gin.Context) {
	if c.Request.Method == "POST" {
		//form 处理用户提交过来的请求数据
		username := c.PostForm("username")
		pwd := c.PostForm("password")
		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"pwd":      pwd,
		})
	} else {
		//返回一个登陆的页面
		c.HTML(http.StatusOK, "login.html", nil)
	}
}
*/

func searchHandler(c *gin.Context) {
	//取querystring
	name := c.Query("name")
	age := c.Query("age")
	c.JSON(http.StatusOK, gin.H{
		"name": name,
		"age":  age,
	})
}

func postsHandler(c *gin.Context) {
	//取道path参数
	year := c.Param("year")
	month := c.Param("month")
	day := c.Param("day")
	c.JSON(http.StatusOK, gin.H{
		"year":  year,
		"month": month,
		"day":   day,
	})
}

func loginHandler(c *gin.Context) {
	if c.Request.Method == "POST" {
		//处理用户提交过来的请求数据
		var u UserInfo //UserInfo 类型变量
		//ShouldBind 会根据请求头中的Content-Type
		err := c.ShouldBind(&u)
		if err != nil {
			//解析出问题
			c.JSON(http.StatusOK, gin.H{
				"err": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"username": u.UserName,
			"pwd":      u.PassWord,
		})
	} else {
		//返回一个登陆的页面
		c.HTML(http.StatusOK, "login.html", nil)
	}

}

func main() {
	r := gin.Default() //得到一个默认的处理引擎
	//加载HTML文件
	r.LoadHTMLGlob("templates/*")
	// r.LoadHTMLFiles("xx.html","oo.html",...)
	//加载静态文件
	r.Static("/dsb", "./statics")

	r.GET("/index", indexHandler)
	//所有请求的URL以v1开头的都交给下面的v1Group
	v1Group := r.Group("/v1")
	{
		v1Group.GET("/index", indexHandler)
	}

	//登录
	// r.GET("/login", loginHandler)
	// r.POST("/login", loginHandler)
	r.Any("/login", loginHandler)

	//query string参数
	r.GET("/search", searchHandler)

	//path参数
	r.GET("/posts/:year/:month/:day", postsHandler)

	r.Run(":8090")
}
