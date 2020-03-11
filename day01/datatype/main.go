package main

import (
	"fmt"
	"math"
)

func main() {

	//整型 int,uint
	var a int = 10   //十进制数
	var b int = 077  //八进制数
	var c int = 0xff //十六进制数
	fmt.Println(a, b)
	fmt.Printf("%b\n", a) //打印二进制
	fmt.Printf("%o\n", b) //打印八进制
	fmt.Printf("%x\n", c) //打印十六进制
	fmt.Println(b)        //将八进制转成十进制，打印出十进制数
	fmt.Println(c)        //将十六进制转成十进制，打印出十进制数
	//求c变量的内存地址--内存地址默认为十六进制
	fmt.Printf("%p\n", &c) //& 取地址

	//浮点型 float
	//浮点数常量
	fmt.Println(math.MaxFloat64)
	fmt.Println(math.MaxFloat32)

	//复数

	//布尔值 bool false/true,默认为false

	//字符串
	//字符串转义 \ 转义
	fmt.Println("\"c:\\go\"")
	var s1 = "单行文本"
	fmt.Println(s1)
	// ``中的字符不用转义
	var s2 = `这是
	多行\
	文本\
	`
	fmt.Println(s2)

}
