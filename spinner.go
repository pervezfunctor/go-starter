package main

import (
	"fmt"
	"time"
)

func spinner(done chan bool, delay time.Duration) {
	for {
		select {
		case <-done:
			fmt.Println()
			return
		default:
			for _, r := range `-\|/` {
				fmt.Printf("\r%c", r)
				time.Sleep(delay)
			}
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}

func SpinnerMain() {
	done := make(chan bool)
	go spinner(done, 100*time.Millisecond)
	const n = 45
	fibN := fib(n)
	done <- true
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}
