package main

import "fmt"

func bubblesort20200116(nums []int) []int {
	swapped := false
	i, imax, swpn := 0, len(nums)-1, 0
	steps := 0
	for i < imax {
		steps++
		if nums[i] > nums[i+1] {
			nums[i], nums[i+1] = nums[i+1], nums[i]
			swapped = true
			swpn = i
		}
		if i == imax-1 && !swapped {
			break
		} else if i == imax-1 && swapped {
			i = -1
			swapped = false
			imax = swpn
		}
		i++
	}
	fmt.Println(steps)
	return nums
}

func main() {
	nums := []int{1, 5, 3, 2, 7, 5, 7, 9, 6, 11, 10}
	res := bubblesort20200116(nums)
	fmt.Println(res)
}
