package main

func VSum(s ...int) int {
	sum := 0
	for _, e := range s {
		sum += e
	}
	return sum
}

func AddBy(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

func ForClosure() []func() {
	funs := []func(){}

	for i := 0; i < 5; i++ {
		funs = append(funs, func() { print(i) })
	}

	return funs
}
