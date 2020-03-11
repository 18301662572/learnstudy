package main

import "fmt"

func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}

//变量f1,f2是个函数并且它引用了其外部作用域中的base变量，此时f1,f2就是一个闭包。 在f1,f2的生命周期内，变量base也一直有效。
func main() {
	f1, f2 := calc(10)        // base:10
	fmt.Println(f1(1), f2(2)) //11 9   base:9
	fmt.Println(f1(3), f2(4)) //12 8   base:8
	fmt.Println(f1(5), f2(6)) //13 7   base:7
}
