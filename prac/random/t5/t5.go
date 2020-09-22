package main

import (
	"fmt"
	time2 "github.com/zhaobing/bingo/utils/time"
	"time"
)

func main() {
	defer time2.TraceCost("测试")()
	time.Sleep(1 * time.Second)
	fmt.Println("done")
}
