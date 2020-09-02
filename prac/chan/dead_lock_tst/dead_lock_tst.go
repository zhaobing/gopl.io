package main

import (
	"fmt"
	"time"
)

//在同一个goroutine中对阻塞通道进行读写操作，产生deadLock
//在同一个goroutine中对缓冲通道进行读写操作，不会产生deadLock
func main() {
	//deadLockDemo()
	noDeadLockDemo()
}

//使用阻塞通道，单个goroutine对同一个阻塞通道进行读写,这样就会产生deadLock
func deadLockDemo() {
	ch := make(chan struct{})
	ch <- struct{}{}
	time.Sleep(1 * time.Second)
	<-ch
	fmt.Println("work once")
}

//使用缓冲通道，单个goroutine对同一个缓冲通道进行读写,这样就不会产生deadLock
func noDeadLockDemo() {
	ch := make(chan struct{}, 1)
	ch <- struct{}{}
	time.Sleep(1 * time.Second)
	<-ch
	fmt.Println("work once")
}
