package onLineUsers

import (
	"fmt"
	"net"
	"strconv"
)

var clnOnLineChannel chan net.Conn

var clnOffLineChannel chan net.Conn

//string为用户id，Conn为连接
var connList map[string]net.Conn

func InitNetChannel() {
	clnOnLineChannel = make(chan net.Conn)  //客户在线通道
	clnOffLineChannel = make(chan net.Conn) //客户下线通道
	connList = make(map[string]net.Conn)    //在线用户列表
}

//获取上线通道
func GetOnLineChan() chan net.Conn {
	return clnOnLineChannel
}

//获取离线通道
func GetOffLineChan() chan net.Conn {
	return clnOffLineChannel
}

//获取在线用户列表
func GetConnList() map[string]net.Conn {
	return connList
}

func AddConnList(userID string,clnConn net.Conn){
	connList[userID]=clnConn
}
func RemoveConnList(userID string){
	delete(connList, userID)
}

//统计在线人数，将上线用户存入map里
func ShowOnLines() {

	fmt.Println("Online Number: " + strconv.Itoa(len(connList))) //strconv.Itoa将整数转换为十进制字符串形式
	//遍历在线用户
	for key := range connList {
		fmt.Println("目前在线的用户：", key)
	}
}
