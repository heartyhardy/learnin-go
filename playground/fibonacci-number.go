package main

import "fmt"

func fib(N int) int {
	if N < 2 {
		return N
	}
	return fib(N-2) + fib(N-1)
}

func main() {
	fmt.Println(fib(0))
}
