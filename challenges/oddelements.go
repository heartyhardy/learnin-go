package main

import (
	"fmt"
)

func getoddindices(arr []int) []int {
	odds := []int{}
	for i := 0; i < len(arr); i++ {
		if i%2 == 1 {
			odds = append(odds, arr[i])
		}
	}

	return odds
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	odds := getoddindices(arr)

	fmt.Println(odds)
}
