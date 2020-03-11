package main

import "fmt"

//为甚么要用接口

//Cat结构体
type Cat struct{}

func (c Cat) Say() string { return "迷哦啊喵" }

//Dog 结构体
type Dog struct{}

func (d Dog) Say() string { return "汪汪汪" }

//定义一个接口
type Sayer interface {
	Say() string
}

func main() {
	// c := Cat{}
	// fmt.Println("猫：" + c.Say())
	// d := Dog{}
	// fmt.Println("狗：" + d.Say())

	var animalList []Sayer
	c := Cat{} //造一个猫
	d := Dog{} //造一个狗
	animalList = append(animalList, c, d)

	fmt.Println(animalList)

	for _, item := range animalList {
		ret := item.Say()
		fmt.Println(ret)
	}
}
