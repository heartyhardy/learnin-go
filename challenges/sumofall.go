package main

import "fmt"

func sumofall(arr []int) int {
	sum := 0
	for i := 0; i < len(arr)/2; i++ {
		sum += arr[i] + arr[len(arr)-1-i]
	}
	if len(arr)%2 == 1 {
		sum += arr[len(arr)/2]
	}

	return sum
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	sum := sumofall(arr)

	fmt.Println(sum)
}
