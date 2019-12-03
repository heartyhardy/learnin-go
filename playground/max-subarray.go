package main

// func maxSubArray(nums []int) int {
// 	var max, min, last int = 0, 0, 0
// 	var haszero = false

// 	if len(nums) == 0 {
// 		return -2147483648
// 	}

// 	for _, n := range nums {

// 		if n == 0 {
// 			haszero = true
// 		}

// 		max += n
// 		if max < 0 {
// 			max = 0
// 			if n > min || min == 0 {
// 				min = n
// 			}
// 		}
// 		if max > last {
// 			last = max
// 		}
// 	}

// 	if last == 0 && !haszero {
// 		return min
// 	} else {

// 		return last
// 	}
// }

func maxSubArray(nums []int) int {
	var max, cur int = -2147483647, 0
	for _, n := range nums {
		cur += n
		if n > cur {
			cur = n
		}
		if cur > max {
			max = cur
		}
	}
	return max
}

// func main() {
// 	a := []int{-3, -1, 2}
// 	fmt.Println(maxSubArray(a))
// }
