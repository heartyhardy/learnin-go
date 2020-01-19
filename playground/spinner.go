package main

import (
	"fmt"
	"time"
)

func spinner(delay time.Duration) {
	fmt.Println("")
	for {
		for _, v := range `🕛🕐🕑🕒🕓🕔🕕🕖🕗🕘🕙🕚` {
			fmt.Printf("\r%c\t Calculating... \t", v)
			time.Sleep(delay)
		}
	}
}

func fibofast(n, first, second int) int {
	memo := map[int]int{}
	if v, ok := memo[n]; ok {
		return v
	}
	if n == 0 {
		return first
	}
	if n == 1 {
		return second
	}
	res := fibofast(n-1, second, first+second)
	memo[n] = res
	return res
}

func fibo(n int) int {
	if n < 2 {
		return n
	}
	return fibo(n-2) + fibo(n-1)
}

// func main() {
// 	go spinner(100 * time.Millisecond)
// 	res := fibo(40)
// 	fmt.Println("\nFibonacci n:{} is", res)
// }
