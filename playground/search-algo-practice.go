package main

import "fmt"

func linearSearch(nums []int, src int) int {
	m, l, r := 0, 0, len(nums)-1
	for l <= r {
		m = l + (r-l)/2
		switch {
		case src == nums[m]:
			return m
		case src < nums[m]:
			r = m - 1
		case src > nums[m]:
			l = m + 1
		}
	}
	return -1
}

func main() {
	a := []int{1, 2, 2, 5, 6, 6, 10, 15}
	res := linearSearch(a, 110)
	fmt.Println(res)
}
