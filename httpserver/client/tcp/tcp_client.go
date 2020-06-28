package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

//TCPClient 客户端
func TCPClient() {
	//1.连接服务器
	conn, err := net.Dial("tcp", "localhost:9091") //创建套接字
	defer conn.Close()
	if err != nil {
		fmt.Printf("connect failed, error:%v", err)
	}
	//2.读取命令行输入
	inputReader := bufio.NewReader(os.Stdin)
	// handler.Send(conn, "my name is TCP Client!")
	for {
		//3.一直读取，直到读到\n
		input, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Printf("read from console failed, err:%v", err)
			break
		}
		//4.读到一行只有一个Q时停止
		trimmedInput := strings.TrimSpace(input)
		if trimmedInput == "Q" {
			break
		}
		//5.回复服务器消息
		_, err = conn.Write([]byte(trimmedInput))
		if err != nil {
			fmt.Printf("write failed, err:%v", err)
			break
		}
	}
}

func main() {
	TCPClient()
}
