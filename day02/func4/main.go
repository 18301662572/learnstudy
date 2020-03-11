package main

import "fmt"

//defer执行时机 的经典案例
//在Go语言的函数中return语句在底层并不是原子操作，它分为给返回值赋值和RET指令两步。而defer语句执行的时机就在返回值赋值操作后，RET指令执行前。
func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x //5
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5 //6
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x //5
}
func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5 //5
}
func main() {
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
}
