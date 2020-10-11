package main

import (
	"database/sql"
	"fmt"
	_ "github.com/xeodou/go-sqlcipher"
)

func main() {
	b, err := sql.Open("sqlite3", ""+"?_key=password")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer b.Close()
}
