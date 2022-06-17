package main

import (
	"io"
	"log"
	"math/rand"
	"time"
)

func RandSleep() {
	time.Sleep(500 + time.Duration(rand.Intn(500))*time.Millisecond)
}

func RandInt(start, stop int) int {
	return start + rand.Intn(stop-start)
}
func Assert(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ErrHandler(err error, closer io.Closer) {
	if err != nil {
		log.Fatal(err)
	}

	if closer != nil {
		closer.Close()
	}
}

func Cancelled(done chan bool) bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}
