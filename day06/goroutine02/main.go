package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"
)

//GOMAXPROCS  m:n 调度

var wg sync.WaitGroup

func a() {
	defer wg.Done()
	for i := 1; i < 10; i++ {
		fmt.Println("A：", i)
	}
}

func b() {
	// defer wg.Done()
	// for i := 1; i < 10; i++ {
	// 	fmt.Println("B:", i)
	// }
	var i int64
	for i < math.MaxInt64 {
		i++
	}
}

//为了避免cpu占满，电脑崩了，尽量不要运行这个程序
func main() {
	runtime.GOMAXPROCS(1) //设置我的Go程序只用一个逻辑核心 ，即m:n中设置n为1
	wg.Add(10)
	//go a()
	for i := 0; i < 10; i++ {
		go b()
	}
	wg.Wait()
}
