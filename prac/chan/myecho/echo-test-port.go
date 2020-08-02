package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/astaxie/beego/logs"
	"io"
	"net"
	"strings"
	"time"
)

//监听输入
func main() {

	var ip = flag.String("ip", "localhost", "input ip like 192.168.1.66")
	var port = flag.Int("port", 8000, "input port like 8000")
	flag.Parse()

	listenAddr := fmt.Sprintf("%s:%d", *ip, *port)
	tcpAddr, err := net.ResolveTCPAddr("", listenAddr)
	if err != nil {
		logs.Error(err)
	}
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		logs.Error(err)
	}

	logs.Info("listening...", listenAddr)

	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			logs.Error(err)
			continue
		}
		go handle(conn)
	}

}

func handle(conn *net.TCPConn) {
	defer conn.CloseRead()

	for {
		reader := bufio.NewReader(conn)
		line, err := reader.ReadString('\n')
		if err != nil {
			logs.Error(err)
		}
		go echo(line, conn)
	}

}

func echo(line string, conn net.Conn) {
	logs.Info("receive:", line)
	time.Sleep(1 * time.Second)
	io.WriteString(conn, line)
	time.Sleep(1 * time.Second)
	io.WriteString(conn, strings.ToUpper(line))
	time.Sleep(1 * time.Second)
	io.WriteString(conn, strings.ToLower(line))
}
