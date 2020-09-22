package main

import (
	"bufio"
	"fmt"
	"github.com/astaxie/beego/logs"
	"net"
)

type client chan string

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

func main() {
	fmt.Println("start charting...")
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		logs.Error(err)
		return
	}

	go brodcast()

	for {
		conn, err := listener.Accept()
		if err != nil {
			logs.Error(err)
			continue
		}
		go handleConn(conn)
	}

}

func brodcast() {
	clients := make(map[client]bool)
	for {

		select {
		case cli := <-entering:
			clients[cli] = true
		case cli := <-leaving:
			close(cli)
			delete(clients, cli)
		case msg := <-messages:
			for cli := range clients {
				cli <- msg
			}

		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(client)
	go clientWrite(ch, conn)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who + "\n"
	messages <- who + " is entered\n"
	entering <- ch

	input := bufio.NewScanner(conn)
	for input.Scan() {
		text := input.Text()
		messages <- who + ":" + text + "\n"
	}

	leaving <- ch
	messages <- who + " is leaved\n"
	conn.Close()
}

func clientWrite(ch client, conn net.Conn) {
	for msg := range ch {
		fmt.Fprint(conn, msg)
	}
}
