package main

import "fmt"

func oddCells(n int, m int, indices [][]int) int {
	matrix := make([][]int, n)
	l := 0
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, m)
		for j := 0; j < m; j++ {
			l++
		}
	}
	return l
}

func main() {
	fmt.Println(oddCells(2, 3, nil))
}
