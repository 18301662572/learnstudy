package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

//粘包现象  server端

func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn) //创建一个缓冲区读数据
	var buf [1024]byte
	//循环读
	for {
		n, err := reader.Read(buf[:])
		if err != io.EOF {
			break
		}
		if err != nil {
			fmt.Println("read from client failed,err:", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("收到client发来的数据：", recvStr)
	}
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("listen failed,err:", err)
		return
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed,err:", err)
			continue
		}
		go process(conn)
	}
}
