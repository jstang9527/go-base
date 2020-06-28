package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/nsqio/go-nsq"
)

//Myhandler 消费者类型
type Myhandler struct {
	Titile string
}

//HandleMessage 需要实现处理消息的方法
func (m *Myhandler) HandleMessage(msg *nsq.Message) (err error) {
	fmt.Printf("%s recv from %v, msg:%v\n", m.Titile, msg.NSQDAddress, string(msg.Body))
	return
}

func initConsumer(topic string, channel string, address string) (err error) {
	config := nsq.NewConfig()
	config.LookupdPollInterval = time.Second * 15
	c, err := nsq.NewConsumer(topic, channel, config)
	if err != nil {
		fmt.Println("create consumer failed, err:", err)
		return
	}
	consumer := &Myhandler{Titile: "沙河2号"}
	c.AddHandler(consumer)
	// if err := c.ConnectToNSQD(address); err != nil{  //直接连NSQD
	if err := c.ConnectToNSQLookupd(address); err != nil { //通过lookupd查询
		return err
	}
	return nil
}

func main() {
	err := initConsumer("topic_demo", "firest", "127.0.0.1:4161")
	if err != nil {
		fmt.Println("init consumer failed, err", err)
		return
	}
	c := make(chan os.Signal)        //定义一个信号通道
	signal.Notify(c, syscall.SIGINT) //转发键盘的信号给C  (ctrl+c)
	<-c                              //阻塞
	fmt.Println("\n退出")
}
