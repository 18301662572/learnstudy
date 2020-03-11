package main

//关于defer的面试题

func f1() int {
	x := 5
	defer func() {
		x++ //只是修改了函数里面的x，并没有返回值赋值给x
	}()
	return x //1.（汇编）返回值=x x=5, 2.执行defer 3.（汇编）RET=x
} //返回5

func f2() (x int) { //全局的x
	defer func() {
		x++
	}()
	return 5 //1.（汇编）返回值=x x=5  2.x++ 3.（汇编）RET==>6
} //返回6

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x //1.（汇编）返回值=y y=5 2.x++ 3.（汇编）RET=>5
} //返回5

func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5 //1.（汇编）返回值=x x=5 2.x是函数内部的x 3.（汇编）RET=>5
} //5

func main() {

}
