// package main

// import "fmt"

// func removeDuplicates(nums []int) int {

// 	if len(nums) < 2 {
// 		return len(nums)
// 	}

// 	i := 0
// 	done := false

// 	for done == false {
// 		if nums[i]^nums[i+1] == 0 {
// 			nums = append(nums[:i+1], nums[i+2:]...)
// 			if i-1 > 0 {
// 				i--
// 			}
// 		} else {
// 			i++
// 			done = i >= len(nums)-1 || len(nums) == 0
// 		}
// 		done = i >= len(nums)-1 || len(nums) == 0
// 	}
// 	return len(nums)
// }

// func main() {
// 	// var a []int = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
// 	var a []int = []int{1, 1, 2, 3, 3, 4}
// 	res := removeDuplicates(a[:])
// 	fmt.Println(res)
// }
