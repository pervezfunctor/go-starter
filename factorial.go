package main

func Factorial(n int8) int64 {
	var i, fact int64 = 1, 1

	for i <= int64(n) {
		fact *= i
		i++
	}

	return fact
}

func Ncr(n int8, r int8) (ncr int64) {
	return Factorial(n) / (Factorial(r) * Factorial(n-r))
}

func Pascal(lineCount int8) [][]int64 {
	lines := make([][]int64, lineCount+1)
	var n, r int8

	for n = 0; n < lineCount; n++ {
		lines[n] = make([]int64, n+1)

		for r = 0; r <= n; r++ {
			lines[n][r] = Ncr(n, r)
		}
	}

	return lines
}
