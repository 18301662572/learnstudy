package main

import (
	"fmt"
	"reflect"
)

//结构体反射示例

//Student 一个自定义类型
type Student struct {
	Name  string `json:"name" hj:"mz"`
	Score int    `json:"score" hj:"fs"`
}

func main() {
	stu1 := Student{
		Name:  "haojie",
		Score: 0,
	}

	t := reflect.TypeOf(stu1)
	fmt.Println(t.Name(), t.Kind()) // student struct
	// 方法1：通过for循环遍历结构体的所有字段信息
	//NumField() 返回结构体成员字段数量
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", field.Name, field.Index, field.Type, field.Tag.Get("json"))
		//field.Tag.Get("hj")
	}

	// 方法2：通过字段名获取指定结构体字段信息
	if scoreField, ok := t.FieldByName("Score"); ok {
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", scoreField.Name, scoreField.Index, scoreField.Type, scoreField.Tag.Get("json"))
	}
}
