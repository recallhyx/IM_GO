package main

import (
	"fmt"
	"net"
	"os"
	//"strconv"
	//"strings"
	"time"

	"../database"
)

var clnOnLineChannel chan net.Conn
var clnOffLineChannel chan net.Conn

func clnMgr() {
	//通道初始化
	clnOnLineChannel = make(chan net.Conn)
	clnOffLineChannel = make(chan net.Conn)

	for {
		select {
		//用户上线
		case clnConn := <-clnOffLineChannel:
			{
				fmt.Println(clnConn.RemoteAddr().String() + "exit")
				clnConn.Close()
				break
			}

		}

	}
}
func main() {
	//初始化数据库
	database.SetupDB()

	service := ":1200"
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
		if err != nil {
			continue
		}
		go handleClient(conn)
	}
}

//处理每一个客户端
func handleClient(conn net.Conn) {
	conn.SetReadDeadline(time.Now().Add(2 * time.Minute)) // set 2 minutes timeout
	request := make([]byte, 128)                          // set maxium request length to 128B to prevent flood attack
	defer conn.Close()                                    // close connection before exit
	for {
		read_len, err := conn.Read(request)
		if err != nil {
			clnOffLineChannel <- conn
			//fmt.Println(err)
			break
		}
		fmt.Println(string(request[:read_len]))
		if read_len == 0 {
			clnOffLineChannel <- conn
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
