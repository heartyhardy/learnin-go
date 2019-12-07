package main

import (
	"fmt"
	"math"
)

// func maxProfit(prices []int) int {
// 	// Naive solution
// 	profit := 0
// 	for i := 0; i < len(prices); i++ {
// 		for j := i; j < len(prices); j++ {
// 			if profit < prices[j]-prices[i] {
// 				profit = prices[j] - prices[i]
// 			}
// 		}
// 	}

// 	return profit
// }

func maxProfit(prices []int) int {
	buy, sell := math.MaxInt32, 0
	l, i := len(prices), 0

	for i < l {
		if buy > prices[i] {
			buy = prices[i]
		} else if (prices[i] - buy) > sell {
			sell = prices[i] - buy
		}
		i++
	}
	return sell
}

func main() {
	a := []int{7, 1, 5, 3, 6, 4}
	res := maxProfit(a)
	fmt.Println(res)
}
