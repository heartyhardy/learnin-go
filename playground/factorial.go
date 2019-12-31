package main

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

func reverseString(s string) string {
	if s == "" {
		return s
	}
	return reverseString(s[1:]) + string(s[0])
}

// func main() {
// 	res := factorial(5)
// 	fmt.Println("Factorial", res)

// 	res = fibonacci(6)
// 	fmt.Println("Fibonacci", res)

// 	fmt.Println(reverseString("hello world"))
// }
