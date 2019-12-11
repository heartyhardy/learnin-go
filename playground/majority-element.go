package main

import "fmt"

func majorityElement(nums []int) int {
	nummap := map[int]int{}
	for _, num := range nums {
		if _, ok := nummap[num]; !ok {
			nummap[num] = 1
		} else {
			nummap[num] = nummap[num] + 1
		}
	}

	maxval, max := 0, 0
	for k, v := range nummap {
		if maxval < v {
			maxval = v
			max = k
		}
	}
	return max
}

func main() {
	a := []int{0, -1, -1, 0}
	res := majorityElement(a)
	fmt.Println(res)
}
