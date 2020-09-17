package main

import (
	"github.com/astaxie/beego/logs"
	"sync"
	"time"
)

func main() {
	logs.Info("测试开始")
	wg := sync.WaitGroup{}
	wg.Add(1)
	go noise(&wg)
	wg.Wait()
	logs.Info("work done")
}

func noise(wg *sync.WaitGroup) {
	defer wg.Done()
	logs.Info("blablabla...")
	time.Sleep(1 * time.Second)
	logs.Info("blablabla...")
	time.Sleep(1 * time.Second)
	logs.Info("blablabla...end")
}
