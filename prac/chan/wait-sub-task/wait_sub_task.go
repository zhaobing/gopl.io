package main

import (
	"github.com/astaxie/beego/logs"
	"time"
)

func main() {
	logs.Info("测试")
	workDone := make(chan struct{})
	go noise(workDone)
	<-workDone
	logs.Info("work done")
}

func noise(c chan struct{}) {
	logs.Info("blablabla...")
	time.Sleep(1 * time.Second)
	logs.Info("blablabla...")
	time.Sleep(1 * time.Second)
	logs.Info("blablabla...")
	time.Sleep(1 * time.Second)
	logs.Info("blablabla...")
	time.Sleep(1 * time.Second)
	logs.Info("blablabla...")
	time.Sleep(1 * time.Second)
	c <- struct{}{}
}
