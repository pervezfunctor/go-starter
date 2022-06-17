package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func ClockServer() {
	listener, err := net.Listen("tcp", "localhost:8000")
	Assert(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		fmt.Println("Server is running...")
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
