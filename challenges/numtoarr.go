// package main

// import "fmt"

// func numtoarr(num uint64) []uint64 {

// 	cnum := num
// 	arr := []uint64{}
// 	var div, rest uint64 = 0, 0

// 	for cnum > 0 {
// 		div = cnum / 10
// 		rest = cnum - div*10
// 		arr = append(arr, rest)
// 		cnum = div
// 	}

// 	return arr

// }

// func main() {
// 	arr := numtoarr(2891123356778289070)
// 	fmt.Println(arr)
// }
