package main

import "fmt"

func rotate(nums []int, k int) {
	var r int
	// if k < len(nums) {
	// 	r = len(nums) - k
	// } else {
	// 	r = k % len(nums)
	// }

	r = len(nums) - 1*(len(nums)-k) - 1
	for i := 0; i < len(nums)-1; i++ {
		t := nums[i]
		nums[i] = nums[i+1]
		nums[i+1] = t
		r--
		if i == len(nums)-2 {
			i = 0
		}
		if r == 0 {
			break
		}
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 2)
	fmt.Println(nums)
}
