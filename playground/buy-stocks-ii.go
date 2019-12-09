package main

import "fmt"

func maxProfit(prices []int) int {

	maxprofit := 0

	for i := 0; i < len(prices)-1; i++ {
		if profit := prices[i+1] - prices[i]; profit > 0 {
			maxprofit += prices[i+1] - prices[i]
		}
	}

	return maxprofit
}

func main() {
	stocks := []int{1, 2, 4, 3, 10}
	res := maxProfit(stocks)
	fmt.Println(res)
}
