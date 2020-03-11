package main

import (
	"fmt"
	"sync"
)

//启动goroutine

//利用sync.WaitGroup实现优雅的等待
var wg sync.WaitGroup //是一个结构体，它里面有一个计算器

func hello(i int) {
	defer wg.Done() //计算器-1，defer:不管有没有错误，计算器都-1
	fmt.Println("Hello 沙河！", i)
	if i == 8 {
		panic("报错了")
	}
}

func main() {
	defer fmt.Println("hahha") //在main()中跑的
	wg.Add(10)                 //计算器+10
	for i := 0; i < 10; i++ {
		//创建的每个线程都是并行的，没有顺序
		go hello(i) //1.创建一个goroutine 2.在新的goroutine 中执行hello()
	}
	fmt.Println("Hello main func")
	//time.Sleep(time.Second)
	//等hello执行完（执行hello函数的那个goroutine执行完）
	wg.Wait() //阻塞，一直等待所有的goroutine结束
	fmt.Println("main函数结束")
}
