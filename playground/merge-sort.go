package main

import "fmt"

func mergeSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	m := len(nums) / 2
	return merge(mergeSort(nums[:m]), mergeSort(nums[m:]))
}

func merge(l, r []int) []int {
	size, i, j := len(l)+len(r), 0, 0
	res := make([]int, size, size)

	for k := 0; k < size; k++ {
		switch {
		case i == len(l):
			res[k] = r[j]
			j++
		case j == len(r):
			res[k] = l[i]
			i++
		case l[i] > r[j]:
			res[k] = r[j]
			j++
		case l[i] < r[j]:
			res[k] = l[i]
			i++
		case l[i] == r[j]:
			res[k] = r[j]
			j++
		}
	}
	return res
}

func main() {
	a := []int{1, 5, 3, 4, 6, 2}
	res := mergeSort(a)
	fmt.Println(res)
}
