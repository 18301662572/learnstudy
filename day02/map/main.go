package main

import (
	"fmt"
	"math/rand"
	"sort"
)

//map
func main() {

	//声明
	//声明一个map类型，key是string类型，value是int类型
	var a map[string]int
	//只声明一个,map类型，但是没有初始化 ，初始值是nil
	fmt.Println(a == nil)

	//map初始化
	a = make(map[string]int, 8)
	fmt.Println(a == nil)

	//map中添加键值对,赋值
	a["沙河娜扎"] = 100
	a["沙河小王子"] = 200
	fmt.Println(a)
	fmt.Printf("a=%#v\n", a)
	fmt.Printf("a=%T\n", a)

	//map声明并初始化
	b := map[int]bool{
		1: true,
		2: false,
	}
	fmt.Printf("b：%#v\n", b)
	fmt.Printf("b：%T\n", b)

	//var c map[int]int
	//c[100] = 200 //c没有初始化，没有内存地址，不鞥你直接赋值
	//fmt.Println(c)

	//判断某个键存不存在 v ok:=
	var score = make(map[string]int, 8)
	score["张三"] = 10
	score["李四"] = 100
	//score["王五"] = 1000
	//判断某个是是否存在score中
	v, ok := score["王五"] //v 存在返回value，不存在返回0，ok 返回true/false
	if ok {
		fmt.Println("王五在，分数为", v)
	} else {
		fmt.Println("没有王五")
	}

	//map遍历 k键 v值
	//map是无序的，键值对和添加的顺序无关
	for k, v := range score {
		fmt.Println(k, v)
	}
	//只遍历map中的key
	for k := range score {
		fmt.Println(k)
	}
	//只遍历map中的value
	for _, v := range score {
		fmt.Println(v)
	}

	//map中删除键值对
	delete(score, "张三")
	fmt.Println(score)

	//按照某个固定顺序遍历map
	var scoreaa = make(map[string]int, 100)
	//添加键值对
	for i := 0; i < 50; i++ {
		key := fmt.Sprintf("stu%2d", i)
		value := rand.Intn(100) //动态生成0-99的随机整数
		scoreaa[key] = value
	}
	for key, v := range scoreaa {
		fmt.Println(key, v)
	}
	//按照key从小到大的顺序去遍历scoreaa
	//1.先取出所有的key存放到切片中
	keys := make([]string, 0, 100)
	for k := range scoreaa {
		keys = append(keys, k)
	}
	//2.对key做排序
	sort.Strings(keys) //keys目前是一个有序的切片
	//3.按照排序后的key对scoreaa排序
	for _, key := range keys {
		fmt.Println(key, scoreaa[key])
	}

}
