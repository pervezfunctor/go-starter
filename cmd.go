package main

import "os"

func Cmd() string {
	s, sep := "", ""

	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}

	return s
}

// fmt.Println(strings.Join(os.Args[1:], " "))
