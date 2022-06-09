package main

import "fmt"

func main() {
	s := []int{3, 1, 7, 5, 8, 4, 3, 2, 9, 5, 10, 21, 4, 3}
	fmt.Printf("%3v\n", s)
	InsertionSort(s)
	fmt.Printf("%5v", s)
}
