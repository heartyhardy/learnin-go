// package main

// import (
// 	"fmt"
// )

// func oddCells(n int, m int, indices [][]int) int {

// 	if n == 0 || m == 0 || len(indices[0]) < 1 {

// 		return 0
// 	}

// 	odds := 0
// 	matrix := make([][]int, n)

// 	for i := 0; i < n; i++ { // O(n)
// 		matrix[i] = make([]int, m)
// 		for j := 0; j < m; j++ {
// 			matrix[i][j] = 0
// 		}
// 	}

// 	for i := 0; i < len(indices); i++ {
// 		for k := 0; k < n; k++ { // O(n)
// 			for j := 0; j < m; j++ {
// 				if k == indices[i][0] {
// 					matrix[k][j] = matrix[k][j] + 1
// 				}
// 				if j == indices[i][1] {
// 					matrix[k][j] = matrix[k][j] + 1
// 				}
// 			}
// 		}
// 	}

// 	for k := 0; k < n; k++ { // O(n)
// 		for j := 0; j < m; j++ {
// 			if matrix[k][j]%2 != 0 {
// 				odds++
// 			}
// 		}
// 	}

// 	return odds
// }

// func main() {
// 	a := [][]int{[]int{}, []int{}}
// 	fmt.Println(oddCells(2, 3, a))
// }
