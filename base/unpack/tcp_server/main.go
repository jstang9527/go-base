package main

import (
	"fmt"
	"net"

	"github.com/jstang007/gateway_demo/base/unpack/unpack"
)

func main() {
	//1、创建套接字
	listener, err := net.Listen("tcp", "0.0.0.0:9090")
	if err != nil {
		fmt.Printf("listen fail, err: %v\n", err)
		return
	}
	//2、监听套接字
	for {
		conn, err := listener.Accept() //阻塞态
		if err != nil {
			fmt.Printf("accept fail, err: %v\n", err)
			continue
		}
		// 3、处理请求
		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close() //处理完请求函数后执行关闭连接
	for {
		bt, err := unpack.Decode(conn)
		if err != nil {
			fmt.Printf("read from connect failed, err: %v", err)
			break
		}
		str := string(bt)
		fmt.Printf("receive from client, data: %v\n", str)
	}
}
