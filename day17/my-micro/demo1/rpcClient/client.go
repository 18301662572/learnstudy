package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Params struct {
	Width,Height int
}

//调用服务
func main()  {
	//1.连接远程RPC服务
	rpcClient,err:= rpc.DialHTTP("tcp","127.0.0.1:8080")
	if err!=nil{
		log.Fatal(err)
	}
	//2.调用远程方法
	//定义接收服务端传回来的计算结果的变量
	ret:=0
	//求面积
	//第一个参数：服务方法名，第二个参数传的参数，第三个参数是服务端返回给客户端的接受值（接受值是指针类型）
	err=rpcClient.Call("Rect.Area",Params{50,20},&ret)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("面积：",ret)
	//求周长
	err=rpcClient.Call("Rect.Perimeter",Params{50,20},&ret)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("周长：",ret)
}
