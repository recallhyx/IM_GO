package main

import (
	"fmt"
	"log"

	pb "../example/protoc"
	"github.com/gogo/protobuf/proto"
)

func Protoc() {
	p := &pb.Frame{
		ProtoSign:  1234,
		MsgLength:  1,
		MsgType:    1,
		SenderTime: 100000,
		Src: &pb.User{
			UserName: "lzy",
			UserPwd:  "123",
		},
	}
	// 编码
	out, err := proto.Marshal(p)
	if err != nil {
		log.Fatal("failed to marshal: ", err)
	}
	// 解码
	p2 := &pb.Frame{}
	if err := proto.Unmarshal(out, p2); err != nil {
		log.Fatal("failed to unmarshal: ", err)
	}
	fmt.Println(p2)
	fmt.Println(out)
	fmt.Println(string(out))

}
func main() {
	//Protoc()
	H:=0x01
	L:=0x02
	test := (int16(H<<8|L))
	fmt.Println(test)
	/*
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
		// 解码
		p2 := &pb.Person{}
		if err := proto.Unmarshal(out, p2); err != nil {
			log.Fatal("failed to unmarshal: ", err)
		}
		fmt.Println(p2)
		fmt.Println(out)
		fmt.Println(string(out))
	*/
}
