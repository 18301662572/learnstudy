package main

import "fmt"

//函数定义
//定义一个不需要参数也没有返回值的参数：sayHello
func sayHello() {
	fmt.Println("Hello World")
}

//定义一个带参数的函数
//定义一个接受string类型的name参数
func sayHello2(name string) {
	fmt.Println("HW", name)
}

//定义接受多个参数的函数
func intSum(a int, b int) int {
	ret := a + b
	return ret
}

//定义一个声明了返回值的函数
func intSum2(a int, b int) (ret int) {
	ret = a + b
	return
}

//函数接受
//函数接受可变参数，在参数名后面加...表示可变参数
//可变参数在函数中是切片类型
func intSum3(a ...int) int {
	// fmt.Println(a)
	// fmt.Printf("%T\n", a)
	ret := 0
	for _, v := range a {
		ret = ret + v
	}
	return ret
}

//固定参数和可变参数同时出现时，可变参数要放在最后
//go语言的函数中没有默认参数
func intSum4(a int, b ...int) int {
	ret := a
	for _, v := range b {
		ret = ret + v
	}
	return ret
}

//Go语言中函数类型简写
func intSum5(a, b int) (ret int) {
	return a + b
}

//定义具有多个返回值的函数,多返回值也支持类型简写
func calc(a, b int) (sum, sub int) {
	sum = a + b
	sub = a - b
	return
}

func main() {
	//函数调用
	sayHello()
	name := "SB"
	sayHello2(name)
	r := intSum(1, 2)
	fmt.Println(r)
	e1 := intSum3(10, 20) //a=[10,20]
	fmt.Println(e1)
	e2 := intSum4(0) //a=0 b=[]
	fmt.Println(e2)
	e3 := intSum4(12, 0, 20) //a=12 b=[0,20]
	fmt.Println(e3)
	intSum5(1, 2)
	x, y := calc(1, 2)
	fmt.Println(x, y)
}
