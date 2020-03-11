package main

import "fmt"

func main() {
	age := 16
	if age > 18 {
		fmt.Printf("%d成年", age)
	} else if age < 18 {
		fmt.Printf("%d未成年", age)
	} else {
		fmt.Println("18岁")
	}

	if age2 := 28; age2 > 18 {
		fmt.Println("成年了")
	}
	//fmt.Println(age2) //局部变量不能输出
}
