package main

import "fmt"

//switch 的坑

func main() {
	f := func() bool {
		return false
	}

	switch f(); { //switch _=f();  //switch不写变量， 默认走true
	case true:
		fmt.Println("真")
	case false:
		fmt.Println("假")
	} // --真

	switch f() {
	case true:
		fmt.Println("真")
	case false:
		fmt.Println("假")
	} // --假
}
