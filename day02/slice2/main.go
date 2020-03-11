package main

import "fmt"

func main() {

	//切片支持自动扩容（扩容策略：每一次都是上一次的2倍）
	var a = []int{}
	fmt.Println(a, len(a), cap(a))
	fmt.Printf("a:%v len:%d cap:%d 内存地址ptr:%p\n", a, len(a), cap(a), a)
	a = append(a, 1)
	fmt.Println(a, len(a), cap(a))
	fmt.Printf("a:%v len:%d cap:%d 内存地址ptr:%p\n", a, len(a), cap(a), a)
	a = append(a, 2)
	fmt.Println(a, len(a), cap(a))
	fmt.Printf("a:%v len:%d cap:%d 内存地址ptr:%p\n", a, len(a), cap(a), a)
	a = append(a, 1)
	fmt.Println(a, len(a), cap(a))
	fmt.Printf("a:%v len:%d cap:%d 内存地址ptr:%p\n", a, len(a), cap(a), a)
	a = append(a, 1)
	fmt.Println(a, len(a), cap(a))
	fmt.Printf("a:%v len:%d cap:%d 内存地址ptr:%p\n", a, len(a), cap(a), a)
	a = append(a, 1)
	fmt.Println(a, len(a), cap(a))
	fmt.Printf("a:%v len:%d cap:%d 内存地址ptr:%p\n", a, len(a), cap(a), a)
	a = append(a, 1)
	fmt.Println(a, len(a), cap(a))
	fmt.Printf("a:%v len:%d cap:%d 内存地址ptr:%p\n", a, len(a), cap(a), a)

	//切片是引用类型
	aa := []int{1, 2, 3}

	bb := aa //1.直接赋值--aa,bb同占一个内存空间

	var cc []int           //2.声明一个切片，还没有申请内存---aa,cc不在同一个内存空间
	cc = make([]int, 3, 3) //  申请内存
	copy(cc, aa)           //  深拷贝

	bb[1] = 100
	fmt.Println(aa)
	fmt.Println(bb)
	fmt.Println(cc)

	//切片中删除某个元素
	aaa := []string{"北京", "上海", "深圳", "成都"}
	//aaa[:1]=["北京"]
	//aaa[2:]=["深圳", "成都"]
	aaa = append(aaa, "广州") //将字符串追加到数组里
	//aaa[2:] 是一个切片，不能追加到数组或切片中
	//...可以将后面的多个元素（字符串）一块追加
	aaa = append(aaa[:1], aaa[2:]...)
	fmt.Println(aaa)

}
