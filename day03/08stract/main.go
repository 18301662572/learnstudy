package main

import "fmt"

//结构体的匿名字段
type stu struct {
	name string
	string
	int
}

func main() {
	var a1 = stu{
		name:   "123",
		string: "dasd",
		int:    12,
	}
	fmt.Println(a1.name)
	fmt.Println(a1.string)
	fmt.Println(a1.int)
}
