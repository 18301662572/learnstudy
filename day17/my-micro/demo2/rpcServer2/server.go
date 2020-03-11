package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

//声明算术运算结构体
type Arith struct {
}

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

//乘法运算
func (this *Arith) Mulriply(req ArithRequest,res *ArithResponse)error{
	res.Pro=req.A*req.B
	return nil
}

//商和余数
func (this *Arith) Divide(req ArithRequest,res *ArithResponse)error{
	if req.B==0{
		return errors.New("除数不能为0")
	}
	//商
	res.Quo=req.A/req.B
	//余数
	res.Rem=req.A%req.B
	return nil
}

//golang的RPC调用
//func main(){
//	//注册服务
//	rpc.Register(new(Arith))
//	//采用http作为rpc载体
//	rpc.HandleHTTP()
//	err:=http.ListenAndServe(":8081",nil)
//	if err!=nil{
//		log.Fatal(err)
//	}
//}

//jsonRPC编码方式调用，跨语言
func main(){
	//注册服务
	rpc.Register(new(Arith))
	//监听服务
	lis,err:=net.Listen("tcp","127.0.0.1:8081")
	if err!=nil{
		log.Fatal(err)
	}
	//循环监听服务
	for{
		conn,err:=lis.Accept()
		if err!=nil{
			continue
		}
		//协程
		go func(conn net.Conn){
			fmt.Println("new Client")
			jsonrpc.ServeConn(conn)
		}(conn)
	}
}