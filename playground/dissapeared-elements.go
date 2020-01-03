package main

import (
	"math"
)

// func findDisappearedNumbers(nums []int) []int {
// 	uq := make(map[int]int)
// 	res := make([]int, 0)
// 	for i := 0; i < len(nums); i++ {
// 		if _, ok := uq[nums[i]]; !ok {
// 			uq[nums[i]] = nums[i]
// 		}
// 	}
// 	for i := 1; i < len(nums)+1; i++ {
// 		if _, ok := uq[i]; !ok {
// 			res = append(res, i)
// 		}
// 	}
// 	return res
// }

// func findDisappearedNumbers(nums []int) []int {
// 	all := make([]int, len(nums)+1)
// 	res := make([]int, 0)
// 	for i := 0; i < len(nums); i++ {
// 		all[nums[i]] = nums[i]
// 	}

// 	for i := 1; i < len(all); i++ {
// 		if all[i] == 0 {
// 			res = append(res, i)
// 		}
// 	}
// 	return res
// }

func findDisappearedNumbersSmall(nums []int) []int {
	res := make([]int, 0)
	nsum := 0
	for i := 0; i < len(nums); i++ {
		nsum |= int(math.Pow(2, float64(nums[i])))
	}
	for i := 1; i < len(nums)+1; i++ {
		if nsum&int(math.Pow(2, float64(i))) == 0 {
			res = append(res, i)
		}
	}
	return res
}

func findDisappearedNumbers(nums []int) []int {

	l := len(nums)
	res := make([]int, 0)
	for _, v := range nums {
		nums[(v-1)%l] += l
	}
	for i, v := range nums {
		if v <= l {
			res = append(res, i+1)
		}
	}
	return res
}

// func main() {
// 	a := []int{4, 3, 2, 7, 8, 2, 3, 1}
// 	res := findDisappearedNumbers(a)
// 	fmt.Println(res)
// }
