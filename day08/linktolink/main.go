package main

import "fmt"

//链式操作  原理：每一次执行完方法之后，返回操作的对象本身

type student struct {
	name string
}

func (s student) learn() student {
	fmt.Printf("%s热爱学习\n", s.name)
	return s
}

func (s student) doHomework() student {
	fmt.Printf("%s交作业\n", s.name)
	return s
}

func main() {
	stu := student{
		"豪杰",
	}
	stu.learn().doHomework()
}
