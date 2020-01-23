package main

import (
	"fmt"
	"time"
)

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

func bbsortStr(strs []string, c chan<- []string) {
	if len(strs) < 2 {
		c <- strs
		return
	}
	size, r, i, swp := len(strs)-1, len(strs)-1, 0, false
	for r >= 0 {
		if strs[i] > strs[i+1] {
			strs[i], strs[i+1], swp, r = strs[i+1], strs[i], true, i
		}
		if i == size-1 && !swp {
			c <- strs
			return
		} else if i == size-1 && swp {
			swp, i = false, 0
			continue
		}
		i++
	}
}

func selsortAnimated(nums []int) {
	min, swp := 0, false
	fmt.Print("\033[2J \n")
	for i := 0; i < len(nums); i++ {
		min = i
		highlight(nums, min, 0)
		clearDelayed(1000)
		for j := i + 1; j < len(nums); j++ {
			if nums[min] > nums[j] {
				highlight(nums, j, 1)
				clearDelayed(1000)
				min = j
				swp = true
			} else {
				highlight(nums, j, 3)
				clearDelayed(500)
			}
		}
		if swp {
			highlight(nums, i, 2)
			clearDelayed(1500)
			highlight(nums, min, 2)
			clearDelayed(1500)
		}
		nums[i], nums[min] = nums[min], nums[i]
		if swp {
			highlight(nums, i, 2)
			clearDelayed(1500)
			highlight(nums, min, 2)
			clearDelayed(1500)
			swp = false
		} else {
			highlight(nums, i, 0)
			clearDelayed(500)
		}
	}
}

func clearDelayed(t time.Duration) {
	time.Sleep(t * time.Millisecond)
	fmt.Print("\033[2J \n")
}

func highlight(nums []int, index int, htype int) {
	for i, v := range nums {
		if i == index {
			switch htype {
			case 0:
				fmt.Printf(" [%1v] ", v)
			case 1:
				fmt.Printf(" \u001b[34m[%1v]\u001b[0m ", v)
			case 2:
				fmt.Printf(" \u001b[31m[%1v]\u001b[0m ", v)
			case 3:
				fmt.Printf(" \u001b[33m[%1v]\u001b[0m ", v)
			}
			continue
		}
		fmt.Printf(" %1v ", v)

	}
}

func selsort(nums []int, c chan<- []int) {
	min := 0
	for i := 0; i < len(nums); i++ {
		min = i
		for j := i + 1; j < len(nums); j++ {
			if nums[min] > nums[j] {
				min = j
			}
		}
		nums[i], nums[min] = nums[min], nums[i]
	}
	c <- nums
}

func main() {
	// in := make(chan []int)
	// nums1 := []int{187, 52, 85, 92, 58, 182, 30, 26, 77, 200, 159, 62, 123, 7, 96, 170, 149, 57, 108, 132, 183, 114, 180, 131, 32}
	// nums2 := []int{22, 91, 122, 68, 153, 103, 37, 72, 152, 189, 83, 197, 176, 23, 134, 161, 40, 4, 58, 16, 0, 48, 33, 54, 181}
	// go bbsort(nums1, in)
	// go bbsort(nums2, in)
	// res := <-in
	// fmt.Printf("%-3v  \n", res)
	// res = <-in
	// fmt.Printf("%-3v  \n", res)
	// fmt.Scanln()

	// in := make(chan []string)
	// str1 := `c p a b l e bb sd ca ac oi kj wd po`
	// str2 := `🍉 🥞 🍨 🍷 🍔 🍟`
	// strs1 := strings.Fields(str1)
	// strs2 := strings.Fields(str2)
	// go bbsortStr(strs1, in)
	// go bbsortStr(strs2, in)
	// res := <-in
	// fmt.Printf(" %4v \n", res)
	// res = <-in
	// fmt.Printf(" %4v  \n", res)

	// in := make(chan []int)
	// nums1 := []int{187, 52, 85, 92, 58, 182, 30, 26, 77, 200, 159, 62, 123, 7, 96, 170, 149, 57, 108, 132, 183, 114, 180, 131, 32}
	// nums2 := []int{22, 91, 122, 68, 153, 103, 37, 72, 152, 189, 83, 197, 176, 23, 134, 161, 40, 4, 58, 16, 0, 48, 33, 54, 181}
	// go selsort(nums1, in)
	// go selsort(nums2, in)
	// res := <-in
	// fmt.Printf(" %2v \n", res)
	// res = <-in
	// fmt.Printf(" %2v \n", res)

	nums1 := []int{187, 52, 85, 92, 58, 182, 30, 26, 77, 200, 159, 62, 123, 7, 96, 170, 149, 57, 108, 132, 183, 114, 180, 131, 32}
	selsortAnimated(nums1)
}
