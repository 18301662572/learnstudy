package main

import "fmt"

//panic错误
func main() {
	defer func() { //定义一个延迟函数
		//recover --恢复,捕获panic异常
		err := recover() //尝试将函数从当前的异常状态恢复过来
		fmt.Println("recover抓到了panic异常：", err)
	}()
	var a []int
	a[0] = 100 //panic--惊恐 恐慌，异常
	fmt.Println("这是一个main函数")
}
