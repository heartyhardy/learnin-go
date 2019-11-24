// package main

// import "fmt"

// func main() {
// 	intCnt := IntCounter(10)
// 	var inc Incrementer = &intCnt

// 	for i := 0; i < 11; i++ {
// 		fmt.Printf("Current value of intCnt: %v \n", inc.Increment())
// 	}
// }

// type IntCounter int

// type Incrementer interface {
// 	Increment() int
// }

// func (inc *IntCounter) Increment() int {
// 	*inc++
// 	return int(*inc)
// }
