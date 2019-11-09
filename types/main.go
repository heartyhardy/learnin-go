package main

import "fmt"

func main() {
	// var i uint8 = 250
	// var j int8 = 1
	// fmt.Printf("%v, %T \n", i, i)
	// fmt.Println(i + uint8(j))

	var name string = "Charith Rathnayake"
	var bName []byte = []byte(name) // bName := []byte(name)

	fmt.Println(bName)
}
