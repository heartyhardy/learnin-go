package main

func dominantIndex(nums []int) int {
	mx, mi := -1, -1
	for i, v := range nums {
		if mx < v {
			mx = v
			mi = i
		}
	}
	for i, v := range nums {
		if mx < v*2 && i != mi && mx != v {
			return -1
		}
	}
	return mi
}

// func main() {
// 	a := []int{0, 0, 0, 0}
// 	res := dominantIndex(a)
// 	fmt.Println(res)
// }
