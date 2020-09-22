package main

import (
	"sync"
	"time"
)

var m *sync.RWMutex

func main() {

	m = new(sync.RWMutex)
	//可以多个同时读
	go read(1)
	go read(2)
	time.Sleep(2 * time.Second)
}

func read(i int) {
	println(i, "read start")
	m.RLock()
	println(i, "reading")
	time.Sleep(1 * time.Second)
	m.RUnlock()
	println(i, "read end")
}
