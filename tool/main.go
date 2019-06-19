package main

import (
	"fmt"
	"github.com/ZRothschild/algorithm/tool/mysql"
)

func main() {
	if mysql.Err != nil {
		fmt.Printf("%v", mysql.Err)
	}
	stmt, err := mysql.Db.Prepare(`insert users (name,age,update_time,add_time) values (?,?,?,?),(?,?,?,?)`)

	if err != nil {
		fmt.Printf("%v", err)
	}
	result, er := stmt.Exec("hello", 14, 123, 321, "hel", 14, 123, 321)
	if err != nil {
		fmt.Printf("%v", er)
	}
	id, _ := result.LastInsertId()
	fmt.Printf("%d", id)
}
