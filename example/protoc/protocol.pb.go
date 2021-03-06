// Code generated by protoc-gen-go.
// source: protocol.proto
// DO NOT EDIT!

/*
Package protocol is a generated protocol buffer package.

It is generated from these files:
	protocol.proto

It has these top-level messages:
	User
	Msg
	DstUser
	Action
	Frame
*/
package protocol

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type User struct {
	UserName string `protobuf:"bytes,1,opt,name=userName" json:"userName,omitempty"`
	UserPwd  string `protobuf:"bytes,2,opt,name=userPwd" json:"userPwd,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *User) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *User) GetUserPwd() string {
	if m != nil {
		return m.UserPwd
	}
	return ""
}

type Msg struct {
	Msg string `protobuf:"bytes,1,opt,name=msg" json:"msg,omitempty"`
}

func (m *Msg) Reset()                    { *m = Msg{} }
func (m *Msg) String() string            { return proto.CompactTextString(m) }
func (*Msg) ProtoMessage()               {}
func (*Msg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Msg) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type DstUser struct {
	Dst []*User `protobuf:"bytes,1,rep,name=dst" json:"dst,omitempty"`
}

func (m *DstUser) Reset()                    { *m = DstUser{} }
func (m *DstUser) String() string            { return proto.CompactTextString(m) }
func (*DstUser) ProtoMessage()               {}
func (*DstUser) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *DstUser) GetDst() []*User {
	if m != nil {
		return m.Dst
	}
	return nil
}

type Action struct {
	RslCode    int32  `protobuf:"varint,1,opt,name=rslCode" json:"rslCode,omitempty"`
	ActionType int32  `protobuf:"varint,2,opt,name=actionType" json:"actionType,omitempty"`
	RslMsg     string `protobuf:"bytes,3,opt,name=rslMsg" json:"rslMsg,omitempty"`
}

func (m *Action) Reset()                    { *m = Action{} }
func (m *Action) String() string            { return proto.CompactTextString(m) }
func (*Action) ProtoMessage()               {}
func (*Action) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Action) GetRslCode() int32 {
	if m != nil {
		return m.RslCode
	}
	return 0
}

func (m *Action) GetActionType() int32 {
	if m != nil {
		return m.ActionType
	}
	return 0
}

func (m *Action) GetRslMsg() string {
	if m != nil {
		return m.RslMsg
	}
	return ""
}

// 登录，注册，添加好友，删除好友，发送消息
type Frame struct {
	ProtoSign  int32    `protobuf:"varint,1,opt,name=protoSign" json:"protoSign,omitempty"`
	MsgLength  int64    `protobuf:"varint,2,opt,name=msgLength" json:"msgLength,omitempty"`
	MsgType    int32    `protobuf:"varint,3,opt,name=msgType" json:"msgType,omitempty"`
	SenderTime int64    `protobuf:"varint,4,opt,name=senderTime" json:"senderTime,omitempty"`
	Src        *User    `protobuf:"bytes,5,opt,name=src" json:"src,omitempty"`
	Dst        *DstUser `protobuf:"bytes,6,opt,name=dst" json:"dst,omitempty"`
	Msg        *Msg     `protobuf:"bytes,7,opt,name=msg" json:"msg,omitempty"`
	FbAction   *Action  `protobuf:"bytes,8,opt,name=fbAction" json:"fbAction,omitempty"`
}

func (m *Frame) Reset()                    { *m = Frame{} }
func (m *Frame) String() string            { return proto.CompactTextString(m) }
func (*Frame) ProtoMessage()               {}
func (*Frame) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Frame) GetProtoSign() int32 {
	if m != nil {
		return m.ProtoSign
	}
	return 0
}

func (m *Frame) GetMsgLength() int64 {
	if m != nil {
		return m.MsgLength
	}
	return 0
}

func (m *Frame) GetMsgType() int32 {
	if m != nil {
		return m.MsgType
	}
	return 0
}

func (m *Frame) GetSenderTime() int64 {
	if m != nil {
		return m.SenderTime
	}
	return 0
}

func (m *Frame) GetSrc() *User {
	if m != nil {
		return m.Src
	}
	return nil
}

func (m *Frame) GetDst() *DstUser {
	if m != nil {
		return m.Dst
	}
	return nil
}

func (m *Frame) GetMsg() *Msg {
	if m != nil {
		return m.Msg
	}
	return nil
}

func (m *Frame) GetFbAction() *Action {
	if m != nil {
		return m.FbAction
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "protocol.User")
	proto.RegisterType((*Msg)(nil), "protocol.Msg")
	proto.RegisterType((*DstUser)(nil), "protocol.DstUser")
	proto.RegisterType((*Action)(nil), "protocol.Action")
	proto.RegisterType((*Frame)(nil), "protocol.Frame")
}

func init() { proto.RegisterFile("protocol.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 311 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x50, 0x3b, 0x4f, 0xf3, 0x30,
	0x14, 0x55, 0xea, 0xe6, 0xd1, 0xfb, 0xe9, 0xab, 0x8a, 0x07, 0xb0, 0x10, 0x82, 0x28, 0x2c, 0x95,
	0x40, 0x1d, 0x60, 0x65, 0x41, 0x20, 0x26, 0x82, 0x90, 0x29, 0x0b, 0x5b, 0x9b, 0x18, 0x13, 0x29,
	0x8f, 0xca, 0x37, 0x08, 0xf1, 0x3b, 0xf8, 0xc3, 0xc8, 0xd7, 0x79, 0x74, 0x60, 0xf3, 0x79, 0xf8,
	0xde, 0x73, 0x0f, 0xcc, 0x77, 0xa6, 0x69, 0x9b, 0xac, 0x29, 0x57, 0xf4, 0xe0, 0x51, 0x8f, 0x93,
	0x1b, 0x98, 0xbe, 0xa2, 0x32, 0xfc, 0x18, 0xa2, 0x4f, 0x54, 0xe6, 0x69, 0x53, 0x29, 0xe1, 0xc5,
	0xde, 0x72, 0x26, 0x07, 0xcc, 0x05, 0x84, 0xf6, 0xfd, 0xfc, 0x95, 0x8b, 0x09, 0x49, 0x3d, 0x4c,
	0x8e, 0x80, 0xa5, 0xa8, 0xf9, 0x02, 0x58, 0x85, 0xba, 0xfb, 0x67, 0x9f, 0xc9, 0x05, 0x84, 0xf7,
	0xd8, 0xd2, 0xe4, 0x18, 0x58, 0x8e, 0xad, 0xf0, 0x62, 0xb6, 0xfc, 0x77, 0x35, 0x5f, 0x0d, 0x49,
	0xac, 0x28, 0xad, 0x94, 0xbc, 0x41, 0x70, 0x9b, 0xb5, 0x45, 0x53, 0xdb, 0x4d, 0x06, 0xcb, 0xbb,
	0x26, 0x77, 0x21, 0x7c, 0xd9, 0x43, 0x7e, 0x0a, 0xb0, 0x21, 0xcf, 0xfa, 0x7b, 0xa7, 0x28, 0x86,
	0x2f, 0xf7, 0x18, 0x7e, 0x08, 0x81, 0xc1, 0x32, 0x45, 0x2d, 0x18, 0xa5, 0xe8, 0x50, 0xf2, 0x33,
	0x01, 0xff, 0xc1, 0xd8, 0x2b, 0x4e, 0x60, 0x46, 0xbb, 0x5f, 0x0a, 0x5d, 0x77, 0xd3, 0x47, 0xc2,
	0xaa, 0x15, 0xea, 0x47, 0x55, 0xeb, 0xf6, 0x83, 0xc6, 0x33, 0x39, 0x12, 0x36, 0x57, 0x85, 0x9a,
	0x56, 0x33, 0x97, 0xab, 0x83, 0x36, 0x17, 0xaa, 0x3a, 0x57, 0x66, 0x5d, 0x54, 0x4a, 0x4c, 0xe9,
	0xe3, 0x1e, 0x63, 0xaf, 0x47, 0x93, 0x09, 0x3f, 0xf6, 0xfe, 0xba, 0x1e, 0x4d, 0xc6, 0xcf, 0x5d,
	0x3f, 0x01, 0x39, 0x0e, 0x46, 0x47, 0xd7, 0x1f, 0x55, 0xc4, 0xcf, 0x5c, 0xc3, 0x21, 0x99, 0xfe,
	0x8f, 0xa6, 0x14, 0x35, 0x15, 0xce, 0x2f, 0x21, 0x7a, 0xdf, 0xba, 0x16, 0x45, 0x44, 0xae, 0xc5,
	0xe8, 0x72, 0xbc, 0x1c, 0x1c, 0xdb, 0x80, 0xa4, 0xeb, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x02,
	0xdb, 0xe0, 0x66, 0x18, 0x02, 0x00, 0x00,
}
