package main

import (
	"fmt"
	"sort"
)

// With insertion sort
func arrayPairSumII(nums []int) int {
	sum := 0

	for i := 1; i < len(nums); i++ {
		k := nums[i]
		j := i - 1

		for j >= 0 && k < nums[j] {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = k
	}

	for i := 0; i < len(nums)/2; i++ {
		sum += nums[i*2]
	}
	return sum
}

// With Builtin sort
func arrayPairSum(nums []int) int {
	sum := 0
	sort.Ints(nums)
	for i := 0; i < len(nums)/2; i++ {
		sum += nums[i*2]
	}
	return sum
}

func main() {
	a := []int{4, 2, 3, 1}
	res := arrayPairSum(a)
	fmt.Println(res)
}
