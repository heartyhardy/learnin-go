package main

import "fmt"

func removeElement(nums []int, val int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
			i = -1
		}
	}
	return len(nums)
}

func main() {
	a := []int{2, 2, 3, 3}
	res := removeElement(a, 2)
	fmt.Println("New Length: ", res)
}
