package main

import "fmt"

type student struct {
	name  string
	age   int
	sex   string
	hobby []string
}

//自己实现一个构造函数
//跟普通函数没有区别，命名用 new+结构体名 （驼峰命名法）
//返回值 student 是一个结构体，是值类型，每次都会拷贝一份给变量，占很多内存
func newStudent1(name string, age int, sex string, hobby []string) student {
	return student{
		name:  name,
		age:   age,
		sex:   sex,
		hobby: hobby,
	}
}

//返回结构体的地址，通过 *student 取地址
func newStudent(name string, age int, sex string, hobby []string) *student {
	return &student{
		name:  name,
		age:   age,
		sex:   sex,
		hobby: hobby,
	}
}

func main() {
	haojie := newStudent("豪杰", 12, "男", []string{"篮球", "求"})
	fmt.Println(haojie.name)
}
