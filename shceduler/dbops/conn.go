package dbops

import (
	"database/sql"

	"github.com/lunny/log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dbConn *sql.DB
	err    error
)

func init() {
	dbConn, err = sql.Open("mysql", "root:12345678@tcp(localhost:3306)/video_server?charset=utf8")
	if err != nil {
		log.Println("connect db is error")
		panic(err)
	}
	err = dbConn.Ping()
	if err != nil {
		log.Println("ping db is error")
		panic(err)
	}

}
