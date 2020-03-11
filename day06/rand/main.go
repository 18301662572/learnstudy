package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//给rand加随机数种子实现每一次执行都能产生真正的随机数
	rand.Seed(time.Now().UnixNano())
	ret := rand.Int63()    //int64正整数
	ret2 := rand.Intn(101) //[1,101)的随机数
	fmt.Println(ret)
	fmt.Println(ret2)
}
