package main

import (
	"fmt"
	"math/rand"
)

//生产者消费者模型
//使用goroutine和channel实现一个简易的生产者消费者模型

//生产者：产生随机数 math/rand
//消费者：计算每个随机数的每个位的数字的和 1212123=？

//1个生产者 20个消费者

var itemChan chan *item
var resultChan chan *result
var doneChan chan struct{}

//Item 是一个类型
type item struct {
	id  int64
	num int64
}

type result struct {
	item *item
	sum  int64
}

//生产者
func producer(ch chan *item) {
	//1.生产随机数
	var id int64
	for i := 0; i < 100; i++ {
		id++
		number := rand.Int63() //int64正整数
		tmp := &item{
			id:  id,
			num: number,
		}
		//2.把随机数发送到通道中
		ch <- tmp
	}
	close(ch)
}

//计算一个数字每个位的和
func calc(num int64) int64 {
	// 123%10=12...3  sum=0+3
	// 12%10=1...2
	// 1%10=0...1
	var sum int64
	for num > 0 {
		sum = sum + num%10
		num = num / 10
	}
	return sum
}

//消费者
func consumer(ch chan *item, resultchan chan *result) {
	for tmp := range ch {
		//tmp := <-ch //结构体指针 *item
		sum := calc(tmp.num)
		//构造result结构体
		retobj := &result{
			item: tmp,
			sum:  sum,
		}
		resultchan <- retobj
	} //结构体指针 *item
	doneChan <- struct{}{} //发送一个空结构体给doneChan

}

func startWorker(n int, ch chan *item, resultchan chan *result) {
	for i := 0; i < n; i++ {
		go consumer(ch, resultchan)
	}
}

func closeResult(donech chan struct{}, resultChan chan *result) {
	for i := 0; i < 20; i++ {
		<-donech
	}
	close(donech)
	close(resultChan)
}

//打印结果
func printResult(resultchan chan *result) {
	for ret := range resultchan {
		fmt.Printf("id:%v，num:%v，sum:%v\n", ret.item.id, ret.item.num, ret.sum)
		//time.Sleep(time.Second)
	}
	//不知道什么时候应该关闭resultChan
	//当20个消费者的goroutine都执行完毕的时候 要关闭resultChan
}

func main() {
	itemChan = make(chan *item, 100)
	resultChan = make(chan *result, 100)
	doneChan = make(chan struct{}, 20)
	go producer(itemChan)
	startWorker(20, itemChan, resultChan)

	//启动一个检测关闭resultChan的goroutine
	go closeResult(doneChan, resultChan)

	//打印结果
	printResult(resultChan)
}
