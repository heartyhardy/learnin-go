package main

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

func bubblesortDesc(nums []int) {
	i, l, swp := 0, len(nums)-1, false
	for {
		if nums[i] < nums[i+1] {
			t := nums[i]
			nums[i] = nums[i+1]
			nums[i+1] = t
			swp = true
		}
		if i == l-1 && swp {
			l--
			i = 0
			swp = false
			continue
		} else if i == l && !swp {
			break
		}
		i++
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
