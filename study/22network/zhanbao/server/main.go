package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	var buf [1024]byte
	for {
		n, err := reader.Read(buf[:])
		if err == io.EOF {
			fmt.Println("客户端数据读取完毕.")
			break
		}
		if err != nil {
			fmt.Println("读取客户端数据失败, err:", err)
			break
		}
		fmt.Println("收到客户端数据:", string(buf[:n]))
	}
}

func main() {
	listenner, err := net.Listen("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println("创建服务端失败, Info: ", err)
		return
	}
	defer listenner.Close()
	for {
		conn, err := listenner.Accept()
		if err != nil {
			fmt.Println("监听请求发生错误, Info: ", err)
			continue
		}
		go process(conn)
	}
}
