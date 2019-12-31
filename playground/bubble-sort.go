package main

import "fmt"

func bubblesort(nums []int) {
	swapped, passes, curpass := false, len(nums)-1, 0
	for {
		if nums[curpass] > nums[curpass+1] {
			tmp := nums[curpass]
			nums[curpass] = nums[curpass+1]
			nums[curpass+1] = tmp
			swapped = true
		}
		if curpass == passes-1 && swapped {
			passes--
			curpass = 0
			swapped = false
			continue
		} else if curpass == passes && !swapped {
			break
		}
		curpass++
	}
}

func main() {
	a := []int{1, 6, 3, 7, 2, 5, 4, 9, 8, 0}
	bubblesort(a)
	fmt.Println("Bubble sorted: ", a)
}
