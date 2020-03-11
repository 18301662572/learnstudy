package main

import "fmt"

//类型断言
//type nullInterfcae interface{}//定义空接口太繁琐

func showType(x interface{}) {
	//因为我这个函数可以接受任意类型的变量
	//类型断言
	v, ok := x.(int)
	if !ok {
		fmt.Println("x不是一个int类型")
	} else {
		fmt.Println("x就是一个int类型", v)
	}
	v1, ok := x.(string)
	if !ok {
		//说明猜错了
		fmt.Println("x不是一个string类型")
	} else {
		fmt.Println("x就是一个string类型", v1)
	}

}

func justifyType(x interface{}) {
	switch v := x.(type) {
	case string:
		fmt.Printf("x is a string value is %v\n", v)
	case int:
		fmt.Printf("x is a int is %v\n", v)
	case bool:
		fmt.Printf("x is a bool is %v\n", v)
	case *string:
		fmt.Printf("x is a *string，value is %v\n", v)
	default:
		fmt.Println("unsupport type!")
	}
}

func main() {
	// var x interface{}
	// x = 100
	showType(100)
	showType("534")
	justifyType(100)
	justifyType("534")
	a := "哈哈"
	justifyType(&a)
}
