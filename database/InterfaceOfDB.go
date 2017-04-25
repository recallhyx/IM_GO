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
