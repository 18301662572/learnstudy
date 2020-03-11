package main

import "fmt"

func main() {
	var a int
	fmt.Println(a)               //0
	b := &a                      //取变量a的内存地址
	fmt.Println(b)               //0xc000056058 -取a的内存地址
	fmt.Printf("b=%v\n", b)      //b=0xc000056058
	fmt.Printf("b:type %T\n", b) // *int
	fmt.Printf("b= %v\n", *b)    //b= 0 -取内存地址对应的值
	c := "豪杰"
	//b=&c //取c的内存地址不能赋值给b，因为b已经被a赋值，b已经是一个int类型的指针
	//&取内存地址
	fmt.Printf("&c=%v\n", &c) //&c=0xc0000421c0
	d := 100
	b = &d
	fmt.Println(b) //0xc0000560a0
	//*取地址对应的值
	fmt.Println(*b) //100
	//指针可以做逻辑判断
	fmt.Println(b == &d) //true

	//指针的应用
	aa := [3]int{1, 2, 3}
	modifyArray(aa)   //在函数中复制了数组赋值给了内部的a1
	fmt.Println(aa)   //[1,2,3]
	modifyArray2(&aa) //传入数组的指针
	fmt.Println(aa)   //[100 2 3]

}

//定一一个修改的数组
func modifyArray(a1 [3]int) {
	a1[0] = 100 //只是修改了内部a1数组
}

//*[3]int 接受的参数是一个数组的指针
func modifyArray2(a1 *[3]int) {
	//(*a1)[0] = 100 //修改地址
	//语法糖：因为GO语言中指针不支持修改
	a1[0] = 100 //修改地址
}
