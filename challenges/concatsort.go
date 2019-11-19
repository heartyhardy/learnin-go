// package main

// import "fmt"

// func concat(arr1 []int, arr2 []int) []int {
// 	looping := true
// 	pos1, pos2 := 0, 0
// 	newarr := []int{}

// 	for looping {

// 		if pos1 == len(arr1) && pos2 < len(arr2) {
// 			newarr = append(newarr, arr2[pos2])
// 			pos2++
// 			continue
// 		} else if pos1 < len(arr1) && pos2 == len(arr2) {
// 			newarr = append(newarr, arr1[pos1])
// 			pos1++
// 			continue
// 		} else if pos1 == len(arr1) && pos2 == len(arr2) {
// 			looping = false
// 			continue
// 		}

// 		if arr1[pos1] < arr2[pos2] {
// 			newarr = append(newarr, arr1[pos1])
// 			pos1++
// 			continue
// 		} else {
// 			newarr = append(newarr, arr2[pos2])
// 			pos2++
// 			continue
// 		}
// 	}

// 	return newarr
// }

// func main() {
// 	arr1 := []int{2, 3, 11, 17}
// 	arr2 := []int{1, 4, 15, 17, 18, 22, 27}

// 	newarr := concat(arr1, arr2)

// 	fmt.Println(newarr)
// }
