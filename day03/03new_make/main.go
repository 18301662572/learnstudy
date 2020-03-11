package main

import "fmt"

func main() {
	//var a *int    //a 是一个int类型的指针 ，错误的写法
	//var b *string //b 是一个string类型的指针 ，错误的写法
	//var c *[3]int //c 是一个[3]int类型的指针 ，错误的写法

	//指针类型需要初始化才能使用,用new初始化 / make引用类型初始化中使用
	var a = new(int) //得到一个int类型的指针
	fmt.Println(a)
	*a = 10
	fmt.Println(a)
	fmt.Println(*a)

	var c = new([3]int) //初始化一个数组类型的指针
	fmt.Println(c)
	c[0] = 1 //(*c)[0]=1
	fmt.Println(*c)

}
