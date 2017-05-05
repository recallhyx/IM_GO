package DeEncode

//对接收的信息进行解码并处理
import (
	"fmt"
	"log"

	"encoding/binary"
	"net"

	"../Handle"

	pb "../example/protoc"

	"../onLineUsers"

	"github.com/golang/protobuf/proto"
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

	//特殊帧
	//特殊帧帧头
	SpFrameHead byte = 0xaa
	//特殊帧长度
	SpFrameLength int32 = 6
	//一般文件消息
	FileMsg int16 = 5

	//==========登录返回值=============//

	//登录成功
	LoginSuccess int32 = 200
	//登录失败
	LoginFailed int32 = 201

	//注册成功
	RegisterSuccess int32 = 200
	//注册失败
	RegisterFail int32 = 201
)

//编码登录帧
func EncodeLoginProtoc(msgType int32, userName, userPwd string, userID int32) ([]byte, error) {
	p := &pb.Frame{
		ProtoSign:  1234,
		MsgLength:  1,
		MsgType:    msgType,
		SenderTime: 100000,
		Src: &pb.User{
			UserName: userName,
			UserPwd:  userPwd,
			UserID:   userID,
		},
	}
	// 编码
	out, err := proto.Marshal(p)
	return out, err
}

// 编码注册帧
func EncodeRegisterProtoc(msgType int32, userName, userPwd string) ([]byte, error) {
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
	out, err := proto.Marshal(p)
	return out, err
}

//编码聊天信息帧
func EncodeChatMsgProtoc(msgType int32, srcName string, srcID int32, dstName string, dstID int32, msgChat string) ([]byte, error) {
	p := &pb.Frame{
		ProtoSign:  1234,
		MsgLength:  1,
		MsgType:    msgType,
		SenderTime: 100000,
		Src: &pb.User{
			UserName: srcName,
			UserID:   srcID,
		},
		Dst: &pb.DstUser{
			[]*pb.User{
				{UserName: dstName, UserID: dstID},
			},
		},
		Msg: &pb.Msg{
			Msg: msgChat,
		},
	}
	// 编码
	out, err := proto.Marshal(p)
	return out, err
}

//编码文件帧
func EncodeFileMsg(msgType int16, srcID int16,
	dstID int16, fileData []byte) []byte {
	var srcIDByte = make([]byte, 2)
	var dstIDByte = make([]byte, 2)
	//go的int为小端方式存数据
	binary.LittleEndian.PutUint16(srcIDByte,  uint16(srcID))
	binary.LittleEndian.PutUint16(dstIDByte, uint16(dstID))
	var file = make([]byte,len(fileData)+int(SpFrameLength))

	head := []byte{SpFrameHead, (byte(msgType)), srcIDByte[0], srcIDByte[1], dstIDByte[0], dstIDByte[1]}
	//合并两个字节数组
	copy(file,head)
	for i:=0;i<len(fileData);i++{
		file = append(file,fileData[i])
	}
	return file
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

//字节数组转int16(注：由于java int用大端方式，go用小端方式，所以需要逆序)
func byteArrayToint16(high, low byte) int16 {
	return (int16(high<<8 | low))
}

//处理特殊帧
func handleSPMsg(request []byte, readlen int, conn net.Conn) {
	//todo 传空文件服务器会crash
	//把帧头读出
	head := request[:SpFrameLength]
	log.Println(head)
	log.Println(request)
	log.Println(readlen)
	file := request[SpFrameLength:readlen]
	log.Println(file)
	//src:=byteArrayToint16(request[3],request[2])
	dstID := byteArrayToint16(request[5], request[4])
	switch msgType := (int16(head[1])); msgType {
	case FileMsg: //一般文件信息
		connList := onLineUsers.GetConnList()
		dstConn, exists := connList[string(dstID)]
		if exists { //如果用户在线，直接转发
			log.Print("send file to user:")
			log.Println(dstID)
			dstConn.Write(request[:readlen])
		} else { //用户不在线，先保存在服务器
			log.Println("user offline save file")
			log.Println(dstID)
			err := Handle.Save(dstID, file)
			if err != nil {
				log.Println(err)
			}
			//handle feedback
		}
		break
	default:
	}
}

//处理登录帧
func handleLogin(frame *pb.Frame, conn net.Conn) {
	if Handle.Login(frame.Src) { // 登录成功
		log.Println("login..check.ok")

		//添加到在线用户列表
		onLineUsers.AddConnList(string(frame.Src.UserID), conn)
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

//处理注册帧
func handleRegister(frame *pb.Frame, conn net.Conn) {
	if Handle.Register(frame.Src) {
		//注册成功
		log.Println("register success")

		//发送返回帧
		//编码
		data, err := EncodeFeedBackProtoc(FeedBack, "", RegisterSuccess, Register)
		if err != nil {
			log.Println(err)
			return
		}
		conn.Write(data)
	} else {
		//注册失败
		log.Println("register fail")
		//发送返回帧
		//编码
		data, err := EncodeFeedBackProtoc(FeedBack, "", RegisterFail, Register)
		if err != nil {
			log.Println(err)
			return
		}
		conn.Write(data)
	}
}

//处理聊天信息帧
func handleChatMsg(frame *pb.Frame, conn net.Conn, rawFrame []byte) {
	log.Println(frame)
	dstUserID := frame.GetDst().GetDst()[0].GetUserID()
	connList := onLineUsers.GetConnList()
	dstConn, exists := connList[string(dstUserID)]
	if exists {
		log.Print("send to user:")
		log.Println(dstUserID)
		dstConn.Write(rawFrame)
	} else {
		log.Println("user offline")
		//handle feedback
	}
}

//消息分发
func msgMux(frame *pb.Frame, rawFrame []byte, conn net.Conn) {
	switch msgType := frame.MsgType; msgType {
	case Login:
		handleLogin(frame, conn)
		break
	case ChatMsg:
		handleChatMsg(frame, conn, rawFrame)
		break
	case Register:
		handleRegister(frame, conn)
		break
	default:
	}
}

//对收到的消息进行解码并且分发
func HandleMsg(request []byte, readlen int, clnConn net.Conn) error {
	//读到标志位，表示这个是特殊帧
	if request[0] == SpFrameHead {
		go handleSPMsg(request, readlen, clnConn)
		return nil
	} else {
		//解码得到信息帧
		frame, err := DeCodeProtoc(request, readlen)
		if err != nil {
			log.Println("failed to unmarshal: ", err)
		}
		fmt.Println(frame)
		//分发消息
		go msgMux(frame, request, clnConn)
		return err
	}

}
