package main

func findMaxConsecutiveOnes(nums []int) int {
	last, cur := 0, 0
	for _, v := range nums {
		if cur+v > last {
			last = cur + v
		}
		cur = (cur + v) * v
	}
	return last
}

// func main() {
// 	a := []int{1, 1, 0, 1, 1, 1}
// 	res := findMaxConsecutiveOnes(a)
// 	fmt.Println(res)
// }
