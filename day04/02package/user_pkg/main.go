package main

import "fmt"

//导入包时起别名
import m "code.oldbody.com/studygolang/learnstuday/day04/02package/math_pkg"

//这个文件是干什么的
// import (
// 	"fmt"

// 	m "code.oldbody.com/studygolang/learnstuday/day04/02package/math_pkg" //起别名
// )

func main() {
	a := m.Add(1, 2)
	fmt.Println(a)
	stu := m.Student{Name: "aa", Age: 1}
	fmt.Println(stu.Name, stu.Age)
}
