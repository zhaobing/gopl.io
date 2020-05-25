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

var vFlag = flag.Bool("v", false, "show vFlag progress message")

func main() {
	flag.Parse()

	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	fileSize := make(chan int64)
	var w sync.WaitGroup
	for _, root := range roots {
		w.Add(1)
		go walkDir(root, &w, fileSize)
	}

	go func() {
		w.Wait()
		close(fileSize)
	}()

	var tick <-chan time.Time
	if *vFlag {
		tick = time.Tick(500 * time.Millisecond)
	}

	var nFiles, nBytes int64

loop:
	for {
		select {
		case fBytes, ok := <-fileSize:
			if !ok {
				break loop
			}
			nFiles++
			nBytes += fBytes
		case <-tick:
			printUsage(nFiles, nBytes)
		}
	}

	fmt.Println("done!")
	printUsage(nFiles, nBytes)
}

func printUsage(nFiles, nBytes int64) {
	fmt.Printf("%d files %.1fMB %.1fGB\n", nFiles, float64(nBytes)/1e6, float64(nBytes)/1e9)
}

func walkDir(dir string, w *sync.WaitGroup, fileSize chan int64) {
	defer w.Done()

	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			w.Add(1)
			subDir := filepath.Join(dir, entry.Name())
			go walkDir(subDir, w, fileSize)
		} else {
			fileSize <- entry.Size()
		}
	}
}

var sema = make(chan struct{}, 20)

func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}
	defer func() { <-sema }()
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}
