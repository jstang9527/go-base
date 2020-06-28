package handler

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
)

//MsgHeader 消息头
var MsgHeader = "12345678"

// Decode 解码，将字节流解码成utf-8
func decode(bytesBuffer io.Reader) (bodyBuf []byte, err error) {
	//1.首先读取消息头数据存到MagicBuf
	MagicBuf := make([]byte, len(MsgHeader))
	if _, err = io.ReadFull(bytesBuffer, MagicBuf); err != nil {
		return nil, err
	}

	//2、读取消息内容长度到lengthBuf
	lengthBuf := make([]byte, 4)
	if _, err := io.ReadFull(bytesBuffer, lengthBuf); err != nil {
		return nil, err
	}

	//3.读取消息内容到bodyBuf
	length := binary.BigEndian.Uint32(lengthBuf) //将二进制转进行解码
	bodyBuf = make([]byte, length)
	if _, err := io.ReadFull(bytesBuffer, bodyBuf); err != nil {
		return nil, err
	}
	return bodyBuf, err
}

// Process 处理请求
func Process(conn net.Conn) {
	defer conn.Close()
	for {
		bt, err := decode(conn)
		if err != nil {
			fmt.Printf("对端套接字关闭, 我方相关套接字也将close, err: %v\n", err)
			break
		}
		str := string(bt)
		fmt.Printf("接收到客户端数据: %s\n", str)
	}
}

// Read 读消息
func Read(conn net.Conn) {
	defer conn.Close()
	for {
		var buf [128]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Printf("对端套接字关闭, 我方相关套接字也将close, err: %v\n", err)
			break
		}
		str := string(buf[:n])
		fmt.Printf("接收到客户端数据: %v\n", str)
	}
}

// Send 发送消息
func Send(conn net.Conn, content string) error {
	return encode(conn, content)
}

// encode 将string字符串转换成字节流
func encode(bytesBuffer io.Writer, contetnt string) error {
	//1.写入消息头
	bytesMsgHeader := []byte(MsgHeader)
	if err := binary.Write(bytesBuffer, binary.BigEndian, bytesMsgHeader); err != nil {
		return err
	}
	//2.写入内容长度
	clen := int32(len([]byte(contetnt)))
	if err := binary.Write(bytesBuffer, binary.BigEndian, clen); err != nil {
		return err
	}
	//3.写入内容
	bytesContent := []byte(contetnt)
	if err := binary.Write(bytesBuffer, binary.BigEndian, bytesContent); err != nil {
		return err
	}
	return nil
}
