package main

import "fmt"

// Related constants
// const (
// 	a = (iota + 1) * iota //iota = 0 + 1 * 0
// 	b                     //iota = 1 now so 1+1 * 1 = 2
// 	c                     //iota = 2 now so 2+1 * 2 = 6
// )

// const (
// 	a = 1 + (1 + iota) //iota = 1 + 1 + 0
// 	b                  //iota = 1 now so 1+1 *+2 = 3
// 	c                  //iota = 2 now so 2+1 +3= 4
// )

// const (
// 	_ = iota // tells the compiler to ignore this value
// 	a = 1 << (10 * iota)
// 	b
// 	c
// )

const (
	_ = 1 << iota // tells the compiler to ignore this value
	a
	b
	c
	d
)

func main() {
	// const i = 10 // If this was const i int32 this wont work
	// var a int16 = 255
	// // supports implicit conversions
	// fmt.Printf("%v, %T\n", i+a, i+a)
	// println((4|5)&4 == 4)
	// fmt.Println(a, b, c)
	var store byte = a | b
	fmt.Printf("Is a? %v \n", store&b == a)
}
