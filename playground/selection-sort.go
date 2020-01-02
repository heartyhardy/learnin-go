package main

func selectionSort(nums []int) {
	l, c, swp := 0, 0, false
	for i := 0; i < len(nums)-1; i++ {
		l, c = i, i
		for j := i + 1; j < len(nums); j++ {
			if nums[c] > nums[j] {
				c = j
				swp = true
			}
		}
		if swp {
			nums[l], nums[c] = nums[c], nums[l]
			swp = false
		}
	}
}

// func main() {
// 	a := []int{1, -5, 3, 2, 7, -8, 6}
// 	selectionSort(a)
// 	fmt.Println(a)
// }
