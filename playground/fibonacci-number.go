package main

// import "fmt"

// //With memoization
// func fib3(N int) int {
// 	memo := map[int]int{}

// 	if v, ok := memo[N]; ok {
// 		return v
// 	}

// 	if N < 2 {
// 		return N
// 	}
// 	res := fib(N-2) + fib(N-1)
// 	memo[N] = res
// 	return res
// }

// //Reucsion
// func fib2(N int) int {
// 	if N < 2 {
// 		return N
// 	}
// 	return fib(N-2) + fib(N-1)
// }

// //Loop
// func fib(N int) int {
// 	l, r := 0, 1
// 	for ; N > 0; N-- {
// 		l, r = r, l+r
// 	}
// 	return l
// }

// func main() {
// 	fmt.Println(fib(30))
// }
