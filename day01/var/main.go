package main

import "fmt"

//全局变量声明必须用var（函数外部）
var test = "test"

//go语言函数里面可以有多个返回值
//定义有两个返回值的函数
func foo1() (int, string) {
	return 10, "Q1mi"
}

func main() {
	var aa int = 3
	fmt.Println(aa)
	//声明变量
	var (
		a int
		b bool
		c string
	)
	//赋值变量
	a = 1
	b = true
	c = "31231"
	fmt.Println(a, b, c)

	//变量声明并赋值
	var d string = "222"
	fmt.Println(d)

	//类型推导(编译器根据变量初始值的类型 指定给变量)
	var x = 200
	var y = "100"
	fmt.Println(x)
	fmt.Println("y=" + y)

	//简短变量声明 （只能在函数内部使用）
	nazha := "NAZHA" //函数内部通常会用这种声明方式 <=>  var nazha string="NAZHA"
	fmt.Println(nazha)
	//格式化输出
	var name string = "张三"
	name2 := "李四"
	fmt.Printf("欢迎%s登录", name)
	fmt.Printf("欢迎%s登录", name2)
	fmt.Println()

	//调用foo函数
	//_（匿名变量）用于接收不需要的变量值,并且可以重复声明
	bb, _ := foo1()
	fmt.Println(bb)
	_, cc := foo1()
	fmt.Println(cc)

}
