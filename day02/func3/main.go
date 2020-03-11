package main

import "fmt"

//匿名函数和闭包

//定义一个函数，返回值是一个函数
//把函数作为返回值
func a(name string) func() {
	return func() {
		fmt.Println("hello:", name)
	}
}

func main() {
	//执行匿名函数
	//方法1：将匿名函数赋值给一个变量，调用变量执行
	sayHello := func() {
		fmt.Println("匿名函数1")
	}
	sayHello()

	//方法2：立即执行， ()定义并执行匿名函数
	func() {
		fmt.Println("匿名函数2")
	}()

	//闭包=函数+外层变量的引用
	r := a("沙河娜扎") //r此时就是一个闭包
	r()            //相当于执行了a函数内部的匿名函数

}
