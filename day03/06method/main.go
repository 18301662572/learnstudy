package main

import "fmt"

//函数
func dream1() {
	fmt.Println("我的梦想是不用上班还有钱")
}

//函数：谁都可以调用的
//方法:某个具体的类型才能调用的方法
type people struct {
	name string
	age  int
}

//函数指定接收者之后就是方法
//在GO语言中约定成俗不用 this 也不用self, 而是使用后面类型的首字母的小写
func (p *people) dream() {
	p.age = 15
	fmt.Println(p.name + "的梦想是不用上班还有钱")
	fmt.Printf("年龄：%d\n", p.age)
}

func main() {
	person := &people{
		name: "豪杰",
		age:  18,
	}
	person.dream()
	fmt.Println(person.age)
}
