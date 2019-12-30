package main

import "fmt"

func factorial(n int) int {
	if n == 1 {
		return 1
	} else {
		return n * factorial(n-1)
	}
}

func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	res := factorial(5)
	fmt.Println("Factorial", res)

	res = fibonacci(6)
	fmt.Println("Fibonacci", res)
}
