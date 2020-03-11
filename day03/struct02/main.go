package main

import "fmt"

//结构体的内存布局
//结构体占用一块连续的内存
//内存是以字节为单位的十六进制
//1字节=8位=8bit

func main() {
	type test struct {
		a int8
		b int8
		c int8
	}
	var t = test{
		a: 1,
		b: 2,
		c: 3,
	}
	fmt.Println(&(t.a))
	fmt.Println(&(t.b))
	fmt.Println(&(t.c))

}
