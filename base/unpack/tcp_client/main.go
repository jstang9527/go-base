package main

import (
	"fmt"
	"net"

	"github.com/jstang007/gateway_demo/base/unpack/unpack"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:9091") //连接服务器,获取套接字conn
	defer conn.Close()                             //在main函数执行完毕即将退出之前，关闭连接
	if err != nil {
		fmt.Printf("connect failed, err : %v\n", err.Error())
		return
	}
	unpack.Encode(conn, "hello world 0!!!")
}
