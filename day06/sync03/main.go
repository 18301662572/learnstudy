package main

import (
	"fmt"
	"sync"
)

func f1(a int) {
	fmt.Println(a)
}

func closer(x int) func() {
	return func() {
		f1(x)
	}
}

var onlyonce sync.Once

func main() {
	f := closer(10) //利用闭包实现
	onlyonce.Do(f)
}
