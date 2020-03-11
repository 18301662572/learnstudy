package main

import (
	"fmt"
	"sort"
)

func main() {
	var a []int
	var b []string
	c := []bool{false, true}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	//make函数构造切片
	//类型 切片长度 切片容量，第三个参数可以省略
	d := make([]int, 5, 10)
	fmt.Println(d)
	fmt.Printf("d=%T\n", d)

	f := make([]int, 5)
	fmt.Println(f)
	fmt.Printf("f=%T\n", f)

	//nil 切片
	var aa []int     //只有声明没有初始化,。没有创建内存地址，此时a就是nil
	var bb = []int{} //声明并初始化,创建了内存地址
	cc := make([]int, 0)
	if aa == nil {
		fmt.Println("aa是一个nil")
	}
	if bb == nil {
		fmt.Println("bb是一个nil")
	}
	if cc == nil {
		fmt.Println("cc是一个nil")
	}
	fmt.Println(aa, len(aa), cap(aa))
	fmt.Println(bb, len(bb), cap(bb))
	fmt.Println(cc, len(cc), cap(cc))

	//切片的赋值拷贝
	aaa := make([]int, 3) //[0,0,0]
	bbb := aaa            //aaa,bbb公用一个底层数组
	bbb[0] = 1
	fmt.Println(aaa)
	fmt.Println(bbb)

	//切片的遍历
	aaaa := []int{1, 2, 3, 4, 5}
	//根据索引遍历
	for i := 0; i < len(aaaa); i++ {
		fmt.Println(i, aaaa[i])
	}
	fmt.Println()
	//根据 for range 遍历
	for index, value := range aaaa {
		fmt.Println(index, value)
	}

	//append动态给切片添加元素
	var aaaaa []int           //声明时，并没有申请内存
	aaaaa = append(aaaaa, 10) //append可能会使切片扩容，所以一定要有一个接收值接收
	fmt.Println(aaaaa)
	for i := 0; i < 10; i++ {
		aaaaa = append(aaaaa, i)
		fmt.Printf("%v len%d cap%d ptr%p\n", aaaaa, len(aaaaa), cap(aaaaa), aaaaa) //值 长度 容量 内存
	}
	//append一次追加多个元素
	var bbbbb []int
	bbbbb = append(bbbbb, 1, 2, 3, 4, 5, 6)
	fmt.Println(bbbbb)

	ccccc := []int{12, 13, 14, 15}
	bbbbb = append(bbbbb, ccccc...)
	fmt.Println(bbbbb)

	//切片的copy
	x := []int{1, 2, 3, 4}
	y := []int{5, 6, 7}
	z := y
	copy(y, x) //将x的元素拷贝到y中
	y[0] = 100
	fmt.Printf("x：%v\n", x)
	fmt.Printf("y：%v\n", y)
	fmt.Printf("z：%v\n", z)

	//切片删除元素
	m := []string{"123", "456", "789", "43242"}
	m = append(m[:1], m[2:]...)
	fmt.Println(m)

	//使用 sort 函数对数组进行排序
	var g = [...]int{3, 7, 8, 9, 1}
	//g[:]得到的是一个切片，指向了底层的数组g
	sort.Ints(g[:])
	fmt.Println(g)

}
