package main

import "fmt"

func main() {

	//一维数组
	//声明
	var a [5]int  //定义一个长度为5存放int类型的数组
	var b [10]int //定义一个长度为10存放int类型的数组
	//初始化
	a = [5]int{1, 2, 3, 4, 5}
	b = [10]int{1, 2, 3, 4, 5}
	fmt.Println(a)
	fmt.Println(b)

	//声明并赋值
	var c [3]string = [3]string{"11", "22"}
	var c1 = [3]string{"11", "22"}
	c2 := [3]string{"11", "22"}
	fmt.Println(c)
	fmt.Println(c1)
	fmt.Println(c2)

	// ... 表示让编译器数一下有多少个初始值，然后赋值给变量数据类型
	var d = [...]int{1, 2, 3, 44, 55}
	fmt.Println(d)
	fmt.Printf("d=%T\t c=%T\n", d, c)

	//根据索引值初始化
	var e [20]int
	e = [20]int{19: 1}
	fmt.Println(e)

	//数组的基本使用
	fmt.Println(e[3])

	//遍历数组方法1
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
	//方法2
	//for range
	for index, value := range a {
		fmt.Println(index, a[index], value)
	}
	for index := range a {
		fmt.Println(index, a[index])
		fmt.Printf("index=%v,value=%v\n", index, a[index])
	}

	//求所有数组元素的和
	ff := [...]int{1, 3, 5, 7, 8}
	var values = 0
	for _, value := range ff {
		values += value
	}
	fmt.Printf("ff数组之和为:%v\n", values)
	//找出ff数组中和为8的两个数的索引值
	//方法1
	for i := 0; i < len(ff); i++ {
		for j := 0; j < len(ff); j++ {
			if ff[i]+ff[j] == 8 && i != j {
				fmt.Println(i, j)
			}
		}
	}
	//方法2
	for index1, value1 := range ff {
		other := 8 - value1
		for index2 := index1 + 1; index2 < len(ff); index2++ {
			if other == ff[index2] {
				fmt.Printf("索引为%d,%d\n", index1, index2)
			}
		}
	}

	//多维数组
	//声明
	var aa [3][2]int
	//初始化
	aa = [3][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
	}
	fmt.Println(aa)

	//声明并初始化
	bb := [...][2]int{
		{1, 2},
		{5, 7}}
	cc := [3][2]int{
		{1, 2},
		{5, 7},
	}
	fmt.Println(bb)
	fmt.Println(cc)

	//注意事项：多维数组除了第一层能用...,其他层都不能用...

	//多位数组的使用
	fmt.Println(cc[1][1])

	//多位数组的遍历
	for i := 0; i < len(cc); i++ {
		for j := 0; j < len(cc[i]); j++ {
			fmt.Printf("%d-%d-%d\n", i, j, cc[i][j])
		}
	}
	for _, value := range cc {
		fmt.Println(value)
		for _, value2 := range value {
			fmt.Println(value2)
		}
	}

	//数组是值类型
	va1 := [2]int{1, 2}
	va2 := va1 //把va1整个拷贝一份赋值给va2
	va2[0] = 100
	fmt.Println(va1) //[1 2]
	fmt.Println(va2) //[100 2]

	ii := [3][2]int{
		[2]int{1, 2},
	}
	hh := ii //值类型，赋值的时候都是直接拷贝
	hh[2][0] = 100
	fmt.Println(ii)
	fmt.Println(hh)

}
