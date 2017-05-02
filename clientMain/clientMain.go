// clientMain
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"

	"../DeEncode"
)

func SendMsg(conn net.Conn) {
	reader := bufio.NewReader(os.Stdin)
	msg, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	_, err = conn.Write([]byte(msg))
	if err != nil {
		fmt.Println(err)
		return
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
	fmt.Println(string(returnBuf))
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
	conn, _ := net.Dial("tcp4", "192.168.1.115:6666")
	log.Println("connect....")
	loginData, err := DeEncode.EncodeLoginProtoc(DeEncode.Login, "lzy", "123")
	if err != nil {
		log.Fatalln(err)
	}
	//发送一条登录帧
	_, err = conn.Write(loginData)
	if err != nil {
		log.Fatalln(err)
		return
	}
	log.Println("Send ok....")
	defer conn.Close()
	for {

		recv(conn)
		//SendMsg(conn)
		//	go recv(conn)
		// listen for reply
		//message, _ := bufio.NewReader(conn).ReadString('\n')
		//fmt.Print("Message from server: " + message)
	}
}
