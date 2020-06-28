package unpack

/*
* 对消息进行编码解码
 */
import (
	"encoding/binary"
	"errors"
	"io"
)

// MsgHeader x
const MsgHeader = "12345678"

// Encode 消息编码, 将字节流消息写入到bytesBuffer中，准备传输
func Encode(bytesBuffer io.Writer, content string) error {
	//msg_header + content_len + content
	// 8 + 4 + 内容长度
	if err := binary.Write(bytesBuffer, binary.BigEndian, []byte(MsgHeader)); err != nil {
		return err
	}
	clen := int32(len([]byte(content))) //32比特4字节
	if err := binary.Write(bytesBuffer, binary.BigEndian, clen); err != nil {
		return err
	}
	if err := binary.Write(bytesBuffer, binary.BigEndian, []byte(content)); err != nil {
		return err
	}
	return nil
}

// Decode 消息解码, 将从系统内核接收到的字节流消息转换成string、图片等消息
func Decode(bytesBuffer io.Reader) (bodyBuf []byte, err error) {
	MagicBuf := make([]byte, len(MsgHeader))                     //读取消息头长度
	if _, err = io.ReadFull(bytesBuffer, MagicBuf); err != nil { //读取字节流的消息头
		return nil, err
	}
	if string(MagicBuf) != MsgHeader { //比较消息头是否和本地消息头相同
		return nil, errors.New("msg_header error")
	}

	lengthBuf := make([]byte, 4)
	if _, err = io.ReadFull(bytesBuffer, lengthBuf); err != nil { //读取字节流的消息长度
		return nil, err
	}

	length := binary.BigEndian.Uint32(lengthBuf) //将字节流的长度解码为整型
	//根据消息内容长度读取对应长度的消息
	bodyBuf = make([]byte, length)
	if _, err = io.ReadFull(bytesBuffer, bodyBuf); err != nil {
		return nil, err
	}
	return bodyBuf, err
}
