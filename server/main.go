package main 

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var port string = ":3030"

func main () {

	db, err := sql.Open("mysql", "root:@/todos_db?charset=utf8")
	CheckError(err)

	CreateApiServer(db)

}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}