package main

import "time"

//https://zhuanlan.zhihu.com/p/91044663

func AsyncCall(t int) <-chan int {
	c := make(chan int, 1)
	go func() {
		// simulate real task
		time.Sleep(time.Millisecond * time.Duration(t))
		c <- t
	}()
	return c
}

func AsyncCall2(t int) <-chan int {
	c := make(chan int, 1)
	go func() {
		// simulate real task
		time.Sleep(time.Millisecond * time.Duration(t))
		c <- t
	}()
	// gc or some other reason cost some time
	time.Sleep(200 * time.Millisecond)
	return c
}

func main() {
	select {
	case resp := <-AsyncCall(50):
		println(resp)
	case resp := <-AsyncCall(200):
		println(resp)
	case resp := <-AsyncCall2(3000):
		println(resp)
	}
}
