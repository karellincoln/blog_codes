package main

import (
	"net"
	"log"
	"io"
	"time"
	"fmt"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:9202")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(time.Now().Format(time.RFC3339))

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()

	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}