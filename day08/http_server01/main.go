package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//HTTP Server端

func sayHello(w http.ResponseWriter, t *http.Request) {
	// fmt.Fprint(w,"Hello shahe!")
	// w.Write([]byte("Hello shahe!"))

	//从hello.txt文件中读取数据写入到w中
	data, err := ioutil.ReadFile("./hello.html") //./hello.txt
	if err != nil {
		fmt.Println("read from failed,err:", err)
		return
	}
	w.Write(data)
}

func main() {
	http.HandleFunc("/", sayHello) //注册一个处理/ 的函数
	//启动服务
	err := http.ListenAndServe("127.0.0.1:9090", nil)
	if err != nil {
		panic(err)
	}
}
