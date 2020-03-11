package main

import "fmt"

func main() {

	a := [...]int{1, 3, 5, 7, 9, 11, 13, 14} //数组是值类型,存的是具体的值，不是地址
	b := a[:]                                //基于数组得到的切片，切片是引用类型
	b[0] = 100                               //切片修改了，数组同时也会修改，因为切片是基于数组的
	fmt.Println(a[0])                        //100
	fmt.Printf("b:%p\n", b)                  //内存 第一个元素是1，跟其他切片内存都不一样

	c := a[2:5]             //因为内存是连续不断的，切片从索引2开始到最后的数组是底层数组，然后再取到索引为5之前的数
	fmt.Println(c)          // 5, 7, 9
	fmt.Println(len(c))     //3
	fmt.Println(cap(c))     //容量 6  //底层 数组 最大能存放的元素的个数,跟内存有关
	fmt.Printf("c:%p\n", c) //内存 第一个元素是5，跟d 内存一样

	d := c[:2]
	fmt.Println(d)          //5,7
	fmt.Println(len(d))     //2
	fmt.Println(cap(d))     //容量 6  //
	fmt.Printf("d:%p\n", d) //内存  第一个元素是5，跟c 内存一样

	e := d[2:]
	fmt.Println(e)          //[]
	fmt.Println(len(e))     //0
	fmt.Println(cap(e))     //容量 4 //底层 数组 最大能存放的元素的个数
	fmt.Printf("e:%p\n", e) //内存 第一个元素是0，跟其他切片内存都不一样

	e = append(e, 100, 200, 300) //追加内容
	fmt.Println(e)
	fmt.Println(len(e))     //3
	fmt.Println(cap(e))     //容量扩充为原来的2倍 =8
	fmt.Printf("e:%p\n", e) //内存不够用，容量扩充，改变内存地址
}
