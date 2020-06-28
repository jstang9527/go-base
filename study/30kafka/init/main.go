package main

import (
	"fmt"

	"github.com/Shopify/sarama"
)

//基于sarama第三方库开发的kafka client
func main() {
	config := sarama.NewConfig()
	//tailf包使用
	config.Producer.RequiredAcks = sarama.WaitForAll          //发生哇数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner //新选出一个partition
	config.Producer.Return.Successes = true                   //成功交付的消息将在success channel返回
	//构造一个消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = "web_log"
	msg.Value = sarama.StringEncoder("this is a test log")
	//连接kafka
	client, err := sarama.NewSyncProducer([]string{"192.168.231.128:9092"}, config)
	if err != nil {
		fmt.Println("producer closed, err: ", err)
		return
	}
	defer client.Close()
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send msg failed, err: ", err)
		return
	}
	fmt.Printf("pid:%v offset:%v\n", pid, offset)
}
