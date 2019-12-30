package main

import "fmt"

func missingNumber(nums []int) int {
	missing, i := 0, 0
	for ; i < len(nums); i++ {
		missing ^= i ^ nums[i]
		fmt.Println(missing, i, nums[i], i^nums[i])
	}
	return missing ^ i
}

// func main() {
// 	a := []int{3, 2, 1, 0, 5}
// 	res := missingNumber(a)
// 	fmt.Println(res, 0^3)
// }
