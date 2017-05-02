package main

import (
	"fmt"
	"log"
	"net"
	"os"

	//"strconv"
	//"strings"
	"time"

	"../DeEncode"
	"../database"
	"../onLineUsers"
)

//此处的string到时候可以存放用户名，key存放的是ip
var connList map[string]string

func clnMgr() {
	onLineUsers.InitNetChannel()
	clnOffLineChannel := onLineUsers.GetOffLineChan()
	clnOnLineChannel := onLineUsers.GetOnLineChan()
	//存储在线用户
	connList := make(map[string]net.Conn)
	for {
		select {
		//用户下线处理统计
		case clnConn := <-clnOffLineChannel:
			{
				//fmt.Println(clnConn.RemoteAddr().String() + "exit")
				clnSap := clnConn.RemoteAddr().String()
				fmt.Println(clnSap + " offline")
				delete(connList, clnSap)
				clnConn.Close()
				onLineUsers.ShowOnLines(connList)
			}
		//用户上线处理统计
		case clnConn := <-clnOnLineChannel:
			{
				clnSap := clnConn.RemoteAddr().String()
				fmt.Println(clnSap + " online")
				connList[clnSap] = clnConn
				onLineUsers.ShowOnLines(connList)
			}
		}

	}
}

func main() {
	//初始化数据库
	database.SetupDB()

	service := ":6666"
	//以ipv4处理
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	//启动监听
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	//处理客户端的业务功能
	go clnMgr()

	fmt.Println("Listening...")
	for {
		//接受客户端连接请求
		conn, err := listener.Accept()
		fmt.Println("accept")
		if err != nil {
			continue
		}
		//conn.Write(encodeProtoc())
		go handleClient(conn)
	}
}

//处理每一个客户端
func handleClient(conn net.Conn) {
	conn.SetReadDeadline(time.Now().Add(2 * time.Minute)) // set 2 minutes timeout
	request := make([]byte, 128)                          // set maxium request length to 128B to prevent flood attack
	defer conn.Close()                                    // close connection before exit
	for {
		readlen, err := conn.Read(request)
		if err != nil {
			onLineUsers.GetOffLineChan() <- conn
			//fmt.Println(err)
			break
		}
		//原来的输入接受
		//fmt.Println(string(request[:read_len]))
		// 处理消息
		_, err = DeEncode.HandleMsg(request[:readlen], readlen, conn)
		//若用户正常上线，则将此用户的conn传进用户上线通道
		//clnOnLineChannel <- conn
		if err != nil {
			log.Println(err)
			continue
		}
		//deCodeProtoc(request, readlen)

		if readlen == 0 {
			onLineUsers.GetOffLineChan() <- conn
			break
		}
		request = make([]byte, 128) // clear last read content
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
