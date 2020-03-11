package main

import "fmt"

//channel

func main() {
	//定义一个ch1变量,声明
	//是一个channel类型,,声明
	//这个channel内部传递的数据是int类型
	var ch1 chan int
	var ch2 chan string

	//channel 是引用类型
	fmt.Println("ch1:", ch1) //nil
	fmt.Println("ch2:", ch2) //nil

	//make函数初始化(分配内存)： slice map channel
	ch3 := make(chan int, 10) //10:指定容量，否则会没有缓冲区，出现死锁

	//通道的操作：发送 、接受、关闭

	//发送和接受都用一个符号：  <-
	ch3 <- 10 //把10发送到ch3中
	//<-ch3        //从ch3中接收值，直接丢弃
	ret := <-ch3     //从ch3中接收值，保存到变量ret中
	fmt.Println(ret) //10

	//关闭
	close(ch3)
	//1.关闭的通道再接收，如果管道里面没有值了，能取到对应类型的零值
	ret2 := <-ch3
	fmt.Println(ret2) //0
	//2.往关闭的通道中发送值,会引发panic
	//ch3 <- 20 //panic: send on closed channel
	//3.关闭一个已经关闭的通道，会引发panic
	//close(ch3)//panic: close of closed channel
	//4.从一个关闭的管道里面取值，是可以取得。一次只能拿一个值。（跟队列一样）
	ch4 := make(chan int, 8)
	ch4 <- 9
	ch4 <- 8
	ch4 <- 7
	close(ch4)
	ret3 := <-ch4
	fmt.Println(ret3) //9

}
