package main

import "fmt"

//函数进阶之变量作用域

//定义全局变量num
var num = 10

//定义函数
func testGlobal() {
	//可以在函数中访问全局变量
	fmt.Println("全局变量num", num)
	num := 100
	//可以在函数中访问变量--找自己函数内部变量
	fmt.Println("变量num", num)
}

func add(x, y int) int {
	return x + y
}
func sub(x, y int) int {
	return x - y
}
func calc(x, y int, op func(int, int) int) int {
	return op(x, y)
}

func main() {
	//函数可以作为变量
	abc := testGlobal
	fmt.Printf("%T\n", abc)
	abc()

	//函数可以作为参数
	a := calc(100, 200, add)
	fmt.Println(a)
	b := calc(100, 200, sub)
	fmt.Println(b)
}
