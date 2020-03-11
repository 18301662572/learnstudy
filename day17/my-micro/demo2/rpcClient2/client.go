package main

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
)

//声明接受的参数结构体
type ArithRequest struct {
	A,B int
}

//声明返回客户端参数结构体
type ArithResponse struct {
	//乘积
	Pro int
	//商
	Quo int
	//余数
	Rem int
}

//调用服务
func main(){

	//1.链接远程的rpc
	//golang的RPC调用
	//conn,err:=rpc.DialHTTP("tcp","127.0.0.1:8081")
	//jsonRPC调用
	conn,err:=jsonrpc.Dial("tcp","127.0.0.1:8081")

	if err!=nil{
		log.Fatal(err)
	}
	req:=ArithRequest{9,2}
	var res ArithResponse

	err=conn.Call("Arith.Mulriply",req,&res)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Printf("%d * %d = %d\n",req.A,req.B,res.Pro)

	err=conn.Call("Arith.Divide",req,&res)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Printf("%d / %d, 商=%d, 余=%d \n",req.A,req.B,res.Quo,res.Rem)
}



