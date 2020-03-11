package main

import (
	"fmt"
	"reflect"
)

//反射的ValueOF

func reflectValue(x interface{}) {
	v := reflect.ValueOf(x) // 获取接口的值信息
	k := v.Kind()           //拿到值对应的种类
	switch k {
	case reflect.Int64:
		//v.Int()从反射中获取整形的原始值，然后通过int64()强制类型转换
		fmt.Printf("type is int64,value is %d\n", int16(v.Int()))
	case reflect.Float32:
		//v.Float() 从反射中获取浮点型的原始值，然后通过float32()强制类型转换
		fmt.Printf("type is Float32,value is %f\n", float32(v.Float()))
	case reflect.Float64:
		//v.Float() 从反射中获取浮点型的原始值，然后通过float64()强制类型转换
		fmt.Printf("type is Float64,value is %f\n", float64(v.Float()))
	}
}

func main() {
	var a float32 = 3.14
	var b int64 = 100
	reflectValue(a) //type is Float32,value is 3.140000
	reflectValue(b) //type is int64,value is 100
	//将int类型的原始值转换成reflect.Value类型
	c := reflect.ValueOf(10)
	fmt.Printf("type c:%T\n", c) //type c:reflect.Value

}
