package main

import "fmt"

//无缓冲通道和有缓冲通道

func recv(ch chan bool) {
	ret := <-ch //阻塞，一直等待ch接收值
	fmt.Println(ret)
}

func main() {
	//无缓冲通道
	// ch := make(chan bool) //无缓冲通道
	// //recv(ch)  //一直等待ch接收值，出现deadlock!
	// go recv(ch)
	// ch <- true
	// fmt.Println("main函数结束") //true  main函数结束

	//有缓冲通道
	//len:获取数据量 cap:获取容量
	ch := make(chan bool, 1) //1:容量
	ch <- false
	fmt.Println(len(ch), cap(ch))
	go recv(ch)             //函数里把值取走了
	ch <- true              //当前通道为空，所以可以放
	fmt.Println("main函数结束") //true  main函数结束

}
