package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(os.Stdin)
	for {
		var tmp [128]byte
		n, err := conn.Read(tmp[:])
		if err == io.EOF {
			fmt.Println("read from conn finished, Info: ", err)
			break
		}
		if err != nil {
			fmt.Println("read from conn failed, err: ", err)
			break
		}
		fmt.Println("收到客户端消息: ", string(tmp[:n]))

		fmt.Print("请回复: ")
		response, _ := reader.ReadString('\n')
		response = strings.TrimSpace(response)
		if _, err := conn.Write([]byte(response)); err != nil {
			fmt.Println("服务端回复对方时出错了,err:", err)
			break
		}
	}

}

func main() {
	//1.监听端口
	listener, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("start tcp server on 127.0.0.1:8888 failed, err: ", err)
		return
	}
	//2.等待他人连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept failed, err: ", err)
			return
		}
		//3.与客户通信
		go process(conn)
	}

}
