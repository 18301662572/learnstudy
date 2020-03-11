package main

import "fmt"

//defer：延迟执行,在函数执行完快要结束的时候按照插入时的倒叙依次执行
//可以方便处理资源释放的问题
//func main() {
//	fmt.Println("start...")
//	defer fmt.Println(1)
//	defer fmt.Println(2)
//	defer fmt.Println(3)
//	fmt.Println("end...")
//}

//面试题
func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	x := 1
	y := 2
	defer calc("AA", x, calc("A", x, y))
	x = 10
	defer calc("BB", x, calc("B", x, y))
	y = 20

	//结果：
	//A 1 2 3
	//B 10 2 12
	//BB 10 12 22
	//AA 1 3 4
}