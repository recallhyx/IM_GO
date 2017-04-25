package main

import (
	"fmt"
	"log"

	pb "../example"

	"github.com/gogo/protobuf/proto"
)

func main() {
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
}
