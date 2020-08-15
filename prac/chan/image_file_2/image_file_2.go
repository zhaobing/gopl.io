package main

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/zhaobing/bingo/utils/collection"
	"time"
)

//在已知任务数量的情况下，使用与任务数量相同的goroutine来执行任务
//执行任务的过程中，返回任务的执行情况
//收集任务执行情况（包括错误情况），并且等待任务集合执行完成
func main() {
	files := collection.GenerateRandomUuids(6)
	makeThumbnails(files)
	fmt.Println("work done")
}

func makeThumbnails(files []string) (thumbFiles []string, err error) {

	type item struct {
		thumbFile string
		err       error
	}

	ch := make(chan item, len(files))

	//在已知任务数量的情况下,调用与任务数量相同的goroutine来执行任务
	for _, fileName := range files {
		go func(f string) {
			var it item
			it.thumbFile, it.err = imageFile(f)
			ch <- it
		}(fileName)
	}

	//等待所有goroutine完成
	for range files {
		it := <-ch
		if it.err != nil {
			return nil, it.err
		}
		thumbFiles = append(thumbFiles, it.thumbFile)
	}

	return thumbFiles, nil
}

func imageFile(fileName string) (string, error) {
	logs.Info("proc file", fileName)
	time.Sleep(1 * time.Second)
	return fileName, nil
}
