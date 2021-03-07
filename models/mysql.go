package models

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

const CONNECT_STR = "root:root@tcp(192.168.3.104:3306)/test"

var Db *sql.DB

func init() {
	var err error
	if Db, err = sql.Open("mysql", CONNECT_STR); err != nil {
		fmt.Printf(err.Error())
	}

}
