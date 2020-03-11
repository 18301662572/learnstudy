package main

import "fmt"

//格式化打印
func main() {
	a := "a1"
	b := "b1"
	c := 100
	fmt.Println(a, b)
	//%v 俗称占位符
	fmt.Printf("a=%v\n", a)
	// %T 打印类型
	fmt.Printf("a=%T\n", a)
	fmt.Printf("a=%T\n", a)

	// %%转义
	fmt.Printf("%d%%\n", c)

	// %d
	fmt.Printf("|%d|\n", c)
	//+右边对齐，左边留白
	fmt.Printf("|%5d|\n", c)
	//+右边对齐，左边留白
	fmt.Printf("|%8d|\n", c)
	//-左边对齐，右边留白
	fmt.Printf("|%-8d|\n", c)
	//0前面补零
	fmt.Printf("|%08d|\n", c)

	// %f 小数点后几位
	f1 := 3.141567867
	fmt.Printf("%f\n", f1)
	fmt.Printf("%.2f\n", f1)
	fmt.Printf("%.5f\n", f1)

	C1 := "这是一个字符串\""
	fmt.Printf("%s\n", C1)
	fmt.Printf("%q\n", C1)

	fmt.Printf("%20s\n", C1)
	fmt.Printf("%.5s\n", C1)

}
