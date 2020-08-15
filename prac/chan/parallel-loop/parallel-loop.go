package main

import (
	"fmt"
	"strconv"
	"time"
)

//已知任务的数量,主goroutine等待多个子goroutine完成
func main() {
	ch := make(chan struct{})
	start := time.Now()
	roundNum := 20
	tasks := createTasks(roundNum)

	//有多少任务，提交多少个任务到goroutine,并通过channel通知完成情况
	for _, v := range tasks {
		go subTask(v, ch)
	}

	//main-goroutine通过通道等待所有子任务完成
	for range tasks {
		<-ch
	}

	cost := time.Now().Sub(start)
	fmt.Println("work done!!!", "cost=", cost)

}

func subTask(x string, ch chan struct{}) {
	time.Sleep(200 * time.Millisecond)
	fmt.Println(x)
	ch <- struct{}{}
}

func createTasks(taskNum int) []string {
	var tasks []string
	for x := 0; x < taskNum; x++ {
		tasks = append(tasks, "t-"+strconv.Itoa(x))
	}
	return tasks
}
