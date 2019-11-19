package main

import "fmt"

func rotate(arr []int, n int) {

	whead := 0
	rhead := n

	w, r := 0, 0

	for whead <= len(arr)-1 {
		w = arr[whead]
		r = arr[rhead]

		arr[whead] = r
		arr[rhead] = w

		rhead++

		if rhead == len(arr) {
			rhead = 0
		}

		whead++
	}

	return
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	rotate(arr, 2)

	fmt.Println(arr)
}
