package math_pkg

import "fmt"

//Add 函数
func Add(x, y int) int {
	return x + y
}

func init() {
	fmt.Println("测试一个init")
}

//Student 结构体
type Student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	id   int    //不能被外面的包访问
}
