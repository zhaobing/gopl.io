package main

import (
	"github.com/astaxie/beego/logs"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go test(&wg)
	wg.Wait()
	logs.Info("work done")
}

func test(wg *sync.WaitGroup) {
	defer wg.Done()
	logs.Info("start....")
	time.Sleep(1 * time.Second)
	logs.Info("end....")
}
