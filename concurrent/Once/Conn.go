package main

import (
	"net"
	"sync"
	"time"
)

var connMu sync.RWMutex
var conn net.Conn

func GetConn() net.Conn {
	connMu.Lock()
	defer connMu.Unlock()

	if conn != nil {
		return conn
	}
	conn, _ := net.DialTimeout("tcp","baidu.com", 10 * time.Second)
	return conn
}

func main() {
	conn := GetConn()
	if conn == nil {
		panic("err" +
			"")
	}
}
