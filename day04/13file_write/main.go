package main

import (
	"fmt"
	"os"
)

//打开文件支持文件写入

func main() {
	file, err := os.OpenFile("xx.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND|os.O_TRUNC, 0755)
	if err != nil {
		fmt.Println("打开文件失败", err)
		return
	}
	defer file.Close()
	str := "Hello Shahe"
	file.Write([]byte("哈哈\n"))
	file.WriteString(str)
	fmt.Println()
}
