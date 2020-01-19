package main

func largeGroupPositions(S string) [][]int {
	res := make([][]int, 0)
	l, r, ub := 0, 0, len(S)-1
	for i := 0; i < ub; i++ {
		if S[i] != S[i+1] {
			if (r - l) >= 2 {
				res = append(res, []int{l, r})
			}
			l = i + 1
			r = l
		} else {
			r++
		}
		if (r-l) >= 2 && i == ub-1 {
			res = append(res, []int{l, r})
		}
	}
	return res
}

// func main() {
// 	//str := "aaaabcdddeeaabbbbcccdddd"
// 	str := "abbbaaaaddd"
// 	res := largeGroupPositions(str)
// 	fmt.Println(res)

// }
