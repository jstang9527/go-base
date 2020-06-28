package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 10000,
	})
	if err != nil {
		fmt.Println("连接UDP服务器出错, Info: ", err)
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)
	var response [1024]byte
	for {
		fmt.Print("请输入内容: ")
		msg, _ := reader.ReadString('\n')
		_, err := conn.Write([]byte(msg))
		if err != nil {
			fmt.Println("发送数据失败了, Info: ", err)
			continue
		}

		//收数据
		n, addr, err := conn.ReadFromUDP(response[:])
		if err != nil {
			fmt.Println("收数据发生错误, err: ", err)
		}
		fmt.Printf("从地址[%v]接收到数据(%#v): %v", addr, n, string(response[:n]))
	}
}
