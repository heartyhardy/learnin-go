package main

import (
	"math"
)

func findLengthOfLCIS(nums []int) int {
	n, cc, max := math.MinInt32, 0, 0
	for i := 0; i < len(nums); i++ {
		if n < nums[i] {
			cc++
		} else if n >= nums[i] {
			cc = 1
		}
		if max < cc {
			max = cc
		}
		n = nums[i]
	}
	return max
}

// func main() {
// 	a := []int{1, -1, 0, 0, -2}
// 	res := findLengthOfLCIS(a)
// 	fmt.Println(res)
// }
