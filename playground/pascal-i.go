// package main

// import "fmt"

// func generate(numRows int) [][]int {
// 	result := [][]int{}
// 	master := [][]int{}

// 	if numRows < 1 {
// 		return result
// 	}

// 	row := []int{0, 1, 0}
// 	master = append(master, row)
// 	result = append(result, row[1:len(row)-1])
// 	r, c := 0, 0

// 	for numRows > 1 {
// 		row = make([]int, len(master[r])+1)
// 		row[0] = 0
// 		row[len(row)-1] = 0

// 		for c = 1; c < len(row)-1; c++ {
// 			row[c] = master[r][c] + master[r][c-1]
// 		}
// 		r++
// 		numRows--
// 		master = append(master, row)
// 		result = append(result, row[1:len(row)-1])
// 	}

// 	return result
// }

// func main() {
// 	res := generate(2)
// 	fmt.Println(res)
// }
