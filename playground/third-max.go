package main

import (
	"fmt"
	"sort"
)

func thirdMax(nums []int) int {
	max := make(map[int]int)
	keys := []int{}
	for i := 0; i < len(nums); i++ {
		if _, ok := max[nums[i]]; !ok {
			max[nums[i]] = nums[i]
			keys = append(keys, nums[i])
		}
	}
	sort.Ints(keys)

	if len(max) < 3 {
		return keys[len(keys)-1]
	}
	keys = keys[len(keys)-3 : len(keys)-1]
	return keys[0]
}

func main() {
	a := []int{1, 0}
	res := thirdMax(a)
	fmt.Println("Zeroes moved: ", res)
}
