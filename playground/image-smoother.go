package main

func imageSmoother(M [][]int) [][]int {
	// type xy struct {
	// 	x int
	// 	y int
	// }
	// xymap := map[uint8]xy{
	// 	0: {0, 0},
	// 	1: {-1, -1},
	// 	2: {0, -1},
	// 	3: {1, -1},
	// 	4: {-1, 0},
	// 	5: {1, 0},
	// 	6: {-1, 1},
	// 	7: {0, 1},
	// 	8: {1, 1},
	// }
	xymap := [][]int{
		[]int{0, 0},
		[]int{-1, -1},
		[]int{0, -1},
		[]int{1, -1},
		[]int{-1, 0},
		[]int{1, 0},
		[]int{-1, 1},
		[]int{0, 1},
		[]int{1, 1},
	}
	res := make([][]int, len(M))
	for i, rv := range M {
		res[i] = make([]int, len(rv))
	}

	rlen, clen := len(M), len(M[0])
	for r, rv := range M {
		for c := range rv {
			cells, cavg := 0, 0
			for _, v := range xymap {
				cx, cy := c+v[0], r+v[1]
				if cx >= 0 && cx < clen && cy >= 0 && cy < rlen {
					cavg += M[cy][cx]
					cells++
				}
			}
			res[r][c] = cavg / cells
		}
	}
	return res
}

// func main() {
// 	a := [][]int{
// 		{1, 1, 1},
// 		{1, 0, 1},
// 		{1, 1, 1},
// 	}
// 	res := imageSmoother(a)
// 	fmt.Println(res)
// }
