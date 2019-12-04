package main

import "fmt"

func getRow(rowIndex int) []int {
	row := []int{1}
	for c := 1; c <= rowIndex; c++ {
		row = append(row, row[c-1]*(rowIndex-(c-1))/c)
	}
	return row
}

func main() {
	res := getRow(5)
	fmt.Println(res)
}
