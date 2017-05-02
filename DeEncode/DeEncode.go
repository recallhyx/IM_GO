package DeEncode

import (
	"fmt"
	"log"

	"net"

	"IM_GO/Handle"

	pb "IM_GO/example/protoc"

	"IM_GO/onLineUsers"

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
	//聊天信息
	ChatMsg int32 = 4
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
	log.Println(p)
	// 编码
	out, err := proto.Marshal(p)
	return out, err
}

//解码 内部
func DeCodeProtoc(request []byte, readlen int) (*pb.Frame, error) {
	frame := &pb.Frame{}
	err := proto.Unmarshal(request[:readlen], frame)
	return frame, err
}

//处理登录帧
func handleLogin(frame *pb.Frame, conn net.Conn) {
	if Handle.Login(frame.Src) {
		log.Println("login..check.ok")
		onLineUsers.GetOnLineChan() <- conn
		//发送返回帧
		//编码
		data, err := EncodeFeedBackProtoc(2, "lzy", 0, 0, "login ok")
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("Send...")
		log.Println(data)
		conn.Write(data)
	} else {
		log.Println("login..check.failed")
		onLineUsers.GetOffLineChan() <- conn
		//发送返回帧
		//编码
		data, err := EncodeFeedBackProtoc(2, "lzy", 1, 0, "login failed")
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("Send...")
		log.Println(data)
		conn.Write(data)
	}
}

//处理聊天信息帧
func handleChatMsg(frame *pb.Frame, conn net.Conn) {
	log.Println(frame)
}

//消息分发
func msgMux(frame *pb.Frame, conn net.Conn) {
	switch msgType := frame.MsgType; msgType {
	case Login:
		handleLogin(frame, conn)
		break
	case ChatMsg:
		handleChatMsg(frame, conn)
	default:
	}
}

//对收到的消息进行解码并且分发
func HandleMsg(request []byte, readlen int, clnConn net.Conn) (*pb.Frame, error) {
	//解码得到信息帧
	frame, err := DeCodeProtoc(request, readlen)
	if err != nil {
		log.Fatal("failed to unmarshal: ", err)
	}
	fmt.Println(frame)
	//分发消息
	go msgMux(frame, clnConn)
	return frame, err
}
