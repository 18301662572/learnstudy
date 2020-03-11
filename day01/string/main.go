package main

import (
	"fmt"
	"strings"
)

//字符串操作
func main() {
	s1 := "hello word"
	fmt.Println(len(s1))

	//字符串拼接
	s2 := "Python"
	fmt.Println(s1 + s2)

	//Sprintf字符串拼接，返回一个新的字符串
	s3 := fmt.Sprintf("%s---%s", s1, s2)
	fmt.Println(s3)

	//分割字符串
	s4 := "dasdsad"
	s5 := strings.Split(s4, "s")
	fmt.Println(s5)

	fmt.Println(strings.Contains(s4, "aa")) //字符串是否包含字符
	fmt.Println(strings.HasPrefix(s4, "d")) //是否以字符开头
	fmt.Println(strings.HasSuffix(s4, "s")) //是否以字符结尾
	fmt.Println(strings.Index(s4, "s"))     //字符所在的第一个索引值
	fmt.Println(strings.LastIndex(s4, "s")) //字符所在的最后一个索引值
	fmt.Println(strings.Index(s1, "s"))     //-1
	fmt.Println(strings.LastIndex(s1, "s")) //-1

	//join 连接
	a1 := []string{"Python", "PHP", "Golang", "Ruby", "JS"}
	fmt.Println(strings.Join(a1, "-"))

}
