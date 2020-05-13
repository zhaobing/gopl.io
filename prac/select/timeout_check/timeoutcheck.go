package main

import (
	"fmt"
	"time"
)

func worker(ch chan int, sleep int) {
	time.Sleep(time.Duration(sleep) * time.Second)
	ch <- 200
}

func main() {
	ch := make(chan int)
	go worker(ch, 2)

	select {
	case stat := <-ch:
		fmt.Println("stat", stat)
	case <-time.After(3 * time.Second):
		fmt.Println("error-stat", "time out")
	}

}
