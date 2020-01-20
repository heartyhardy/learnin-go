package main

func flipAndInvertImage(A [][]int) [][]int {
	for r, rows := range A {
		for c := range rows {
			if c < (len(rows)/2 + len(rows)%2) {
				A[r][len(rows)-1-c], A[r][c] = A[r][c], A[r][len(rows)-1-c]
			}

			A[r][c] = A[r][c] ^ 1
		}
	}
	return A
}

func main() {
	// [[1,1,0],[1,0,1],[0,0,0]]  [[1,0,0],[0,1,0],[1,1,1]]
	// [[1,1,0,0],[1,0,0,1],[0,1,1,1],[1,0,1,0]]   [[1,1,0,0],[0,1,1,0],[0,0,0,1],[1,0,1,0]]

	// a := [][]int{
	// 	[]int{1, 1, 0},
	// 	[]int{1, 0, 1},
	// 	[]int{0, 0, 0},
	// }
	a := [][]int{
		[]int{1, 1, 0, 0},
		[]int{1, 0, 0, 1},
		[]int{0, 1, 1, 1},
		[]int{1, 0, 1, 0},
	}
	flipAndInvertImage(a)
}
