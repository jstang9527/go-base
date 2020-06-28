package main

import (
	"fmt"

	"github.com/jstang007/gateway_demo/httpserver/server"
)

// TCPTest X
func TCPTest() {
	//启动服务器
	server.SrvForever()

	//启动客户端
	// for i := 0; i < 5; i++ {
	// 	time.Sleep(time.Second * 2)
	// 	fmt.Printf("-------启动客户端[%v]--------\n", i)
	// 	client.TCPClient()
	// }
}

// UDPTest X
func UDPTest() {
	//启动服务端
	server.UDPServer()
	//启动客户端
	// time.Sleep(time.Second * 2)
	// client.UDPClient()
}

type MyFunc func(i string, j string) //函数类型

func (m MyFunc) SrvAdd(w string, r string) { //函数类型的实现体的方法
	m(w, r)
}

func myFunc(res string, req string) { //函数类型的实现
	fmt.Println(res + req)
}

func main() {
	// TCPTest()
	// UDPTest()
	hf := MyFunc(myFunc)     //获取函数类型实现体
	hf.SrvAdd("resp", "req") //调用函数类型的方法
}
