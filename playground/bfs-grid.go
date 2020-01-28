package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func loadGridFromFile(filename string) []string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic("Error occured while loading grid data!")
	}
	content := strings.Split(string(data), "\n")
	return content
}

func buildGrid(content []string) [][]uint8 {
	grid := make([][]uint8, len(content))
	for rn, fields := range content {
		rows := strings.Fields(fields)
		grid[rn] = make([]uint8, len(rows))
		for cn, col := range rows {
			uv, err := strconv.ParseUint(col, 10, 8)
			if err != nil {
				panic("Invalid data found in grid!")
			}
			grid[rn][cn] = uint8(uv)
		}
	}
	return grid
}

func printGrid(grid [][]uint8) {
	for _, rows := range grid {
		for _, col := range rows {
			switch col {
			case 0:
				fmt.Printf("\u001b[41m %v \u001b[0m", col)
			case 1:
				fmt.Printf("\u001b[44m %v \u001b[0m", col)
			}
		}
		fmt.Println()
	}
}

func main() {
	content := loadGridFromFile("grid.txt")
	grid := buildGrid(content)
	printGrid(grid)
}
