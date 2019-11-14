package main

import "fmt"

func getArr() [20]bool {
	arr := [20]bool{}
	for i := 0; i < len(arr); i++ {
		arr[i] = true
	}
	return arr
}

func main() {
	arr := getArr()

	for i := 2; i < len(arr); i++ {
		if arr[i] == true {
			for x := i * i; x < len(arr); x += i {
				arr[x] = false
			}
		}
	}

	fmt.Println(arr)
}
