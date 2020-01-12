package main

func findShortestSubArray(nums []int) int {
	numcnt := make(map[int]int)
	lencnt := make(map[int][]int)
	deg, minlen := -1, 0
	for i, v := range nums {
		numcnt[v]++
		lencnt[v] = append(lencnt[v], i)
		if deg < numcnt[v] {
			deg = numcnt[v]
		}
	}
	for _, v := range lencnt {
		if len(v) == deg {
			if minlen == 0 || minlen > v[len(v)-1]-v[0] {
				minlen = v[len(v)-1] - v[0]
			}
		}
	}
	return minlen + 1
}

// func main() {
// 	a := []int{3, 3, 2, 3, 3, 3, 1, 4, 4, 4, 4, 4}
// 	res := findShortestSubArray(a)
// 	fmt.Println(res)
// }
