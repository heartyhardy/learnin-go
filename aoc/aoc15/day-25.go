package aoc15

import "fmt"

const matrixSize = 10000

func createGrid() [][]int {
	grid := make([][]int, matrixSize+1)
	for r := 0; r <= matrixSize; r++ {
		grid[r] = make([]int, matrixSize+1)
		for c := 0; c <= matrixSize; c++ {
			if r == 0 || c == 0 {
				grid[r][c] = r + c
				continue
			}
		}
	}
	return grid
}

func generateCodes() int {
	i, v := 1, 20151125
	for r := 1; r <= matrixSize; r++ {
		for c := r; c >= 1; c-- {
			if r > 1 {
				v *= 252533
				v %= 33554393
			}
			//(*grid)[c][r-c+1] = v
			if c == 3010 && r-c+1 == 3019 {
				return v
			}
			i++
		}
	}
	return -1
}

func printGrid(grid *[][]int) {
	for r := 0; r <= matrixSize; r++ {
		for c := 0; c <= matrixSize; c++ {
			fmt.Printf("%4v", (*grid)[r][c])
		}
		fmt.Println()
	}
}

//Run Solution
func Run() {
	//grid := createGrid()
	code := generateCodes()
	//printGrid(&grid)
	fmt.Println("Part I: ", code)
}
