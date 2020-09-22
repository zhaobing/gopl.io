package main

import (
	"fmt"
	time2 "github.com/zhaobing/bingo/utils/time"
	"sync"
	"time"
)

//通过close通道的方式，通知多个goroutine退出
func main() {
	t1()
}

func t1() {

	taskNum := 3
	c := make(chan struct{})
	wg := sync.WaitGroup{}
	go func() {
		time2.HeartingPrint(3, 500*time.Millisecond)
		//for i := taskNum; i > 0; i-- {
		//	c <- struct{}{}
		//}

		close(c)
	}()

	for i := taskNum; i > 0; i-- {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()

			select {
			case <-c:
				fmt.Println("Hal-le-lu-jah", "idx", idx)
			}
		}(i)
	}

	wg.Wait()
	fmt.Println("done!")
}
