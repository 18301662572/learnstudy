package main

import "fmt"

//结构体内嵌 模拟继承

type animal struct {
	name string
}

//定义一个动物会动的方法
func (a *animal) move() {
	fmt.Printf("%s会动", a.name)
}

type dog struct {
	feet int
	*animal
}

//定义一个狗的方法
func (d dog) wangwang() {
	fmt.Printf("%s 在叫\n", d.name)
}

func main() {
	var a = dog{
		feet: 4,
		animal: &animal{
			name: "旺财",
		},
	}
	a.wangwang() //调用狗的方法
	a.move()     //调用动物的方法
}
