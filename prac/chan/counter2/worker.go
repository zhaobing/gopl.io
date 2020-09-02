package main

import "fmt"

func workerInit() {

	for i := 0; i < 20; i++ {
		go func(idx int) {
			for ts := range ch {
				fmt.Println("worker-id", idx, ts)
			}
		}(i)
	}

}
