package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)
	fmt.Printf("%T, %v", now, now)

	now.Year()
	now.Month()
	now.Day()
	now.Hour()
	now.Format("2006-01-02 15:04:05")


	// 时间常量
	/**
		const (
		Nanosecond  Duration = 1
		Microsecond          = 1000 * Nanosecond
		Millisecond          = 1000 * Microsecond
		Second               = 1000 * Millisecond
		Minute               = 60 * Second
		Hour                 = 60 * Minute
	)
	*/

	// 获取随机数字 time.unix() unixNano()
	// time.sleep()

}
