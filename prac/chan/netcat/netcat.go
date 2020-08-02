package main

import (
	"github.com/astaxie/beego/logs"
	"io"
	"net"
	"os"
)

const tcp = "tcp"

func main() {
	addr := "localhost:8000"
	tcpAddr, err := net.ResolveTCPAddr("", addr)
	conn, err := net.DialTCP(tcp, nil, tcpAddr)
	if err != nil {
		logs.Error(err)
	}

	doneChan := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn)
		logs.Info("work done!")
		doneChan <- struct{}{}
	}()

	mustCopy(conn, os.Stdin)
	conn.CloseWrite()
	<-doneChan //等待后台goroutine完成
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		logs.Error(err)
	}
}
