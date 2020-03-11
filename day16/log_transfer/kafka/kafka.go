package kafka

import (
	"code.oldbody.com/studygolang/learnstuday/day16/log_transfer/es"
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
)

//初始化kafka链接
//从kafka里面取出日志数据

func Init(address []string, topic string) (err error) {
	// 创建新的消费者
	consumer, err := sarama.NewConsumer(address, nil)
	if err != nil {
		fmt.Printf("fail to start consumer, err:%v\n", err)
		return
	}
	// 拿到指定topic下面的所有分区列表
	partitionList, err := consumer.Partitions(topic) // 根据topic取到所有的分区
	if err != nil {
		fmt.Printf("fail to get list of partition:err%v\n", err)
		return
	}
	fmt.Println(partitionList)
	for partition := range partitionList { // 遍历所有的分区
		// 针对每个分区创建一个对应的分区消费者
		var pc sarama.PartitionConsumer
		pc, err = consumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("failed to start consumer for partition %d,err:%v\n",
				partition, err)
			return
		}
		//defer pc.AsyncClose()
		//异步从每个分区消费信息
		fmt.Println("start to cumsume...")
		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				//logDataChan <- msg //为了将同步流程异步化，所以将取出的日志数据先放到通道中
				var m1 map[string]interface{}
				err = json.Unmarshal(msg.Value, &m1)
				if err != nil {
					fmt.Printf("unmarshal msg failed,err:%v", err)
					continue
				}
				es.PutLogData(m1)
			}
		}(pc)
	}
	return
}
