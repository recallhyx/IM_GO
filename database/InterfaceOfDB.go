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

	Existsql := "select user_name from user where user_name=(?) and pwd = (?)"

	rows,staute := exeSQLforResult(Existsql,userName,pwd)
	log.Println(staute)
	defer rows.Close()

	if staute == true {
		if rows.Next() {
			// 存在用户，拒绝注册
			return false
		}
		log.Println("can_register")
		Keepsql := "insert into user(user_name,pwd) values(?,?)"
		exeSQL(Keepsql,userName,pwd)
		return true
	} else {
		log.Println("check_user_error")
		return false
	}
}
