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

//编码登录帧
func EncodeLoginProtoc(msgType int32, userName, userPwd string) ([]byte, error) {
	p := &pb.Frame{
		ProtoSign:  1234,
		MsgLength:  1,
		MsgType:    msgType,
		SenderTime: 100000,
		Src: &pb.User{
			UserName: userName,
			UserPwd:  userPwd,
		},
	}
	// 编码
	out, err := proto.Marshal(p)
	return out, err
}

//编码反馈帧
func EncodeFeedBackProtoc(msgType int32, userName string,
	rslCode int32, actionCode int32, rslMsg string) ([]byte, error) {
	p := &pb.Frame{
		ProtoSign:  1234,
		MsgLength:  1,
		MsgType:    msgType,
		SenderTime: 100000,
		Src: &pb.User{
			UserName: userName,
		},
		FbAction: &pb.Action{
			RslCode:    rslCode,
			RslMsg:     rslMsg,
			ActionType: actionCode,
		},
	}
	// 编码
	out, err := proto.Marshal(p)
	return out, err
}

//解码 内部
func deCodeProtoc(request []byte, readlen int) (*pb.Frame, error) {
	frame := &pb.Frame{}
	err := proto.Unmarshal(request[:readlen], frame)
	return frame, err
}

//消息分发
func msgMux(frame *pb.Frame) {
	switch msgType := frame.MsgType; msgType {
	case Login:
		Handle.HandleLogin(frame.Src)
	default:
	}
}

//对收到的消息进行解码并且分发
func HandleMsg(request []byte, readlen int) (*pb.Frame, error) {
	//解码得到信息帧
	frame, err := deCodeProtoc(request, readlen)
	if err != nil {
		log.Fatal("failed to unmarshal: ", err)
	}
	fmt.Println(frame)
	//分发消息
	go msgMux(frame)
	return frame, err
}
