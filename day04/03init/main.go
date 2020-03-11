package main

import "fmt"

import _ "code.oldbody.com/studygolang/learnstuday/day04/02package/math_pkg"

//init()  示例

var toyad = "周天"

//Week 星期常量
const Week = 7

type stu struct {
	Name string
	age  int
}

//包被导入的时候会自动执行(多用来做初始化操作)
func init() {
	fmt.Println("包被初始化了！")
	stu01 := stu{Name: "haojie"}
	fmt.Printf("toyad:%s,week:%d,name:%s\n", toyad, Week, stu01.Name)
}

func main() {
	fmt.Println("这是一个main函数")
}
