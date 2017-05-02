// SetupDB
package database

import (
	"database/sql"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var rootDbPwd = "Mlf7netS"
var dB = "im"
var lock sync.Mutex

//初始化数据库
func SetupDB() {
	var err error
	connStr := "root:" + rootDbPwd + "@/mysql?charset=utf8&loc=Local&parseTime=true"
	db, err = sql.Open("mysql", connStr)
	if err != nil {
		//log(err.Error(), MSG_ERR)
		log.Println(err.Error())
	}
	err = db.Ping()
	if err != nil {
		//log(err.Error(), MSG_ERR)
		log.Println(err.Error())
	}

	cr_db := "CREATE DATABASE IF NOT EXISTS " + dB + " DEFAULT CHARSET utf8 COLLATE utf8_general_ci;"
	exeSQL(cr_db)
	grantSQL := "grant all on " + dB + ".* to root identified by '" + rootDbPwd + "';"
	exeSQL(grantSQL)
	grantSQL = "grant all on " + dB + ".* to root@'' identified by '" + rootDbPwd + "';"
	exeSQL(grantSQL)
	grantSQL = "grant all on " + dB + ".* to root@'localhost' identified by '" + rootDbPwd + "';"
	exeSQL(grantSQL)
	grantSQL = "grant all on " + dB + ".* to root@'127.0.0.1' identified by '" + rootDbPwd + "';"
	exeSQL(grantSQL)

	db.Close()

	connStr = "root:" + rootDbPwd + "@/" + dB + "?charset=utf8&loc=Local&parseTime=true"
	db, err = sql.Open("mysql", connStr)
	if err != nil {
		log.Println(err.Error())
	}
	err = db.Ping()
	if err != nil {
		log.Println(err.Error())
	}

	//初始化表
	exeSQL(`CREATE TABLE IF NOT EXISTS user(
      user_name varchar(35) PRIMARY KEY,
      pwd varchar(128)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;`)

}

//执行sql,不获取返回值
func exeSQL(command string, args ...interface{}) bool {
	lock.Lock()
	stmt, err := db.Prepare(command)
	defer func() {
		if stmt != nil {
			stmt.Close()
		}
		lock.Unlock()
	}()
	if err != nil {
		log.Println(err.Error())
		return false
	}
	if args != nil {
		_, err = stmt.Exec(args...)

	} else {
		_, err = stmt.Exec()
	}

	if err != nil {
		log.Println(err.Error())
		return false
	}
	return true
}

//获取数据，返回数据列表
func exeSQLforResult(command string, args ...interface{}) (*sql.Rows, bool) {
	lock.Lock()
	ok := true
	stmt, err := db.Prepare(command)
	if err != nil {
		log.Println(err.Error())
		ok = false
	}
	var rows *sql.Rows
	if args != nil {
		rows, err = stmt.Query(args...)
	} else {
		rows, err = stmt.Query()
	}

	if err != nil {
		log.Println(err.Error())
		ok = false
	}
	defer func() {
		if stmt != nil {
			stmt.Close()
		}
		lock.Unlock()
	}()
	return rows, ok
}
