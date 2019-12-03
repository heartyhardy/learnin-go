package main

// import (
// 	"fmt"
// )

// func searchInsert(nums []int, target int) int {
// 	l, r, m := 0, len(nums)-1, 0
// 	m = l + (r-l)/2

// 	// Early exits
// 	if target == 0 || len(nums) == 0 {
// 		return 0
// 	}

// 	// Early exits
// 	if target > nums[r] {
// 		return r + 1
// 	} else if target < nums[l] {
// 		return 0
// 	}

// 	for l <= r {

// 		m = l + (r-l)/2

// 		if nums[m] == target {
// 			return m
// 		} else if nums[m] > target {
// 			r = m - 1
// 		} else {
// 			l = m + 1
// 		}
// 	}

// 	// if nums[m] > target {
// 	// 	return m
// 	// } else {
// 	// 	return m + 1
// 	// }
// 	return l
// }

// func searchInsert2(nums []int, target int) int {
// 	if nums[0] >= target {
// 		return 0
// 	}
// 	for i := 1; i < len(nums); i++ {

// 		if nums[i] >= target {
// 			return i
// 		}
// 	}
// 	return len(nums)
// }

// func main() {
// 	// a := []int{2, 6, 8, 10, 12}
// 	// res := searchInsert(a, 13)
// 	// fmt.Println(res)

// 	a := []int{2, 6, 10, 12}
// 	res := searchInsert(a, 9)
// 	fmt.Println(res)
// }
