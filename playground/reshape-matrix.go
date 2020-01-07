package main

import "fmt"

func matrixReshapeII(nums [][]int, r int, c int) [][]int {
	cc := len(nums[0])
	cr := len(nums)
	res := make([][]int, r)
	tmp := make([]int, 0)
	pos := 0

	if r*c != cc*cr || r == cr && c == cc {
		return nums
	}
	for i := range res {
		res[i] = make([]int, c)
	}
	for i := 0; i < cr; i++ {
		for j := 0; j < cc; j++ {
			tmp = append(tmp, nums[i][j])
		}
	}
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			res[i][j] = tmp[pos]
			pos++
		}
	}
	return res
}

func matrixReshape(nums [][]int, r int, c int) [][]int {
	cr, cc := len(nums), len(nums[0])
	res := make([][]int, r)

	if r*c != cc*cr || r == cr && c == cc {
		return nums
	}
	for i := range res {
		res[i] = make([]int, c)
	}
	for i := 0; i < r*c; i++ {
		res[i/c][i%c] = nums[i/cc][i%cc]
	}
	return res
}

func main() {
	a := [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}
	res := matrixReshape(a, 2, 4)
	fmt.Println(res)
}
