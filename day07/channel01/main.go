package main

import "fmt"

//无缓冲区channel

func f(ch chan int) {
	<-ch //接收值
}

func main() {
	var ch chan int
	ch = make(chan int)

	go f(ch) //开启一个goroutine去执行f函数

	ch <- 100 //无缓冲区的通道，没有任何接收 100 就发送不进去
	fmt.Println("hello 沙河！")
}
