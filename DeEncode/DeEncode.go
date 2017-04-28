package DeEncode

import (
	"fmt"
	"log"

	"../Handle"
	pb "../example/protoc"
	"github.com/gogo/protobuf/proto"
)

const (
	// 登录
	Login int32 = 0
	// 注册
	Register int32 = 1
	// 添加好友
	AddFriend int32 = 2
	// 删除好友
	DelFriend int32 = 3
	//发送信息
	SendMsg int32 = 4
)

//编码 内部
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

//解码 内部
func deCodeProtoc(request []byte, readlen int) (*pb.Frame, error) {
	frame := &pb.Frame{}
	err := proto.Unmarshal(request[:readlen], frame)
	if err != nil {
		log.Fatal("failed to unmarshal: ", err)
	}
	fmt.Println(frame)
	return frame, err
}

//消息分发
func MsgMux(frame *pb.Frame) {
	switch msgType := frame.MsgType; msgType {
	case Login:
		Handle.HandleLogin(frame.Src)
	default:
	}
}

//Decode 解码 外部接口
func Decode(request []byte, readlen int) (*pb.Frame, error) {
	//解码得到信息帧
	frame := &pb.Frame{}
	err := proto.Unmarshal(request[:readlen], frame)
	if err != nil {
		log.Fatal("failed to unmarshal: ", err)
	}
	fmt.Println(frame)
	//分发消息
	go MsgMux(frame)
	return frame, err
}
