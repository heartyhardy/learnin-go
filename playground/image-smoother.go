package main

import "fmt"

func imageSmoother(M [][]int) [][]int {
	pos := 0
	tmp := make([]int, len(M)*len(M[0]))

	for _, r := range M {
		for _, c := range r {
			tmp[pos] = c
			pos++
		}
	}

	fmt.Println(tmp)
	return M
}

func main() {
	a := [][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}
	res := imageSmoother(a)
	fmt.Println(res)
}
