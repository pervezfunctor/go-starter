package main

func Gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}
