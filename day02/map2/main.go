package main

import (
	"fmt"
	"strings"
)

//map映射
func main() {

	//元素类型为map的切片
	//1.只是完成切片的初始化,没有对map进行初始化
	var mapSlice = make([]map[string]int, 8, 8)
	//[nil nil nil nil nil nil nil nil]
	fmt.Println(mapSlice[0] == nil) //目前的map没有初始化
	fmt.Println(mapSlice)
	//2.还需要完成内部map元素的初始化
	mapSlice[0] = make(map[string]int, 8) //完成了map的初始化
	mapSlice[0]["沙河"] = 100
	fmt.Println(mapSlice)
	mapSlice[1] = make(map[string]int, 8)
	mapSlice[1]["沙河1"] = 200
	fmt.Println(mapSlice)

	//值为切片的map
	//1.只完成了map的初始化
	var sliceMap = make(map[string][]int, 8)
	v, ok := sliceMap["中国"]
	if ok {
		fmt.Println(v)
	} else {
		//sliceMap 中没有中国这个键
		sliceMap["中国"] = make([]int, 8) //完成了对切片的初始化
		sliceMap["中国"][0] = 100
		sliceMap["中国"][1] = 200
		sliceMap["中国"][3] = 300
	}
	//遍历sliceMap
	for k, v := range sliceMap {
		fmt.Println(k, v)
	}

	//作业：
	//统计一个字符串中每个单词出现的次数
	test := "sda daqw sda 22"
	//0.定义一个map[string]int
	var count = make(map[string]int, 20)
	//1.字符串中都有哪些单词
	words := strings.Split(test, " ")
	//2.遍历单词做统计
	for _, word := range words {
		v, ok := count[word]
		if ok {
			//map中有这个单词
			count[word] = v + 1
		} else {
			count[word] = 1
		}
	}
	for k, v := range count {
		fmt.Println(k, v)
	}

}
