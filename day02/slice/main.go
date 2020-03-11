package main

import "fmt"

//Slice 切片
func main() {
	//数组
	a := [3]int{1, 2, 3}
	//切片
	//方式1
	//声明并初始化
	b := []int{1, 2, 3}
	fmt.Println(a, b)
	fmt.Printf("a:%T\n", a)
	fmt.Printf("b:%T\n", b)
	fmt.Printf("b的长度：%d\n", len(b)) //3
	fmt.Printf("b的容量：%d\n", cap(b)) //3
	//切片取值,地址
	fmt.Println(b[2])

	//方式2:从数组得到切片,地址
	var c []int
	c = a[0:2] //c=a[:]从数组开始切到结束
	fmt.Printf("c:%T\n", c)
	fmt.Printf("c:%v\n", c)
	//[ : ] 左包含右不包含
	d := a[:2] //从开始切到索引2（不包含2）
	fmt.Println(d)

	//切片的大小（目前元素的数量）
	fmt.Println(len(d))

	//容量（底层数组最大能放多少）
	x := [...]string{"北京", "上海", "深圳", "成都", "重庆"}
	y := x[1:4]
	fmt.Println(y)
	fmt.Printf("切片y的长度是:%d\n", len(y)) //3
	fmt.Printf("切片y的容量是:%d\n", cap(y)) //4 从索引1 开始往后的数组

}
