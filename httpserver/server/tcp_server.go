package server

import (
	"fmt"
	"net"

	"github.com/jstang007/gateway_demo/httpserver/handler"
)

// SrvForever X
func SrvForever() {
	//1、创建套接字
	listenner, err := net.Listen("tcp", "0.0.0.0:9091")
	if err != nil {
		fmt.Printf("listen fail, err:%v\n", err)
		return
	}
	//2、监听套接字
	for {
		conn, err := listenner.Accept() //接收请求, 阻塞态
		if err != nil {
			fmt.Printf("accept fail, err:%v\n", err)
			continue
		}
		//3.1 处理请求
		// go handler.Process(conn)
		//3.2 读客户端发过来的消息
		go handler.Read(conn)
	}
}
