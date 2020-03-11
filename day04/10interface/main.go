package main

import "fmt"

//实现接口时使用指针接收者和使用值接受者的区别

type animal interface {
	speak()
	move()
}

//Cat 结构体
type Cat struct {
	name string
}

//cat实现animal的接口

//方法一：使用值接收者
// func (c Cat) speak() {
// 	fmt.Println("米哦啊秒！！")
// }
// func (c Cat) move() {
// 	fmt.Println("米哦啊会动！！")
// }

//方法二：使用者指针接收者
func (c *Cat) speak() {
	fmt.Println("米哦啊秒！！")
}
func (c *Cat) move() {
	fmt.Println("米哦啊会动！！")
}

func main() {
	var x animal

	//方法一可以使用，
	//方法二报错，实现animal接口的是*Cat类型，hh此时是一个Cat类型
	//hh := Cat{"huahua"}
	//x = hh

	tom := &Cat{"汤姆猫"} //*Cat 类型
	x = tom
	tom.speak() //(*tom).speak()
	tom.move()  //(*tom).move()
	fmt.Println(x)

}
