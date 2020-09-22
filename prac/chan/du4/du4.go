package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var path string
var verbose bool

func init() {
	flag.StringVar(&path, "path", ".", "traversal path")
	flag.BoolVar(&verbose, "verbose", false, "show verbose step")
}

func main() {
	startTime := time.Now()
	flag.Parse()
	fmt.Println(path, verbose)

	fileSizes := make(chan int64)
	var ticker <-chan time.Time
	if verbose {
		//ticker = time.NewTicker(500 * time.Millisecond).C
		ticker = time.Tick(500 * time.Millisecond)
	}
	var nFiles, nBytes int64
	wg := sync.WaitGroup{}
	walkDirWrap(path, fileSizes, &wg)

	go func() {
		wg.Wait()
		close(fileSizes)
	}()

loop:
	for {
		select {
		case <-ticker:
			printDiskUsage(nFiles, nBytes)
		case size, ok := <-fileSizes:
			if ok {
				nFiles++
				nBytes += size
			} else { //fileSizes is close
				break loop
			}
		}
	}
	printDiskUsage(nFiles, nBytes)

	after := time.Now().Sub(startTime)
	fmt.Println("cost", after)
}

func printDiskUsage(nFiles int64, nBytes int64) {
	fmt.Printf("%d files %.1fMB %.1fGB\n", nFiles, float64(nBytes)/1e6, float64(nBytes)/1e9)
}

func walkDirWrap(path string, fileSizes chan<- int64, wg *sync.WaitGroup) {
	wg.Add(1)
	go walkDir(path, fileSizes, wg)
}
func walkDir(path string, fileSizes chan<- int64, wg *sync.WaitGroup) {
	defer wg.Done()
	fileInfos := readDirChildrenFileInfos(path)

	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			wg.Add(1)
			go walkDir(filepath.Join(path, fileInfo.Name()), fileSizes, wg)
		} else {
			fileSizes <- fileInfo.Size()
		}
	}
}

var sema = make(chan struct{}, 100)

func readDirChildrenFileInfos(dir string) []os.FileInfo {
	sema <- struct{}{}
	defer func() { <-sema }()
	fileInfos, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du4 error:%v\n", err)
	}
	return fileInfos
}
