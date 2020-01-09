package main

import (
	"math"
	"sort"
)

func maximumProductII(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	a := nums[0] * nums[1] * nums[2]
	b := nums[0] * nums[len(nums)-1] * nums[len(nums)-2]
	if a > b {
		return a
	}
	return b
}

func maximumProduct(nums []int) int {

	max1, max2, max3 := math.MinInt32, math.MinInt32, math.MinInt32
	min1, min2 := math.MaxInt32, math.MaxInt32

	for _, num := range nums {
		if num >= max1 {
			max3, max2, max1 = max2, max1, num
		} else if num >= max2 {
			max3, max2 = max2, num
		} else if num > max3 {
			max3 = num
		}
		if num <= min1 {
			min2, min1 = min1, num
		} else if num < min2 {
			min2 = num
		}
	}
	if max1*max2*max3 > min1*min2*max1 {
		return max1 * max2 * max3
	}
	return min1 * min2 * max1
}

// func main() {
// 	a := []int{-4, -3, -2, 60}
// 	res := maximumProduct(a)
// 	fmt.Println(a, res)
// }
