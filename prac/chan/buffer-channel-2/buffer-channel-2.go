package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	ch := make(chan string)

	start := time.Now()
	//模拟生产者，生产100个事件
	go func() {
		for x := 0; x < 5; x++ {
			ch <- strconv.Itoa(x)
			time.Sleep(1 * time.Second)
			fmt.Println("生产", x)
		}
		close(ch)
	}()

	//goroutine增加消费能力
	//go consume(ch)
	//main线程模拟消费者
	consume(ch)

	sub := time.Now().Sub(start)
	fmt.Println(sub)
}

//消费能力是弱于生产能力的
func consume(ch chan string) {
	for v := range ch {
		time.Sleep(2 * time.Second)
		fmt.Println("消费", v)
	}
}
