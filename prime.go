package main

import (
	"sync"
)

func isPrime(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func AddPrimes(nos, primes chan int) {
	for i := range nos {
		if isPrime(i) {
			primes <- i
		}
	}
}

func GenNos(nos chan int) {
	for i := 2; i <= 1000000; i++ {
		nos <- i
	}
	close(nos)
}

func Foomain() {
	nos := make(chan int)
	primes := make(chan int)
	var done sync.WaitGroup

	go GenNos(nos)

	for i := 0; i <= 16; i++ {
		done.Add(1)
		go func() {
			defer done.Done()
			AddPrimes(nos, primes)
		}()
	}

	go func() {
		done.Wait()
		close(primes)
	}()

	var s int64 = 0
	for i := range primes {
		s += int64(i)
	}

	println(s)
}
