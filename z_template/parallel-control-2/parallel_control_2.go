package main

import (
	"fmt"
	time2 "github.com/zhaobing/bingo/utils/time"
	"sync"
	"time"
)

//限定同时执行某个任务的goroutine的数量
//后台启动一个goroutine持续的订阅某个任务通道
//每从任务通道中接收到一个任务，就启动一个goroutine来执行该任务,直到该任务通道关闭
//任务执行者方法中，每次开始执行任务前，都向控制goroutine并行度的令牌缓冲通道中写入一个标志，表示占用一个并行度信号量
//任务执行完成后,释放该令牌
//利用缓冲通道的阻塞机制，来控制并行度

var taskCh = make(chan string)
var wg = sync.WaitGroup{}
var workRoutineNum = 10
var sema = make(chan struct{}, workRoutineNum)

func main() {
	defer time2.TraceCost("parallel control 2")()
	taskNum := 20
	wg.Add(taskNum)
	go taskProduce(taskNum)
	go worker()
	wg.Wait()
	fmt.Println("work done.")
}

func taskProduce(taskNum int) {
	for i := 0; i < taskNum; i++ {
		task := fmt.Sprintf("t-%d", i)
		taskCh <- task
	}
}

func worker() {
	for task := range taskCh {
		go func(task string) {
			sema <- struct{}{}
			time.Sleep(2 * time.Second)
			fmt.Printf("exec task-%s\n", task)
			<-sema
			wg.Done()
		}(task)
	}
}
