package main

import "fmt"

func thirdMax(nums []int) int {
	swp := false
	res := map[int]int{}
	pos, third, l := 0, 0, len(nums)-1
	for i := 0; i < l; i++ {
		if nums[i] < nums[i+1] {
			nums[i], nums[i+1] = nums[i+1], nums[i]
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

	for i := 0; i < len(nums) && pos < 3; i++ {
		if _, ok := res[nums[i]]; !ok {
			res[nums[i]] = nums[i]
			third = nums[i]
			pos++
		}
	}

	if len(res) <= 2 {
		for _, v := range res {
			if third < v {
				third = v
			}
		}
	}
	return third
}

func main() {
	a := []int{1, -1, 3, -2, 3}
	res := thirdMax(a)
	fmt.Println("Zeroes moved: ", res)
}
