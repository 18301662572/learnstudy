package main

import "fmt"

//接口实现一个洗衣机
//只要一个类型它实现了wach() 和 dry()方法，我们就称这个类型实现了xi'yi'ji这个接口
type xiyiji interface {
	wash()
	dry()
}

type Haier struct {
	name  string
	price float64
	mode  string
}

func (h Haier) wash() {
	fmt.Println("Haier洗衣服")
}

func (h Haier) dry() {
	fmt.Println("Haier甩干")
}

type tianluo struct {
	name string
}

func (t tianluo) wash() {
	fmt.Println("田螺姑娘可以洗衣服")
}

func (t tianluo) dry() {
	fmt.Println("田螺姑娘可以拧干")
}

func main() {
	var a xiyiji //声明一个xiyiji类型的变量
	fmt.Println(a)
	h1 := Haier{ //实例化一个Haier结构体对象
		name:  "小神童",
		price: 999,
	}
	fmt.Printf("%T\n", h1)
	a = h1 //接口是一种类型，一种抽象的类型
	fmt.Println(a)

	t1 := tianluo{
		name: "SB",
	}
	a = t1
	fmt.Println(a)
}
