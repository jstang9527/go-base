package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 10000,
	})
	if err != nil {
		fmt.Println("创建UDP客户端失败，Info: ", err)
		return
	}
	defer listen.Close()

	//不需要建立连接收发数据
	for {
		var data [1024]byte
		n, addr, err := listen.ReadFromUDP(data[:]) //接收数据
		if err != nil {
			fmt.Println("服务端接收数据出错, Info: ", err)
			continue
		}
		fmt.Printf("data:[%#v] from addr:%v count:%v\n", string(data[:n-1]), addr, n)
		reponseData := strings.ToUpper(string(data[:n]))
		_, err = listen.WriteToUDP([]byte(reponseData), addr) //发送数据
		if err != nil {
			fmt.Println("服务端写数据出错, Info: ", err)
			continue
		}
	}
}
