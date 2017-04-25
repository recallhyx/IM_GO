// clientMain
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
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
func main() {

	// connect to this socket
	conn, _ := net.Dial("tcp4", "127.0.0.1:1200")
	defer conn.Close()
	for {

		SendMsg(conn)
		go recv(conn)

		// listen for reply
		//message, _ := bufio.NewReader(conn).ReadString('\n')
		//fmt.Print("Message from server: " + message)
	}
}
