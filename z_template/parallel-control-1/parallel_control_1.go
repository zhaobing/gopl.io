package main

import (
	"fmt"
	time2 "github.com/zhaobing/bingo/utils/time"
	"sync"
	"time"
)

//限定同时执行某个任务的goroutine的数量
//后台启动一个模拟协程池，持续的订阅某个任务通道
//直到任务通道关闭前，都从这个通道中接收任务
//模拟协程池的协程数量是可以设定的

var taskCh = make(chan string)
var wg = sync.WaitGroup{}

func main() {
	defer time2.TraceCost("parallel control 1")()
	taskNum := 20
	wg.Add(taskNum)
	go taskProduce(taskNum)
	go workers(10)
	wg.Wait()
	fmt.Println("work done.")
}

func taskProduce(taskNum int) {
	for i := 0; i < taskNum; i++ {
		task := fmt.Sprintf("t-%d", i)
		taskCh <- task
	}
}

func workers(workNum int) {
	for i := 0; i < workNum; i++ {
		go worker(i)
	}
}

func worker(workerId int) {
	for task := range taskCh {
		time.Sleep(900 * time.Millisecond)
		fmt.Printf("wk-%d exec task-%s\n", workerId, task)
		wg.Done()
	}
}
