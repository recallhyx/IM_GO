// clientMain
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"

	pb "../example"
	"github.com/gogo/protobuf/proto"
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
	_, err := conn.Read(returnBuf)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("GET...")
	fmt.Println(string(returnBuf))
}
func ProtoBufMsg() []byte {
	p := &pb.Person{
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
}

func main() {

	// connect to this socket
	conn, _ := net.Dial("tcp4", "192.168.191.1:6666")
	log.Println("connect....")
	_, err := conn.Write(ProtoBufMsg())
	if err != nil {
		fmt.Println(err)
		return
	}
	log.Println("Send ok....")
	defer conn.Close()
	for {

		//SendMsg(conn)
		//	go recv(conn)
		// listen for reply
		//message, _ := bufio.NewReader(conn).ReadString('\n')
		//fmt.Print("Message from server: " + message)
	}
}
