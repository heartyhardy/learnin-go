package main

import "fmt"

func findDisappearedNumbers(nums []int) []int {
	uq := make(map[int]int)
	res := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if _, ok := uq[nums[i]]; !ok {
			uq[nums[i]] = nums[i]
		}
	}
	for i := 1; i < len(nums)+1; i++ {
		if _, ok := uq[i]; !ok {
			res = append(res, i)
		}
	}
	return res
}

func main() {
	a := []int{1, 3, 3, 5, 5}
	res := findDisappearedNumbers(a)
	fmt.Println(res)
}
