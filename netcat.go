package main

import (
	"io"
	"log"
	"net"
	"os"
)

func Netcat() {
	conn, err := net.Dial("tcp", "localhost:8000")
	ErrHandler(err, conn)

	mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
