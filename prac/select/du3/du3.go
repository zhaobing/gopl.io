package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
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
