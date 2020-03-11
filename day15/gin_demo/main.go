package main

import (
	"fmt"
	"net/http"
)

//HTTP Server端

func sayHello(w http.ResponseWriter, t *http.Request) {
	fmt.Fprint(w,"Hello shahe!")
}

func main() {
	http.HandleFunc("/", sayHello) //注册一个处理/ 的函数
	//启动服务
	err := http.ListenAndServe("127.0.0.1:9090", nil)
	if err != nil {
		panic(err)
	}
}

