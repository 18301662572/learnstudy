package main

import (
	"flag"
	"fmt"
	"time"
)

//flag 设置命令行参数

func flagDemo1(){
	name:= flag.String("name","kelly","姓名")
	age:= flag.Int("age",23,"年龄")
	married:= flag.Bool("married",true,"婚否")
	flag.Parse()//解析命令行的参数
	fmt.Println(*name,*age,*married) //cmd:  flag_demo.exe -name=haojie -age=14
}

func flagDemo2(){
	var (
		name string
		age int
		married bool
		delay time.Duration
	)
	flag.StringVar(&name, "name", "张三", "姓名")
	flag.IntVar(&age, "age", 18, "年龄")
	flag.BoolVar(&married, "married", false, "婚否")
	flag.DurationVar(&delay, "d", 0, "时间间隔")
	flag.Parse()//解析命令行的参数
	fmt.Println(name,age,married,delay)

	fmt.Println(flag.Args())  //除flag参数之外的其他参数
	fmt.Println(flag.NArg())  //除flag参数之外的其他参数的数量
	fmt.Println(flag.NFlag())  //使用到的flag参数的个数
}

func main(){
	//flagDemo1()
	flagDemo2()
}

