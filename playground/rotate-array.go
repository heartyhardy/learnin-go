package main

// func rotate(nums []int, k int) {
// 	k %= len(nums)
// 	t := nums[len(nums)-k:]
// 	copy(nums, append(t, nums[:len(nums)-k]...))
// }

func rotate(nums []int, k int) {
	k %= len(nums)
	reverse := func(l int, r int) {
		for l < r {
			t := nums[l]
			nums[l] = nums[r]
			nums[r] = t
			l++
			r--
		}
	}

	l := 0
	r := len(nums) - 1
	reverse(l, r)

	l = 0
	r = k - 1
	reverse(l, r)

	r = len(nums) - 1
	l = k
	reverse(l, r)
}

// func main() {
// 	nums := []int{2, 1}
// 	// 5 4 3 2 1
// 	// 3 4 5 2 1
// 	// 3 4 5 1 2

// 	// 7, 2, 3, 1, 5, 6, 4
// 	// 6, 2, 7, 1, 5, 3, 4
// 	// 5, 6, 7, 1, 2, 3, 4
// 	rotate(nums, 2)
// 	fmt.Println(nums)
// }
