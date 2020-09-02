package main

import (
	"time"
)

var ch = make(chan int64, 10)

func main() {
	workerInit()
	for {
		now := time.Now().Unix()
		ch <- now
		time.Sleep(300 * time.Millisecond)
	}
}
