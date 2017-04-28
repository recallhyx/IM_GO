package Handle

import (
	"log"

	"../database"
	pb "../example/protoc"
)

//HandleLogin 登录,参数为pb.User
func HandleLogin(user *pb.User) bool {
	log.Println("login..check.")
	//验证数据库里的情况
	loginFlag := database.User_Login(user.UserName, user.UserPwd)
	//发送返回帧
	//编码
	//data, err := EncodeFeedBackProtoc(1, "lzy", 1, 1, "login ok")
	//if err != nil {
	//	log.Println(err)
	//	return false
	//}

	return loginFlag
}

//Register 注册,参数为pb.User
func Register(user *pb.User) bool {
	if database.User_Register(user.UserName, user.UserPwd) {
		return true
	}
	return false
}
