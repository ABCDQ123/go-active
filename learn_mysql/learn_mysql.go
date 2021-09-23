package learn_mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

const (
	USER_NAME = "root"
	PASS_WORD = "root"
	NET       = "tcp"
	HOST      = "localhost"
	PORT      = "3306"
	DATABASE  = "mysql"
	CHARSET   = "utf8"
)

var sql_db *sql.DB
var sql_err error

func sql_open() {
	sql_string := fmt.Sprintf("%s:%s@%s(%s:%s)/%s?charset=%s", USER_NAME, PASS_WORD, NET, HOST, PORT, DATABASE, CHARSET)
	sql_db, sql_err = sql.Open(DATABASE, sql_string)
	if sql_err != nil {
		fmt.Println("mysql Sql Open err", sql_err.Error())
		return
	}
	sql_db.SetMaxOpenConns(10)
	sql_db.SetMaxIdleConns(10)
	sql_db.SetConnMaxLifetime(time.Minute)
	sql_err = nil
	if sql_err = sql_db.Ping(); sql_err != nil {
		panic("mysql Sql ping err: " + sql_err.Error())
	}
}

func sql_table_create() {
	//**exec query prepare
	//sql_db.Exec()
	//sql_db.Query()
	//stm ,err :=sql_db.Prepare()
	//stm.Exec()
	//stm.Query()
	//**bein
	//tx, err :=sql_db.Begin()
	//tx.Exec()
	//tx.Query()
	//stm ,err =tx.Prepare()
	//stm.Exec()
	//stm.Query()
	//tx.Commit()
}
