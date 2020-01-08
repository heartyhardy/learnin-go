package main

func findUnsortedSubarray(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	max, start, size := nums[0], -4, 0

	for i, n := range nums {
		if n < max {
			if start == -4 {
				start = i - 1
			}
			for start >= 0 && nums[start] > n {
				start--
			}
			size = i - start
		} else {
			max = n
		}
	}
	return size
}

// func main() {
// 	a := []int{2, 1, 1, 1, 1}
// 	res := findUnsortedSubarray(a)
// 	fmt.Println(res)
// }
