package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
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
		ticker = time.NewTicker(500 * time.Millisecond).C
	}
	var nFiles, nBytes int64
	go walkDirWrap(path, fileSizes)

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

func walkDirWrap(path string, fileSizes chan<- int64) {
	walkDir(path, fileSizes)
	close(fileSizes)
}
func walkDir(path string, fileSizes chan<- int64) {
	fileInfos, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du2 error:%v\n", err)
		return
	}

	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			walkDir(filepath.Join(path, fileInfo.Name()), fileSizes)
		} else {
			fileSizes <- fileInfo.Size()
		}
	}
}
