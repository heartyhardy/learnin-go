// package main

// import "fmt"

// type fn func(int) int

// func foreach(arr []int, f fn) {
// 	for i := 0; i < len(arr); i++ {
// 		arr[i] = f(arr[i])
// 	}
// }

// func psqr(num int) int {
// 	return num * num
// }

// func main() {
// 	arr := []int{1, 2, 3}
// 	foreach(arr[:], psqr)

// 	fmt.Println(arr)
// }
