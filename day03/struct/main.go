package main

import "fmt"

//自定义类型（首字母大写：公有的，对外部可见，必须添加注释，注释格式如下
//首字母小写：对包里可见的）
//NewInt 是一个新类型
type NewInt int

//类型别名(类型的重命名):只存在代码编写过程中，代码编译之后根本不存在haojie
//提高代码可读性
type haoJie = int

//byte uint8
// var char uint8 //存储一个数字
// var char byte  //存储字符,byte类型的底层类型是uint8类型

func main() {
	var a NewInt
	fmt.Println(a)
	fmt.Printf("a：%T\n", a)
	var b haoJie
	fmt.Println(b)
	fmt.Printf("b:%T\n", b)

	//byte类型的底层类型是uint8类型
	var c byte = '1'
	fmt.Println(c)          //49
	fmt.Printf("c=%v\n", c) //49
	fmt.Printf("c=%c\n", c) //1
	fmt.Printf("c:%T\n", c) //uint8
	var d uint8 = 1
	fmt.Println(d)          //1
	fmt.Printf("d:%T\n", d) //uint8

}
