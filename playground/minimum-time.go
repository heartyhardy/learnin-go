package main

func minTimeToVisitAllPoints(points [][]int) int {
	var xdiff, ydiff, time int
	for i := 0; i < len(points)-1; i++ {
		if xdiff = points[i+1][0] - points[i][0]; xdiff < 0 {
			xdiff = -xdiff
		}
		if ydiff = points[i+1][1] - points[i][1]; ydiff < 0 {
			ydiff = -ydiff
		}

		switch {
		case xdiff > ydiff:
			time += xdiff
		case xdiff < ydiff:
			time += ydiff
		case xdiff == ydiff:
			time += xdiff
		}
	}
	return time
}

// func main() {
// 	//a := [][]int{[]int{1, 1}, []int{3, 4}, []int{-1, 0}}
// 	b := [][]int{[]int{3, 2}, []int{-2, 2}}

// 	time := minTimeToVisitAllPoints(b)
// 	fmt.Println(time)
// }
