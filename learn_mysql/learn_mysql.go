package learn_mysql

import "database/sql"

const (
	USER_NAME = "root"
	PASS_WORD = "root"
	HOST      = "localhost"
	PORT      = "3306"
	DATABASE  = "golang"
	CHARSET   = "utf8"
)

var mysql_db *sql.DB
var mysql_err error

func sql_link() {

}

func sql_create() {

}
