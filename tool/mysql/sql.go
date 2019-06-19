package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var (
	Err error
	Db  *sql.DB
)

func init() {
	fmt.Printf("hello world")
	Db, Err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test")
}
