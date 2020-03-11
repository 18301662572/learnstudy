package main

import "fmt"

//结构体
//创造新的类型使用 type关键字
type student struct {
	name   string
	age    int
	gender string
	hobby  []string
}

//实例化：声明 初始化：赋值
func main() {
	//实例化方法1
	var haojie = student{
		name:   "豪杰",
		age:    19,
		gender: "男",
		hobby:  []string{"篮球", "足球"},
	}
	//结构体支持 . 访问属性
	fmt.Println(haojie)
	fmt.Println(haojie.name)
	fmt.Println(haojie.age)
	fmt.Println(haojie.hobby)

	//实例化
	//实例化方法1
	//struct 是值类型
	//如果初始化时没有给属性（字段）设置对应的初始值，那么对应的属性类型就是其类型的默认值
	var wangzan = student{}
	fmt.Println(wangzan.name)
	fmt.Println(wangzan.age)
	fmt.Println(wangzan.gender)
	fmt.Println(wangzan.hobby)
	//实例化方法2
	//new(T) 分配内存，主要分配值类型，比如 int ,struct ,返回的是指针,即内存地址
	var dawei = new(student)
	dawei.name = "大卫" //相当于 (*dawei).struct --*dawei,取地址所对应的结构体
	dawei.age = 12
	fmt.Println(dawei.name, dawei.age)
	//实例化方法3
	var nazha = &student{} //naza:结构体的指针或地址
	(*nazha).name = "娜扎"
	nazha.age = 19
	fmt.Println(nazha.name, nazha.age, nazha.hobby)

	//初始化
	//结构体初始化1
	//键值对初始化
	var stu01 = student{
		name:  "姓名",
		age:   18,
		hobby: []string{"篮球", "足球"},
	}
	fmt.Println(stu01)
	//结构体初始化2
	//只填值初始化（必须填写所有的值，按顺序填写）
	var stu02 = student{
		"姓名1",
		181,
		"女1",
		[]string{"篮球1", "足球1"},
	}
	fmt.Println(stu02)
	//结构体初始化3 取地址
	//键值对初始化
	var stu04 = &student{
		name: "姓名3",
		age:  183,
	}
	fmt.Println(stu04)
	//结构体初始化4 取地址
	//只填值初始化（必须填写所有的值，按顺序填写）
	var stu03 = &student{
		"姓名2",
		182,
		"女2",
		[]string{"篮球2", "足球2"},
	}
	fmt.Println(stu03)

}
