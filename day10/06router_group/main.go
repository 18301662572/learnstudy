package main

import (
	"github.com/gin-gonic/gin"
)

//路由分组

func main() {
	r := gin.Default()
	//http://127.0.0.1:8080/shopping/index
	shoppingGroup := r.Group("/shopping")
	{
		shoppingGroup.GET("/index", shopIndexHandler)
		shoppingGroup.GET("/home", shopHomeHandler)
	}
	//http://127.0.0.1:8080/live/index
	liveGroup := r.Group("/live")
	{
		liveGroup.GET("/index", liveIndexHandler)
		liveGroup.GET("/home", liveHomeHandler)
	}
	//http://127.0.0.1:8080/v1/shopping/home
	v1 := r.Group("/v1")
	{
		v1Shopping := v1.Group("/shopping")
		{
			v1Shopping.GET("/home", shopHomeHandler)
		}
	}

	r.Run()
}
