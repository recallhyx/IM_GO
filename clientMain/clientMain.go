// clientMain
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"../DeEncode"
	"io/ioutil"
)

//发送一帧登录请求帧
func sendLoginFrame(conn net.Conn){
	loginData, err := DeEncode.EncodeLoginProtoc(DeEncode.Login, "lzy", "123",2)
	if err != nil {
		log.Fatalln(err)
	}
	//发送一条登录帧
	_, err = conn.Write(loginData)
	if err != nil {
		log.Fatalln(err)
		return
	}
	log.Println("Send LoginFrame ok....")
}

//发送一帧注册请求帧
func sendRegisterFrame(conn net.Conn) {
	registerData, err := DeEncode.EncodeRegisterProtoc(DeEncode.Register, "13022015751", "123456789")
	if err != nil {
		log.Fatalln(err)
	}
	//发送一条登录帧
	_, err = conn.Write(registerData)
	if err != nil {
		log.Fatalln(err)
		return
	}
	log.Println("Send RegisterFrame ok....")
}

//发送聊天消息帧
func sendChatMsgFrame(conn net.Conn){
	chatMsgData, err := DeEncode.EncodeChatMsgProtoc(DeEncode.ChatMsg, "lzy", 2,"wzb",1,"hello from go client")
	if err != nil {
		log.Fatalln(err)
	}
	//发送一条聊天消息帧
	_, err = conn.Write(chatMsgData)
	if err != nil {
		log.Fatalln(err)
		return
	}
	log.Println("Send ChatMsgFrame ok....")
}
//读文件
func readFile(filePth string) ([]byte, error) {
	f, err := os.Open(filePth)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(f)
}
//发送文件帧
func sendFileMsg(conn net.Conn){
	fileByte,err := readFile("./TestSendFile/a.txt")
	if err != nil{
		log.Println(err)
		return
	}
	fileMsg:=DeEncode.EncodeFileMsg(DeEncode.FileMsg, 1, 2, fileByte)
	//发送一条文件消息帧
	_, err = conn.Write(fileMsg)
	if err != nil {
		log.Fatalln(err)
		return
	}
	log.Println("Send fileMsg ok....")
}
func ReadCmd(conn net.Conn) {
	reader := bufio.NewReader(os.Stdin)
	msg, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if strings.Contains(msg,"login"){//输入的字符串包含login
		sendLoginFrame(conn)
	}
	if strings.Contains(msg,"sendMsg"){//输入的字符串包含sendMsg
		sendChatMsgFrame(conn)
	}
	if strings.Contains(msg,"register"){//输入的字符串包含register
		sendRegisterFrame(conn)
	}
	if strings.Contains(msg,"sendFile"){//输入的字符串包含sendFile
		sendFileMsg(conn)
	}
}
func recv(conn net.Conn) {
	returnBuf := make([]byte, 128)
	readlen, err := conn.Read(returnBuf)

	if err != nil {
		log.Fatalln(err)
	}
	//解码
	frame, err := DeEncode.DeCodeProtoc(returnBuf, readlen)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("GET...")
	//fmt.Println(string(returnBuf))
	fmt.Println(frame)
}

/*
func ProtoBufMsg() []byte {
	p := &pb.Frame{
		Id:    1234,
		Name:  "Jerry Hou",
		Email: "https@yryz.net",
		Phones: []*pb.Person_PhoneNumber{
			{Number: "110", Type: pb.Person_HOME},
			{Number: "911", Type: pb.Person_WORK},
		},
	}
	// 编码
	out, err := proto.Marshal(p)
	if err != nil {
		log.Fatal("failed to marshal: ", err)
	}
	return out
}*/

func main() {

	// connect to this socket
	conn, _ := net.Dial("tcp4", "192.168.191.1:6666")
	log.Println("connect....")

	defer conn.Close()
	for {
		ReadCmd(conn)
		go recv(conn)
		//SendMsg(conn)
		//	go recv(conn)
		// listen for reply
		//message, _ := bufio.NewReader(conn).ReadString('\n')
		//fmt.Print("Message from server: " + message)
	}
}
