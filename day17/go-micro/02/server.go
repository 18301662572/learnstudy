package main

import (
	"fmt"
	"github.com/micro/go-micro"
	pb "go-micro/02/proto"
	"context"
	"log"
	"errors"
)

type Example struct {
}

type Foo struct {
}

func (e *Example) Call(ctx context.Context, req *pb.CallRequest, rep *pb.CallResponse) error {
	log.Print("收到 Example Call 请求")
	if len(req.Name)==0{
		return errors.BadRequest("go.micro.api.example","no content")
	}
	rep.Message="RPC Call 收到了你的请求"+req.Name
	return nil
}

func (e *Example) Bar(ctx context.Context, req *pb.EmptyRequest, rep *pb.EmptyResponse) error {
	log.Print("收到 Foo.Bar 请求")
	return nil
}

func main(){
	service:=micro.NewService(
		micro.Name("go.micro.api.example"),
		)
	service.Init()
	//注册example接口
	err:=pb.RegisterExampleHandler(service.Server().new(Example))
	if err!=nil{
		fmt.Println(err)
	}
	//注册Foo接口
	err =pb.RegisterFooHandler(service.Server().new(Foo))
	if err!=nil{
		fmt.Println(err)
	}
	if err:=service.Run(); err!=nil{
		log.Fatal(err)
	}

}