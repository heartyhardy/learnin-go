package main

import "fmt"

func insertionSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		k := nums[i]
		j := i - 1

		for j >= 0 && k < nums[j] {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = k
	}
}

// func main() {
// 	a := []int{3, 1, 2, 6, 5}
// 	//a := []int{16, 19, 4, 2, 20, 12, 13, 9, 3, 7}
// 	insertionSort(a)
// 	fmt.Println(a)
// }
