// DeEncode_test
package DeEncode

import (
	"testing"

	"log"

	"../database"
	pb "../example/protoc"
)

func init() {
	log.Println("init db")
	database.SetupDB()
}
func Test_Decoding(t *testing.T) {
	log.Println("test1")
	//准备编码好的二进制数据
	data := encodeProtoc()
	//处理数据
	_, err := Decode(data, len(data))
	if err != nil {
		t.Error(err)
	}
}

func Test_MsgMux(t *testing.T) {
	log.Println("test2")
	frame := &pb.Frame{
		ProtoSign:  1234,
		MsgLength:  1,
		MsgType:    0,
		SenderTime: 100000,
		Src: &pb.User{
			UserName: "lzy",
			UserPwd:  "123",
		},
	}
	//分发消息
	MsgMux(frame)

}
