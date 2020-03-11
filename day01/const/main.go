package main

import "fmt"

const Pi = 3.1415 //常量在定义的时候必须赋值

//批量声明常量
const (
	a = 100
	b = 10000
	c //c=1000
	d //d=1000
)

//iota 枚举,默认为0，在同一个const中依次累加1
const (
	aa = iota //0
	bb        //iota+=1 ->1
	cc        //iota+=1 ->2
	dd        //iota+=1 ->3
)

const (
	n1 = iota //0
	n2        //=iota ->1
	_         // =iota ->2
	n4        // =iota->3
)

//const 中每新增一行常量声明，将使iota计数一次（iota 可理解为const语句块中行的索引）
const (
	i1 = iota //0
	i2 = 100  //100
	i3 = iota //2
	i4        //3
)
const i5 = iota //0

// << 位移运算符
const (
	_  = iota             //0
	KB = 1 << (10 * iota) // 1<<10 (1位移10位)=>2的10次方 =1024
	MB = 1 << (10 * iota) // 1<<20 =1024*1024
	GB = 1 << (10 * iota) // 1<<30 =1024*1024*1024
	TB = 1 << (10 * iota) // 1<<40 =1024*1024*1024*1024
	PB = 1 << (10 * iota) // 1<<50 =1024*1024*1024*1024*1024
)

const (
	g, h = iota + 1, iota + 2 // iota=0 g=1,h=2
	m, n                      // m,n =iota+1,iota+2  iota=1 m=2,n=3
	e, f                      // e,f=iota+1,iota+2  iota=2 e=3,f=4
)

//常量 ,常量是不允许修改的
func main() {
	fmt.Println(Pi)
	// var Pi = "dasda"
	// fmt.Println(Pi)
	fmt.Println(a, b, c, d)
	fmt.Println(aa, bb, cc, dd)
	fmt.Println(n1, n2, n4)
	fmt.Println(i1, i2, i3, i4, i5)
}
