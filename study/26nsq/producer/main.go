package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/nsqio/go-nsq"
)

var producer *nsq.Producer

//初始化生产者
func initPorducer(str string) (err error) {
	config := nsq.NewConfig()
	producer, err = nsq.NewProducer(str, config)
	if err != nil {
		fmt.Println("create producer failed, err: ", err)
		return
	}
	return nil
}

func main() {
	nsqAddr := "127.0.0.1:4150"
	err := initPorducer(nsqAddr)
	if err != nil {
		fmt.Println("init producer failed, err: ", err)
		return
	}
	reader := bufio.NewReader(os.Stdin)
	for {
		data, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("read string from stdin failed, err:", err)
			continue
		}
		data = strings.TrimSpace(data)
		if strings.ToUpper(data) == "Q" { //输入Q/qQ退出
			break
		}
		if err = producer.Publish("topic_demo", []byte(data)); err != nil {
			fmt.Println("publish msg to nsq failed, err: ", err)
			continue
		}
	}

}
