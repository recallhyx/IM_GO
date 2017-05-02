package onLineUsers

import (
	"fmt"
	"net"
	"strconv"
)

var clnOnLineChannel chan net.Conn

var clnOffLineChannel chan net.Conn

func InitNetChannel() {
	clnOnLineChannel = make(chan net.Conn)  //客户在线通道
	clnOffLineChannel = make(chan net.Conn) //客户下线通道
}

func GetOnLineChan() chan net.Conn {
	return clnOnLineChannel
}
func GetOffLineChan() chan net.Conn {
	return clnOffLineChannel
}

//统计在线人数，将上线用户存入map里
func ShowOnLines(arg_conns map[string]net.Conn) {

	fmt.Println("Online Number: " + strconv.Itoa(len(arg_conns))) //strconv.Itoa将整数转换为十进制字符串形式
	//遍历在线用户
	for key := range arg_conns {
		fmt.Println("目前在线的用户：", key)
	}
}
