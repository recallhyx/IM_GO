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

//编码反馈帧
func Test_EncodingFbFrame(t *testing.T) {
	log.Println("test3")
	//编码
	data, err := EncodeFeedBackProtoc(1, "lzy", 1, 1, "safe")
	if err != nil {
		t.Error(err)
	}

	//处理数据
	frame, err := HandleMsg(data, len(data))
	if err != nil {
		t.Error(err)
	}
	log.Println(frame)

}

func Test_HandleMsg(t *testing.T) {
	log.Println("test1")
	//准备编码好的二进制数据
	data, err := EncodeLoginProtoc(1, "lzy", "abc")
	if err != nil {
		t.Error(err)
	}
	//处理数据
	_, err = HandleMsg(data, len(data))
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
	msgMux(frame)

}
