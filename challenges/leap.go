package main

import (
	"fmt"
)

func main() {
	leapyears := []int{}

	for y := 2019; y < 2019+20; y++ {
		if (y%400 == 0 || y%100 != 0) && y%4 == 0 {
			leapyears = append(leapyears, y)
		}
	}

	fmt.Println(leapyears)
}
