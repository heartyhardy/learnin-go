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

func main() {
	a := []int{1, 3, 1, 5, 4}
	res := findPairs(a, 0)
	fmt.Println(res)
}

// [1,1,1,2,1]
// 1
