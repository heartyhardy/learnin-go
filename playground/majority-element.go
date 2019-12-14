package main

// import "fmt"

// func majorityElement(nums []int) int {
// 	nummap := map[int]int{}
// 	maxval, max := 0, 0
// 	for _, num := range nums {
// 		if _, ok := nummap[num]; !ok {
// 			nummap[num] = 1
// 		} else {
// 			nummap[num] = nummap[num] + 1
// 		}
// 		if maxval < nummap[num] {
// 			maxval = nummap[num]
// 			max = num
// 		}
// 	}
// 	return max
// }

// func main() {
// 	a := []int{0, 3, 3, 2, 1, 3, 3, 3, 3, 1, 1, 2, 2, 2, 2}
// 	res := majorityElement(a)
// 	fmt.Println(res)
// }
