// package main

// import "fmt"

// var fibo_cache = make(map[uint64]uint64)

// func fibogen(n uint64) uint64 {
// 	// Generate Fibonacci sequence
// 	// Use memoization

// 	if fibo_cache[n] != 0 {
// 		return fibo_cache[n]
// 	}

// 	var value uint64 = 0

// 	if n == 1 {
// 		value = 1
// 	} else if n == 2 {
// 		value = 1
// 	} else {
// 		value = fibogen(n-1) + fibogen(n-2)
// 	}

// 	fibo_cache[n] = value
// 	return value
// }

// func main() {
// 	var i uint64 = 0
// 	for i = 0; i < 100; i++ {
// 		fmt.Println(fibogen(i + 1))
// 	}
// }
