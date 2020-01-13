package main

import "fmt"

func pivotIndexII(nums []int) int {
	if len(nums) > 0 {
		suml := make([]int, len(nums))
		sumr := make([]int, len(nums))
		suml[0], sumr[len(nums)-1] = nums[0], nums[len(nums)-1]
		for i := 1; i < len(nums); i++ {
			suml[i] = suml[i-1] + nums[i]
			sumr[len(nums)-1-i] = sumr[len(nums)-i] + nums[len(nums)-1-i]
		}
		for k, v := range suml {
			if v == sumr[k] {
				return k
			}
		}
	}
	return -1
}

func pivotIndex(nums []int) int {
	sum, l, r := 0, 0, 0
	for _, v := range nums {
		sum += v
	}
	for i, v := range nums {
		r = sum - l - v
		if l == r {
			return i
		}
		l += v
	}
	return -1
}

func main() {
	a := []int{-1, 1, 1}
	res := pivotIndex(a)
	fmt.Println(res)
}
