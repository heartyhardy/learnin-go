package main

import (
	"fmt"
)

func checkPossibility(nums []int) bool {
	c, li := 0, 0
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			c++
			if c >= 2 {
				return false
			}
			li = i
		}
	}
	if c == 0 {
		return true
	}
	if c == 1 {
		if li == 1 || li == len(nums)-1 {
			return true
		}
		if nums[li+1] >= nums[li-1] || nums[li] >= nums[li-2] {
			return true
		}
	}
	return false
}

func main() {
	a := []int{4, 2, 3}
	res := checkPossibility(a)
	fmt.Println(res)
}
