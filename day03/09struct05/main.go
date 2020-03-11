package main

import "fmt"

//结构体的嵌套
type address struct {
	provice string
	city    string
}

type student struct {
	name string
	age  int
	ads  address //嵌套了别的结构体
}

type student1 struct {
	name    string
	age     int
	address //嵌套了匿名结构体
}

func main() {
	var stu1 = student{
		"kelly",
		18,
		address{
			provice: "北京",
			city:    "昌平",
		},
	}
	fmt.Println(stu1)
	fmt.Println(stu1.ads.provice)

	var stu2 = student1{
		name: "kelly",
		age:  18,
		address: address{
			provice: "北京",
			city:    "昌平",
		},
	}
	fmt.Println(stu2)
	fmt.Println(stu2.provice) //匿名字段支持直接访问

}
