package learn_mysql

import (
	"database/sql"
	"fmt"
	"time"
)

const (
	DATABASE  = "mysql"
	USER_NAME = "root"
	PASS_WORD = "root"
	NET       = "tcp"
	HOST      = "localhost"
	PORT      = "3306"
	CHARSET   = "utf8"
)

var mysql_db *sql.DB
var mysql_err error

func sql_open() {
	sql_string := fmt.Sprintf("%s:%s@%s(%s:%s)/%s?charset=%s", USER_NAME, PASS_WORD, NET, HOST, PORT, DATABASE, CHARSET)
	mysql_db, mysql_err = sql.Open(DATABASE, sql_string)
	if mysql_err != nil {
		fmt.Println("mysql Sql Open err", mysql_err.Error())
		return
	}
	mysql_db.SetMaxOpenConns(100)
	mysql_db.SetMaxIdleConns(20)
	mysql_db.SetConnMaxLifetime(time.Second * 100)
	mysql_err = nil
	if mysql_err = mysql_db.Ping(); mysql_err != nil {
		panic("mysql Sql ping err: " + mysql_err.Error())
	}
}

func sql_table_create() {

}
