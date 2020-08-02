package main

import "github.com/astaxie/beego/logs"

func main() {
	ch := make(chan string, 3)

	ch <- "A"
	ch <- "B"
	ch <- "C"

	logs.Info(<-ch)
	logs.Info(cap(ch))
	logs.Info(len(ch))
	logs.Info(<-ch)
	logs.Info(cap(ch))
	logs.Info(len(ch))
}
