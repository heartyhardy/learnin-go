package main

import (
	"fmt"
)

func findPairs(nums []int, k int) int {
	res := 0
	tmp := make(map[int]int)
	for _, v := range nums {
		tmp[v]++
	}
	fmt.Println(tmp)
	for i, v := range tmp {
		_, ok := tmp[i+k]
		if k > 0 && ok || k == 0 && v >= 2 {
			res++
		}
	}
	return res
}

// func findPairsII(nums []int, k int) int {
// 	type kv struct {
// 		x int
// 		y int
// 	}
// 	cnt := 0
// 	pairs := make(map[kv]bool)

// 	for i := 0; i < len(nums); i++ {
// 		for j := i + 1; j < len(nums); j++ {
// 			kdiff := int(math.Abs(float64(nums[i]) - float64(nums[j])))
// 			if kdiff == k {
// 				minv := int(math.Min(float64(nums[i]), float64(nums[j])))
// 				maxv := int(math.Min(float64(nums[i]), float64(nums[j])))
// 				npair := kv{minv, maxv}

// 				if _, ok := pairs[npair]; !ok {
// 					pairs[npair] = true
// 					cnt++
// 				}
// 			}
// 		}
// 	}
// 	return cnt
// }

func findPairsII(nums []int, k int) int {
	res := 0
	tmp := make(map[int]int)
	for _, v := range nums {
		tmp[v]++
	}
	for i, v := range tmp {
		_, ok := tmp[k+i]
		if k > 0 && ok || k == 0 && v >= 2 {
			res++
		}
	}

	return res
}

// func main() {
// 	a := []int{1, 2, 3, 4, 5}
// 	res := findPairsII(a, 1)
// 	fmt.Println(res)
// }

// [1,1,1,2,1]
// 1
