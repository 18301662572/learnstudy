package main

import (
	"code.oldbody.com/studygolang/learnstuday/day16/log_transfer/es"
	"code.oldbody.com/studygolang/learnstuday/day16/log_transfer/kafka"
	"code.oldbody.com/studygolang/learnstuday/day16/log_transfer/model"
	"fmt"
	"github.com/go-ini/ini"
)

//log transfer
//从kafka消费日志数据，写入ES

func main() {
	//1.加载配置文件
	var cfg = new(model.Config)
	err := ini.MapTo(cfg, "./config/logtransfer.ini")
	if err != nil {
		fmt.Println("load config failed,err:%v", err)
		panic(err)
	}
	fmt.Println("load config success")
	//2.链接ES
	err = es.Init(cfg.ESConf.Address, cfg.ESConf.Index, cfg.ESConf.GoNum, cfg.ESConf.MaxSize)
	if err != nil {
		fmt.Println("init ES failed,err:%v", err)
		panic(err)
	}
	fmt.Println("init ES success")
	//3.链接kafka
	err = kafka.Init([]string{cfg.KafkaConf.Address}, cfg.KafkaConf.Topic)
	if err != nil {
		fmt.Println("connect to kafka failed,err:%v", err)
		panic(err)
	}
	fmt.Println("connect kafka success")
	//select 不占用CPU，让程序在这儿停顿！！
	select {}
}
