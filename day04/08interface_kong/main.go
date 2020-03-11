package main

import (
	"fmt"
	"time"
)

//空接口

func showType(a interface{}) {
	fmt.Printf("type:%T\n", a)
}

func main() {
	var s interface{}
	s = 100
	fmt.Println(s)
	s = "字符串"
	fmt.Println(s)
	s = true
	fmt.Println(s)
	s = struct{ name string }{name: "huahua"}
	fmt.Println(s)
	showType(s)
	showType(100)
	showType(time.Second)

	//定一个值为空接口的map
	var stuInfo = make(map[string]interface{}, 100)
	stuInfo["name"] = "100"
	stuInfo["age"] = 5
	stuInfo["豪杰"]=99.00
	fmt.Println(stuInfo)
}
