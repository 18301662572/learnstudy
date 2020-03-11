package main

import "fmt"

//从终端获取用户的输入内容
//fmt.Scan包

func main() {
	var (
		name    string
		age     int
		married bool
	)

	//从终端获取用户的输入内容 :fmt.Scan包
	//可以用 换行/空格 填写下一个
	//fmt.Scan(&name, &age, &married)
	//按照格式输入
	//fmt.Scanf("name:%s age:%d married:%t\n", &name, &age, &married)
	//不能用换行,只能用空格填写下一个
	//fmt.Scanln(&name, &age, &married)

	//终端输出函数
	fmt.Println(name, age, married)
	//fmt.Printf("",)
}
