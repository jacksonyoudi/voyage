package main

import "time"

var startTime = time.Now()

var startTime1 time.Time

func init() {
	startTime1 = time.Now()
}

var startTime2 time.Time

func initApp() {
	startTime2 = time.Now()
}

func main() {
	initApp()
}
