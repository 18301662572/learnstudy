package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

//bufio读数据
func readByLine() {
	file, err := os.Open("../11file_open/xx.txt")
	if err != nil {
		fmt.Println("打开文见失败")
		return
	}
	defer file.Close()
	//利用缓冲区从文件读数据
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n') //单引号表示字符
		if err == io.EOF {
			fmt.Println(str)
			return
		}
		if err != nil {
			fmt.Println("读取文件内容失败")
			return
		}
		fmt.Print(str)
	}
}

//ioutil读取文件
func readFile(filename string) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("read file failed,err:", err)
		return
	}
	fmt.Println(string(content))
}

func main() {
	//bufio读取文件
	readByLine()
	fmt.Println()
	//ioutil读取文件
	readFile("../11file_open/xx.txt")

}
