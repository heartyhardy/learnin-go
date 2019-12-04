package main

// import "fmt"

// type line struct {
// 	x1 float64
// 	x2 float64
// 	y1 float64
// 	y2 float64
// }

// func isIntersecting(l1 line, l2 line) (bool, float64, float64) {

// 	var ix, iy float64
// 	var insx, insy bool = false, false

// 	xdiff1 := l1.x1 - l1.x2
// 	ydiff1 := l1.y1 - l1.y2

// 	xdiff2 := l2.x1 - l2.x2
// 	ydiff2 := l2.y1 - l2.y2

// 	if xdiff1 == 0 {
// 		ix = l1.x1
// 	} else if ydiff1 == 0 {
// 		iy = l1.y1
// 	}

// 	if xdiff2 == 0 {
// 		ix = l2.x1
// 	} else if ydiff2 == 0 {
// 		iy = l2.y1
// 	}

// 	if ix != l1.x1 && ix != l1.x2 {
// 		if ix > l2.x1 && ix < l2.x2 {
// 			insx = true
// 		}
// 	} else if ix != l2.x1 && ix != l2.x2 {
// 		if ix > l1.x1 && ix < l1.x2 {
// 			insx = true
// 		}
// 	}

// 	if iy != l1.y1 && ix != l1.y2 {
// 		if iy > l2.y1 && iy < l2.y2 {
// 			insy = true
// 		}
// 	} else if iy != l2.y1 && ix != l2.y2 {
// 		if iy > l1.y1 && iy < l1.y2 {
// 			insy = true
// 		}
// 	}

// 	return insx && insy, ix, iy
// }

// func main() {
// 	l1 := line{3, 8, 5, 5}
// 	l2 := line{6, 6, 7, 3}
// 	isx, x, y := isIntersecting(l1, l2)
// 	fmt.Println(isx, x, y)
// }
