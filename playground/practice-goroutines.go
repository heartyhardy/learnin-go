package main

import "fmt"

// r -  last index of swap position
// swp - if swap didn't happen at the end of a pass, break out

func bbsort(nums []int, c chan<- []int) {
	if len(nums) <= 1 {
		c <- nums
		return
	}
	size, r, i, swp := len(nums)-1, len(nums)-1, 0, false
	for r >= 0 {
		if nums[i] > nums[i+1] {
			nums[i], nums[i+1], swp, r = nums[i+1], nums[i], true, i
		}
		if i == size-1 && !swp {
			c <- nums
			return
		} else if i == size-1 && swp {
			swp, i = false, 0
			continue
		}
		i++
	}
	c <- nums
}

func main() {
	in := make(chan []int)
	nums1 := []int{187, 52, 85, 92, 58, 182, 30, 26, 77, 200, 159, 62, 123, 7, 96, 170, 149, 57, 108, 132, 183, 114, 180, 131, 32}
	nums2 := []int{22, 91, 122, 68, 153, 103, 37, 72, 152, 189, 83, 197, 176, 23, 134, 161, 40, 4, 58, 16, 0, 48, 33, 54, 181}
	go bbsort(nums1, in)
	go bbsort(nums2, in)
	res := <-in
	fmt.Printf("%-3v  \n", res)
	res = <-in
	fmt.Printf("%-3v  \n", res)
	fmt.Scanln()
}
