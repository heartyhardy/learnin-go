package main

func bubblesort(nums []int) {
	swp := false
	l := len(nums) - 1
	for i := 0; i < l; i++ {
		if nums[i] > nums[i+1] {
			t := nums[i]
			nums[i] = nums[i+1]
			nums[i+1] = t
			swp = true
		}
		if i == l-1 && !swp {
			break
		} else if i == l-1 && swp {
			l--
			i = -1
			swp = false
			continue
		}
	}
}

func bubblesortDesc(nums []int) {
	swp := false
	l := len(nums) - 1
	for i := 0; i < l; i++ {
		if nums[i] < nums[i+1] {
			t := nums[i]
			nums[i] = nums[i+1]
			nums[i+1] = t
			swp = true
		}
		if i == l-1 && !swp {
			break
		} else if i == l-1 && swp {
			l--
			i = -1
			swp = false
			continue
		}
	}
}

// func main() {
// 	a := []int{1, 6, 3, 7, 2, 5, 4, 9, 8, 0}
// 	bubblesort(a)
// 	fmt.Println("Bubble sorted: ", a)

// 	a = []int{1, 6, 3, 7, 2, 5, 4, 9, 8, 0}
// 	bubblesortDesc(a)
// 	fmt.Println("Bubblesort Desc: ", a)
// }
