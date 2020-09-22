package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var mu *sync.RWMutex = new(sync.RWMutex)

var m map[string]int

var _m = map[string]int{
	"a": 1,
	"b": 2,
	"c": 3,
	"d": 4,
	"e": 5,
}

var tokens = []string{"a", "b", "c", "d", "e"}
var wg = sync.WaitGroup{}

func main() {

	//mu = new(sync.RWMutex)
	//可以多个同时读
	//go read(1)
	//go read(2)
	//time.Sleep(2 * time.Second)

	wg.Add(1)
	go reload()
	go func() {
		wg.Add(1)
		for i := 0; i < 100; i++ {
			randomRead()
		}
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("done")
}

func reload() {
	mu.Lock()
	defer mu.Unlock()
	time.Sleep(2 * time.Second)
	m = make(map[string]int)
	time.Sleep(2 * time.Second)
	for k, v := range _m {
		m[k] = v
	}

	wg.Done()
}

func randomRead() {
	mu.RLock()
	defer mu.RUnlock()
	k := tokens[rand.Intn(len(tokens))]
	i := m[k]
	if i == 0 {
		fmt.Println("wrong value", k)
	}
}
