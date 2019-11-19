package main

import "fmt"

func rotate(arr []int, n int) {

	if n >= len(arr) {
		panic("n cannot be larger than the length of the given array!")
	} else if n < 0 {
		panic("n cannot be smaller than 0!")
	}

	whead := 0
	rhead := n

	w, r := 0, 0

	for whead <= len(arr)-1 {
		w = arr[whead]
		r = arr[rhead]

		arr[whead] = r
		arr[rhead] = w

		fmt.Printf("%v \n", arr)

		rhead++
		whead++

		if rhead == len(arr) {
			break
		}
	}

	return
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	rotate(arr, 2)

	//fmt.Println(arr)
}
