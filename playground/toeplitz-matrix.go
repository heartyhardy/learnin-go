package main

import "fmt"

func isToeplitzMatrix(matrix [][]int) bool {
	r, c := len(matrix)-1, len(matrix[0])-1
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if matrix[i][j] != matrix[i+1][j+1] {
				return false
			}
		}
	}
	return true
}

func main() {
	matrix := [][]int{
		[]int{1, 2, 3, 4},
		[]int{5, 1, 2, 3},
		[]int{9, 5, 1, 2},
	}
	res := isToeplitzMatrix(matrix)
	fmt.Println(res)
}
