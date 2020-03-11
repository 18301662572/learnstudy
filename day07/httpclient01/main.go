package main

import (
	"fmt"
	"io"
	"net"
)

//HTTP Client

func main() {
	conn, err := net.Dial("tcp", "www.liwenzhou.com:80") //http
	if err != nil {
		fmt.Println("访问 www.liwenzhou.com 失败")
		return
	}
	defer conn.Close()
	// 发送数据
	conn.Write([]byte("GET / HTTP/1.0\r\n\r\n"))
	//接收数据
	var buf [1024]byte
	for {
		n, err := conn.Read(buf[:])
		if err == io.EOF {
			fmt.Println(string(buf[:n]))
			return
		}
		if err != nil {
			fmt.Println("接收数据失败，err:", err)
		}
		fmt.Println(string(buf[:n]))
	}

}
