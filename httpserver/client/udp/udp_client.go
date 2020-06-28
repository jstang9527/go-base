package main

import (
	"fmt"
	"net"
)

// UDPClient X
func UDPClient() {
	//1.连接服务器
	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 9092,
	})
	if err != nil {
		fmt.Printf("connect failed, err:%v\n", err)
		return
	}
	for i := 0; i < 5; i++ {
		//2.发送请求
		_, err = conn.Write([]byte("hello server!"))
		if err != nil {
			fmt.Printf("send data failed, err:%v\n", err)
			return
		}
		//3.接收请求
		result := make([]byte, 1024)
		n, remoteAddr, err := conn.ReadFromUDP(result)
		if err != nil {
			fmt.Printf("receive data failed, err:%v\n", err)
			return
		}
		fmt.Printf("receive from addr:%v data:%v\n", remoteAddr, string(result[:n]))
	}
}

func main() {
	UDPClient()
}
