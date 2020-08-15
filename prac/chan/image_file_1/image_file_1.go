package main

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/zhaobing/bingo/utils/collection"
	"time"
)

//在已知任务数量的情况下，使用与任务数量相等的goroutine执行任务，并等待任务的完成
func main() {
	files := collection.GenerateRandomUuids(6)
	makeThumbnails(files)
	fmt.Println("work done")
}

func makeThumbnails(files []string) {
	ch := make(chan struct{})

	//在已知任务数量的情况下,调用与任务数量相同的goroutine来执行任务
	for _, fileName := range files {
		go func(f string) {
			imageFile(f)
			ch <- struct{}{}
		}(fileName)
	}

	//等待所有goroutine完成
	for range files {
		<-ch
	}

}

func imageFile(fileName string) {
	logs.Info("proc file", fileName)
	time.Sleep(1 * time.Second)
}
