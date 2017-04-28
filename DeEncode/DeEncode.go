package DeEncode

import (
	"fmt"
	"log"

	pb "../example/protoc"
	"github.com/gogo/protobuf/proto"
)

//编码
func encodeProtoc() []byte {
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
	return out
}

//解码
func deCodeProtoc(request []byte, readlen int) (*pb.Frame, error) {
	p2 := &pb.Frame{}
	err := proto.Unmarshal(request[:readlen], p2)
	if err != nil {
		log.Fatal("failed to unmarshal: ", err)
	}
	fmt.Println(p2)
	return p2, err
}
