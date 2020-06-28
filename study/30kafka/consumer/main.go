package main

import (
	"fmt"
	"time"

	"github.com/Shopify/sarama"
)

func main() {
	consumer, err := sarama.NewConsumer([]string{"192.168.231.128:9092"}, nil)
	if err != nil {
		fmt.Println("fail to start consumer, err: ", err)
		return
	}
	fmt.Println("create consumer success.")
	partitonList, err := consumer.Partitions("web_log") //根据topic取到所有分区
	if err != nil {
		fmt.Println("fail to get list of partition, err: ", err)
		return
	}
	fmt.Println("partition list: ", partitonList)
	for partition := range partitonList { //遍历所有分区
		//针对每个分区创建一个对应的分区消费者
		pc, err := consumer.ConsumePartition("web_log", int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Println("failed to start consumer for partition ", partition, "err: ", err)
			return
		}
		defer pc.AsyncClose()
		// 异步消费信息
		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d Offset:%d Key:%v Value:%s\n", msg.Partition, msg.Offset, msg.Key, string(msg.Value))
			}
		}(pc)
	}
	time.Sleep(time.Second * 60) //这个是避免异步消费信息被退出，接收60秒内的产生的数据
}
