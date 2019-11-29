// Array is sorted therefore:
//	Sum of Leftmost two ints gives lower bound.
//	Sum of Rightmost two ints gives upper bound.
//	If sample array = [1, 2, 4, 5, 7, 8]
//	upperbound = 15	lowebound = 3
//	if SUM = 20 => doent exist
//	if SUM = 2 => doesnt exist
package main

import "fmt"

func sumoftwosorted(arr []int, sum int) []int {
	l := 0
	r := len(arr) - 1
	res := []int{}

	if arr[l]+arr[l+1] > sum || arr[r]+arr[r-1] < sum {
		return res
	}

	for l < r {
		switch {
		case arr[l]+arr[r] > sum:
			r--
		case arr[l]+arr[r] < sum:
			l++
		case arr[l]+arr[r] == sum:
			l++
			r++
			res = append(res, l, r)
			return res
		}
	}
	return res
}

func main() {
	a := []int{1, 3, 4, 5, 6, 7, 10}
	res := sumoftwosorted(a, 17)
	fmt.Println(res)
}
