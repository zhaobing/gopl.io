package main

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/zhaobing/bingo/utils/collection"
	"sync"
	"time"
)

//未知任务数量的情况下,任务的传递通过通道
//使用goroutine进行任务处理，每个任务对应一个goroutine
//使用sync.waitGroup做同步，等待所有goroutine的执行完成
//任务执行的结果，通过通道进行通信到主goroutine
func main() {
	files := make(chan string)

	//producer
	go func() {
		for _, file := range collection.GenerateRandomUuids(6) {
			fmt.Println("wait proc file", file)
			files <- file
		}
		close(files)
	}()

	makeThumbnails(files)
	fmt.Println("work done")
}

func makeThumbnails(fileNames <-chan string) int64 {
	sizes := make(chan int64, 10)
	var wg sync.WaitGroup
	for f := range fileNames {
		wg.Add(1)
		//worker
		go func(file string) {
			defer wg.Done()
			file, err := imageFile(f)
			if err != nil {
				logs.Error(err)
			}
			sizes <- int64(len(file))
		}(f)
	}

	//closer
	go func() {
		wg.Wait()
		close(sizes)
	}()

	var total int64
	for size := range sizes {
		total += size
	}
	logs.Info("total", total)
	//flag <- true
	return total
}

func imageFile(fileName string) (string, error) {
	logs.Info("proc file", fileName)
	time.Sleep(1 * time.Second)
	return fileName, nil
}
