package main

import (
	"fmt"
	"os"
)

//os.Args

func main(){
	fmt.Println(os.Args)
	fmt.Println("Hello",os.Args[1:])
}
