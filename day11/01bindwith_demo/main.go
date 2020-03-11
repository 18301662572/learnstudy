package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

//shouldbindwith

type formA struct {
	Foo string `json:"foo" xml:"foo" binding:"required"`
}

type formB struct {
	Bar string `json:"bar" xml:"bar" binding:"required"`
}

//SomeHandler 自定义
func SomeHandler(c *gin.Context) {
	objA := formA{}
	objB := formB{}

	// c.ShouldBind 使用了 c.Request.Body，不可重用。
	// if errA := c.ShouldBind(&objA); errA == nil {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"msg": "传了一个foo数据",
	// 	})
	// 	//else if这一步， 因为现在 c.Request.Body 是 EOF，所以这里会报错。
	// } else if errB := c.ShouldBind(&objB); errB == nil {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"msg": "传了一个bar数据",
	// 	})
	// } else {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"msg": "不知道传的啥",
	// 	})
	// }

	if errA := c.ShouldBindBodyWith(&objA, binding.JSON); errA == nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "传了一个foo数据",
		})
		//else if这一步， 因为现在 c.Request.Body 是 EOF，所以这里会报错。
	} else if errB := c.ShouldBindBodyWith(&objB, binding.JSON); errB == nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "传了一个bar数据",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg": "不知道传的啥",
		})
	}
}

func main() {
	r := gin.Default()
	r.Any("/some", SomeHandler)
	r.Run()
}
