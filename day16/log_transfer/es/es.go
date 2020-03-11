package es

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
)

//ES
//将日志数据写入Elasticsearch

type ESClient struct {
	client      *elastic.Client  //es 链接
	index       string           //index(数据库)
	logDataChan chan interface{} //接收日志数据的channel  *sarama.ConsumerMessage
}

var (
	esClient *ESClient
)

func Init(addr, index string, goroutineNum int, maxSize int) (err error) {
	client, err := elastic.NewClient(elastic.SetURL("http://" + addr))
	if err != nil {
		// Handle error
		panic(err)
	}
	esClient = &ESClient{
		client:      client,
		index:       index,
		logDataChan: make(chan interface{}, maxSize),
	}
	fmt.Println("connect to es success")
	//从通道中取出数据,写入到kafka里面去
	for i := 0; i < goroutineNum; i++ {
		go sendToES()
	}
	return
}

func sendToES() {
	for m1 := range esClient.logDataChan {
		put1, err := esClient.client.Index().
			Index(esClient.index).
			BodyJson(m1).
			Do(context.Background())
		if err != nil {
			// Handle error
			panic(err)
		}
		fmt.Printf("Indexed user %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
	}
}

//通过一个首字母大写的函数,调用方从包外接收msg,发送到chan中
func PutLogData(msg interface{}) {
	esClient.logDataChan <- msg
}
