package main

import (
	"database/sql"
	"fmt"
	_ "github.com/xeodou/go-sqlcipher"
)

func main() {
	b, err := sql.Open("sqlite3", "wxmsg.db"+"?_key=c160836")
	if err != nil {
		fmt.Println(err)
		return
	}
	result, err := b.Exec(".database;")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)

	defer b.Close()
}
