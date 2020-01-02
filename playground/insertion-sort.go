package main

import "fmt"

func insertionSort(nums []int) {
	cnt := 0
	for i := 1; i < len(nums); i++ {
		cnt++
		k := nums[i]
		j := i - 1
		for j >= 0 && k < nums[j] {
			cnt++
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = k
	}
	fmt.Println("Steps: ", cnt)
}

func main() {
	a := []int{3, 1, 2, 6, 5}
	//a := []int{16, 19, 4, 2, 20, 12, 13, 9, 3, 7}
	insertionSort(a)
	fmt.Println(a)
}
