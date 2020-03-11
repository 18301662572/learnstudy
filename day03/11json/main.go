package main

import (
	"encoding/json"
	"fmt"
)

//json序列化

//Student 是一个结构体
//定义元信息: json tag //当json包处理Student，使用tag，转成小写
type Student struct {
	ID     int    `json:"id"`
	Gender string `json:"gender"`
	Name   string `json:"name"`
}

func main() {
	var stu1 = Student{
		ID:     1,
		Gender: "男",
		Name:   "娜扎",
	}
	//序列化：把编程语言里面的数据转换成json格式的字符串
	v, err := json.Marshal(stu1)
	if err != nil {
		fmt.Println("json格式化出错")
		fmt.Println(err)
	}
	fmt.Println(v) //[]byte
	fmt.Printf("%#v:\n", string(v))
	fmt.Println(string(v)) //把[]byte转化成string
	//反序列化：把满足json格式的字符串转换成满足当前编程语言里面的对象
	str := "{\"ID\":1,\"Gender\":\"男\",\"Name\":\"娜扎\"}"
	stu2 := &Student{}
	json.Unmarshal([]byte(str), stu2)
	fmt.Println(stu2)
	fmt.Printf("%T\n", stu2)
}
