package main

import (
	"fmt"
	"gopl.io/ch5/links"
	"log"
)

//并发爬虫
//增加并行度的控制,使用缓冲通道的读写阻塞机制来控制并行度
//使用计数器实现退出机制
func main() {
	workList := make(chan []string)
	var n int //等待发送到任务列表的数量

	n++
	go func() {
		urls := []string{"https://www.yangzhiping.com"}
		workList <- urls
	}()

	seen := make(map[string]bool)
	for ; n > 0; n-- { //使用计数器，控制爬虫的退出
		for linkList := range workList {
			for _, link := range linkList {
				if !seen[link] {
					seen[link] = true
					n++
					go func(link string) {
						workList <- crawl(link)
					}(link)
				}
			}
		}
	}

}

//用来控制并发数量
var tokens = make(chan struct{}, 20)

//!+crawl
func crawl(url string) []string {
	fmt.Println("before read url", url)
	tokens <- struct{}{} //获取令牌
	list, err := links.Extract(url)
	fmt.Println("after read url", url)
	<-tokens //释放令牌
	if err != nil {
		log.Print(err)
	}
	return list
}
