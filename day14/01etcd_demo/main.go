package main

import (
	"context"
	"fmt"
	"time"
	"go.etcd.io/etcd/clientv3"
)

// 代码连接 etcd

// etcd client put/get demo
// use etcd/clientv3
func main(){
	cli, err := clientv3.New(
		clientv3.Config{
			Endpoints: []string{"127.0.0.1:2379"},
			DialTimeout: 5 * time.Second,
		})
	if err != nil {
		// handle error!
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		return
	}
	fmt.Println("connect to etcd success")
	defer cli.Close()
	// put
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	//str:=`[{"path":"E:/logs/s4.log","topic":"s4"},{"path":"E:/logs/web.log","topic":"web_log"}]`
	str:=`[{"path":"E:/logs/s4.log","topic":"s4"},{"path":"F:/logs/web.log","topic":"web_log"},{"path":"G:/logs/addlog.log","topic":"add_log"}]`
	_, err = cli.Put(ctx, "collect_log_192.168.1.5_conf", str)
	if err != nil {
		fmt.Printf("put to etcd failed, err:%v\n", err)
		return
	}
	cancel()
	// get
	ctx, cancel = context.WithTimeout(context.Background(), time.Second*2)
	resp, err := cli.Get(ctx, "collect_log_192.168.1.5_conf")
	cancel()
	if err != nil {
		fmt.Printf("get from etcd failed, err:%v\n", err)
		return
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s:%s\n", ev.Key, ev.Value)
	}
}
