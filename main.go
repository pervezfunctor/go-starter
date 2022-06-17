package main

import (
	"fmt"
	"time"
)

func go1(ch chan int) {
	time.Sleep(time.Second * 2)
	ch <- 1
}

func go2(ch chan int) {
	time.Sleep(time.Second * 2)
	ch <- 2
}

func SelectExample() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go go1(ch1)
	go go2(ch2)

	for {
		select {
		case <-ch1:
			fmt.Println("ch1")
		case <-ch2:
			fmt.Println("ch2")
		case <-time.After(time.Second * 3):
			fmt.Println("timeout")
			return
		}
	}

}

func main() {
	SpinnerMain()
}
