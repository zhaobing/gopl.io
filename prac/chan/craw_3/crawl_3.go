package main

import (
	"gopl.io/ch5/links"
	"log"
)

const Parallelism = 20

//并发爬虫
//增加并行度的控制,使用计数来控制并行度
//使用计数器实现退出机制
func main() {
	workList := make(chan []string)
	unseenLinks := make(chan string)

	go func() {
		urls := []string{"https://www.yangzhiping.com"}
		workList <- urls
	}()

	for i := 0; i < Parallelism; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
				go func() {
					workList <- foundLinks
				}()
			}
		}()
	}

	seen := make(map[string]bool)
	for linkList := range workList {
		for _, link := range linkList {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}

}

//!+crawl
func crawl(url string) []string {
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}
