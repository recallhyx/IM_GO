package model

import (
	"../database"
	pb "../example/protoc"
)

//Login 登录,参数为pb.User
func Login(user *pb.User) bool {
	if database.User_Login(user.UserName, user.UserPwd) {
		return true
	}
	return false
}

//Register 注册,参数为pb.User
func Register(user *pb.User) bool {
	if database.User_Register(user.UserName, user.UserPwd) {
		return true
	}
	return false
}
