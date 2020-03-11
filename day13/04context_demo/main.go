package main

import (
"fmt"
"sync"

"time"
)

var wg sync.WaitGroup

// 使用channel 的方式实现

func worker(ch <-chan struct{}) {
	defer wg.Done()
	LABEL:
	for {
		select{
			case <-ch:
				break LABEL
			default:
				fmt.Println("worker")
				time.Sleep(time.Second)
			}
		}
	// 如何接收外部命令实现退出
}

func main() {
	var exitChan =make(chan struct{})
	wg.Add(1)
	go worker(exitChan)
	// 如何优雅的实现结束子goroutine
	time.Sleep(time.Second*5)
	exitChan<-struct{}{}
	wg.Wait()//等待....
	fmt.Println("over")
}
