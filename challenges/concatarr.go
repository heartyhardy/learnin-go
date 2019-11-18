// package main

// import (
// 	"fmt"
// )

// func concat(arr1 []int, arr2 []int) []int {
// 	newarr := []int{}
// 	maxlen := len(arr1) + len(arr2) - 1

// 	for i := 0; i <= maxlen; i++ {
// 		if i < len(arr1) {
// 			newarr = append(newarr, arr1[i])
// 		} else {
// 			newarr = append(newarr, arr2[i-len(arr1)])
// 		}
// 	}

// 	return newarr
// }

// func main() {
// 	arr1 := []int{1, 2, 3, 4}
// 	arr2 := []int{10, 20, 30, 40}

// 	newarr := concat(arr1, arr2)

// 	fmt.Println(newarr)
// }
