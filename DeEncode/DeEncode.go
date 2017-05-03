package DeEncode

//对接收的信息进行解码并处理
import (
	"fmt"
	"log"

	"net"

	"../Handle"

	pb "../example/protoc"

	"../onLineUsers"

	"github.com/gogo/protobuf/proto"
)

const (
	//==========消息类型=============//
	// 登录
	Login int32 = 0
	// 注册
	Register int32 = 1
	// 反馈
	FeedBack int32 = 2
	// 删除好友
	DelFriend int32 = 3
	//聊天信息
	ChatMsg int32 = 4

	//==========登录返回值=============//

	//登录成功
	LoginSuccess int32 = 200
	//登录失败
	LoginFailed int32 = 201
)

//编码登录帧
func EncodeLoginProtoc(msgType int32, userName, userPwd string,userID int32) ([]byte, error) {
	p := &pb.Frame{
		ProtoSign:  1234,
		MsgLength:  1,
		MsgType:    msgType,
		SenderTime: 100000,
		Src: &pb.User{
			UserName: userName,
			UserPwd:  userPwd,
			UserID:userID,
		},
	}
	// 编码
	out, err := proto.Marshal(p)
	return out, err
}
//编码聊天信息帧
func EncodeChatMsgProtoc(msgType int32, srcName string,srcID int32,dstName string,dstID int32,msgChat string) ([]byte, error) {
	p := &pb.Frame{
		ProtoSign:  1234,
		MsgLength:  1,
		MsgType:    msgType,
		SenderTime: 100000,
		Src: &pb.User{
			UserName: srcName,
			UserID:srcID,
		},
		Dst: &pb.DstUser{
			[]*pb.User{
				{UserName:dstName,UserID:dstID},
			},
		},
		Msg:&pb.Msg{
			Msg:msgChat,
		},
	}
	// 编码
	out, err := proto.Marshal(p)
	return out, err
}

//编码反馈帧
func EncodeFeedBackProtoc(msgType int32, userName string,
	rslCode int32, actionCode int32) ([]byte, error) {
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
	if Handle.Login(frame.Src) { // 登录成功
		log.Println("login..check.ok")

		//添加到在线用户列表
		onLineUsers.AddConnList(string(frame.Src.UserID),conn)
		onLineUsers.GetOnLineChan() <- conn
		//发送返回帧
		//编码
		data, err := EncodeFeedBackProtoc(FeedBack, "IM", LoginSuccess, Login)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("Send...")
		log.Println(data)
		conn.Write(data)
	} else {
		log.Println("login..check.failed")
		//认定用户下线
		onLineUsers.GetOffLineChan() <- conn
		//从在线用户列表删除
		//onLineUsers.RemoveConnList(123)
		//发送返回帧
		//编码
		data, err := EncodeFeedBackProtoc(FeedBack, "IM", LoginFailed, Login)
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
func handleChatMsg(frame *pb.Frame,conn net.Conn,rawFrame []byte ) {
	log.Println(frame)
	dstUserID := frame.GetDst().GetDst()[0].GetUserID()
	connList :=onLineUsers.GetConnList()
	dstConn,exists := connList[string(dstUserID)]
	if exists{
		log.Print("send to user:")
		log.Println(dstUserID)
		dstConn.Write(rawFrame)
	}else{
		log.Println("user offline")
		//handle feedback
	}
}

//消息分发
func msgMux(frame *pb.Frame,rawFrame []byte ,conn net.Conn) {
	switch msgType := frame.MsgType; msgType {
	case Login:
		handleLogin(frame, conn)
		break
	case ChatMsg:
		handleChatMsg(frame,conn,rawFrame)
		break
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
	go msgMux(frame,request ,clnConn)
	return frame, err
}
