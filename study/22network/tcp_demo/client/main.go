package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	//1.连接server端
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("dial 127.0.0.1:8888 failed, err: ", err)
		return
	}
	defer conn.Close()
	//2.发送数据
	reader := bufio.NewReader(os.Stdin) //从标志输入生成对象
	for {
		fmt.Print("请输入内容: ")
		text, err := reader.ReadString('\n') //读到换行
		if err != nil {
			fmt.Println("空行退出, Info:", err)
			break
		}
		text = strings.TrimSpace(text)
		if len(text) == 0 {
			fmt.Println("输入空行，退出程序")
			break
		}
		conn.Write([]byte(text))

		//接收客户端数据
		buf := [512]byte{}
		if n, err := conn.Read(buf[:]); err != nil {
			fmt.Println("读取服务端回信失败. err:", err)
			break
		} else {
			fmt.Println("收到服务端回信:", string(buf[:n]))
		}

	}
}
