package main

import (
	"fmt"
	"time"
)

func main() {
	t1()
}

//使用阻塞通道，单个goroutine对同一个阻塞通道进行读写
func t1() {
	ch := make(chan struct{}, 1)
	ch <- struct{}{}
	time.Sleep(1 * time.Second)
	<-ch
	fmt.Println("work once")
}
