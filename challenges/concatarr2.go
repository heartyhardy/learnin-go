// package main

// import "fmt"

// func concat(arr1 []int, arr2 []int) []int {
// 	maxlen := len(arr1) + len(arr2) - 1
// 	newarr := []int{}

// 	for i := 0; i < maxlen; i++ {
// 		if i < len(arr1) {
// 			newarr = append(newarr, arr1[i])
// 		}
// 		if i < len(arr2) {
// 			newarr = append(newarr, arr2[i])
// 		}
// 	}

// 	return newarr
// }

// func main() {
// 	arr1 := []int{1, 2, 3, 4, 5, 6}
// 	arr2 := []int{10, 20, 30, 40, 50, 60}

// 	newarr := concat(arr1, arr2)
// 	fmt.Println(newarr)
// }
