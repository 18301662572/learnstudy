package main

import (
	"fmt"
	"time"
)

//内置的time包

func timestamp2Timeboj(timestamp int64) {
	timeobj := time.Unix(timestamp, 0) //将时间戳转化为时间格式 unix参数（秒，纳秒）纳秒：[0,999999999]
	fmt.Println(timeobj)
	year := timeobj.Year()
	month := timeobj.Month()
	day := timeobj.Day()
	hour := timeobj.Hour()
	minute := timeobj.Minute()
	second := timeobj.Second()
	fmt.Printf("%4d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}

//定时器
func tickDemo() {
	ticker := time.Tick(time.Second) //定义一个1秒间隔的定时器

	for i := range ticker {
		fmt.Println(i) //每秒都会执行的任务
	}
}

//go 语言格式化
//2006表示年 01表示月 02 表示日 15表示时 04表示分 05表示秒 000表示毫秒
//是go语言诞生的日期
//2006-01-02 15:04:05.000   200612345
func formatDemo() {
	now := time.Now()
	fmt.Println(now.Format("2006-01-02 15:04:05.000"))
	fmt.Println(now.Format("2006/01/02 15:04:05.000"))
	fmt.Println(now.Format("15:04 2006/01/02 "))
	fmt.Println(now.Format("2006/01/02"))
}

func main() {
	//time.Time 结构体
	now := time.Now()
	fmt.Printf("%#v\n", now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	fmt.Println(now.Nanosecond())

	//时间戳
	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano())
	tm := now.Unix() + 3600
	timestamp2Timeboj(tm)

	//定时器
	//tickDemo()

	//日期格式化
	formatDemo()

}
