package main

import "fmt"

//结构体是一个值类型
type student struct {
	name string
	age  int8
}

func main() {
	var stu01 = student{
		name: "haojie",
		age:  18,
	}
	stu02 := stu01 //将结构体stu01的值完整的复制一份给stu02
	stu02.name = "nahza"
	fmt.Println(stu01.name)
	fmt.Println(stu02.name)

	stu03 := &stu01                   //将stu01的地址赋值给了stu03
	fmt.Printf("stu03类型：%T\n", stu03) //*main.student
	stu03.name = "xiaowang"
	(*stu03).age = 15
	fmt.Println(stu01.name)
	fmt.Println(stu01.age)

}
