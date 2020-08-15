package main

import (
	"fmt"
	"gopl.io/ch5/links"
	"log"
)

//并发爬虫
//没有并行度的控制
//也没有退出机制
func main() {
	workList := make(chan []string)

	go func() {
		urls := []string{"https://www.yangzhiping.com"}
		workList <- urls
	}()

	seen := make(map[string]bool)
	for linkList := range workList {
		for _, link := range linkList {
			if !seen[link] {
				seen[link] = true
				go func(link string) {
					workList <- crawl(link)
				}(link)
			}
		}
	}

}

//!+crawl
func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}
