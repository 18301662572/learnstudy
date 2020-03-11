package main

import (
	"fmt"
	"reflect"
)

//通过反射修改值

func modifyValue(x interface{}) {
	v := reflect.ValueOf(x) //v：reflect.Value类型
	kind := v.Kind()
	//fmt.Println(kind) //ptr 指针
	if kind == reflect.Ptr {
		//传入的是一个指针
		//取到指针对应的值再去修改
		v.Elem().SetInt(1200)
	}
}

func main() {
	var a int64 = 100
	//将a传入一个函数，在函数中修改a的值，该函数可接受任意类型的值
	modifyValue(&a) //传入的是指针
	fmt.Println(a)  //1200
}
