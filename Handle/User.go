package Handle

import (
	"IM_GO/database"

	pb "IM_GO/example/protoc"
)

//Login 登录,参数为pb.User
func Login(user *pb.User) bool {
	//验证数据库里的情况
	loginFlag := database.User_Login(user.UserName, user.UserPwd)
	return loginFlag
}

//Register 注册,参数为pb.User
func Register(user *pb.User) bool {
	if database.User_Register(user.UserName, user.UserPwd) {
		return true
	}
	return false
}
