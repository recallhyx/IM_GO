// InterfaceOfDB
package database

import (
	//"errors"
	"log"
	//"strconv"
	_ "github.com/go-sql-driver/mysql"
)

func User_Login(userName, pwd string) bool {
	Existsql := "select user_name from user where user_name=(?) and pwd=(?)"
	rows, staute := exeSQLforResult(Existsql, userName, pwd)
	if staute != true {
		log.Println("check_pwd_error")
		return false
	}
	defer rows.Close()
	if rows.Next() {
		return true
	}
	return false
}
func User_Register(userName, pwd string) bool {
	keepsql := "insert into user(user_name,pwd) values(?,?)"
	staute := exeSQL(keepsql, userName, pwd)
	if staute == true {
		return true
	} else {
		log.Println("insert_newUser_error")
		return false
	}
}

func AddFriend(ownername,friendname string)bool{
	keepsql := "insert into friend_info(ownerid,friendid) values(?,?)"
	staute := exeSQL(keepsql, ownername,friendname)
	if staute == true {
		return true
	} else {
		log.Println("Addfriend_error")
		return false
	}
}

func DelFriend(ownername,friendname string)bool{
	keepsql := "delete * from friend_info where ownerid=(?) and friendid=(?)"
	staute := exeSQL(keepsql, ownername,friendname)
	if staute == true {
		return true
	} else {
		log.Println("Delfriend_error")
		return false
	}
}

func GetFriend(ownerID string) ([]string, bool) { //查找好友，参数用户id，返回好友id的int数组
	seletesql := "select ownerid from friend_info where ownerid=(?)"
	rows, staute := exeSQLforResult(seletesql, ownerID)
	var friendname []string
	if staute != true {
		log.Println("getFriend_error")
	} else {
		defer rows.Close()
		var name string;
		for rows.Next() {
			rows.Scan(&name)
			friendname = append(friendname, name)
		}
		return friendname, true
	}
	return friendname, false
}



