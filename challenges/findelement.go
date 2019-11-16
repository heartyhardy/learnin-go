// package main

// import (
// 	"fmt"
// )

// func find(el int, arr []int) (int, bool) {
// 	i := 0
// 	found := false
// 	for ; i <= len(arr)/2; i++ {
// 		if arr[i] == el || arr[len(arr)-1-i] == el {
// 			if arr[len(arr)-1-i] == el {
// 				i = len(arr) - 1 - i
// 			}
// 			found = true
// 			break
// 		}
// 	}

// 	return i, found
// }

// func main() {
// 	arr := []int{1, 5, 2, 3, 4, 8, 7, 1}
// 	el := 2

// 	pos, res := find(el, arr)
// 	fmt.Printf("Result: %v %v \n", res, pos)
// }
