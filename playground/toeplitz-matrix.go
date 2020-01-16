package main

import "fmt"

func isToeplitzMatrixNaive(matrix [][]int) bool {
	r, c := len(matrix)-1, len(matrix[0])-1
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			fmt.Printf(" %v %v ", matrix[i][j], matrix[i+1][j+1])
			if matrix[i][j] != matrix[i+1][j+1] {
				return false
			}
		}
	}
	return true
}

func isToeplitzMatrix(matrix [][]int) bool {
	r, c := len(matrix), len(matrix[0])
	lrow := matrix[0][:c]
	fmt.Println(lrow)
	for i := 1; i < r; i++ {
		for j := 1; j < c; j++ {
			fmt.Println(matrix[i][j], lrow[j-1])
			if matrix[i][j] != lrow[j-1] {
				return false
			}
		}
		lrow = matrix[i][:c]
	}
	return true
}

// 1, 2, 3, 4, 5, 1, 2, 3, 9, 5, 1, 2, 10, 9, 5, 1
//  _ 1, 2, 3, 4
//  _, 5, 1, 2, 3
//  _,
func main() {
	matrix := [][]int{
		[]int{1, 2, 3, 4},
		[]int{5, 1, 2, 3},
		[]int{9, 5, 1, 2},
		[]int{10, 9, 5, 1},
	}
	res := isToeplitzMatrix(matrix)
	fmt.Println(res)
}
