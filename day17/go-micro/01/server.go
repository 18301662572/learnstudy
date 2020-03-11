package main

import (
	"context"
	"fmt"
	"github.com/go-micro"
	"log"
	pb "code.oldbody.com/studygolang/learnstuday/day17/go-micro/01/proto"
)

//声明结构体
type Hello struct {
}

//实现接口方法
func (g *Hello) Info(ctx context.Context,req *pb.InfoRequest,rep *pb.InfoResponse)error {
	rep.Msg="你好"+req.Username
	return nil
}

func main(){
	//1.得到微服务实例
	service:= micro.NewService(
		//设置微服务的名字，用来做访问用的
		micro.Name("hello"),
		)
	//2.初始化
	service.Init()
	//3.服务注册
	err:=pb.RegisterHelloHandler(service.Server,new(Hello))
	if err!=nil{
		fmt.Println(err)
	}
	//4.启动微服务
	if err=service.Run();err!=nil{
		log.Fatal(err)
	}
}
