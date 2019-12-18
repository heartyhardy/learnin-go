package main

import (
	"fmt"
)

// func containsDuplicate(nums []int) bool {
// 	sort.Ints(nums)
// 	for i := 0; i < len(nums)-1; i++ {
// 		if nums[i]^nums[i+1] == 0 {
// 			return true
// 		}
// 	}
// 	return false
// }

func containsDuplicate(nums []int) bool {
	dups := map[int]bool{}
	for i := 0; i < len(nums); i++ {
		if _, ok := dups[nums[i]]; !ok {
			dups[nums[i]] = true
		} else {
			return true
		}
	}
	return false
}

func main() {
	a := []int{2, 3, 2}
	res := containsDuplicate(a)
	fmt.Println(res)
}
