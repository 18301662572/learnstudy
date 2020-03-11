package main

import (
	"fmt"
	"reflect"
)

//反射的TypeOf()

type cat struct {
	name string
}

func reflectType(x interface{}) {
	v := reflect.TypeOf(x) //拿到x的动态类型信息
	// fmt.Printf("type:%v\n", v) //能拿到数据类型
	// fmt.Printf("%T\n", x) //fmt.Printf 原理就是用的反射，代码补全也是用的反射
	fmt.Printf("name:%v kind:%v\n", v.Name(), v.Kind()) //Name()具体的数据类型，Kind()大的种类,主要针对复杂数据类型
	//复杂类型的v.Name() 是空，但是 fmt.Printf("type:%v\n", v)，能拿到该数据类型
}

func main() {
	// reflectType(100)
	// reflectType(false)
	// reflectType("dasdad")
	// reflectType([3]int{1, 2, 3})
	// reflectType(map[string]int{})

	//测试自定义的结构类型
	c1 := cat{name: "maomao"}
	reflectType(c1)

	var n1 int32 = 100
	reflectType(&n1)
	n2 := "豪杰"
	reflectType(&n2)
	reflectType([]int{1, 2, 3})
	reflectType(map[string]int{})
	reflectType(123)

}
