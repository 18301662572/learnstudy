package main

import "fmt"

//可以给任意类型追加方法
//不能给别的包定义的类型添加方法

//MyInt 是一个新类型
type MyInt int

func (m MyInt) sayHi() {
	fmt.Println("Hello")
}

func main() {
	var a MyInt
	fmt.Println(a)
	a.sayHi()
}
