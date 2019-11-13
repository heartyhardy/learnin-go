package main

import "fmt"

type person struct {
	name       string
	age        int
	contactnos []int
}

// if Exporting
// type Person struct {
// 	Name       string
// 	Age        int
// 	Contactnos []int
// }

// Note that structs are passed by value and a new copy is made everyime you use
// for this use & operator

func main() {
	// Non positional
	personx := person{ // non positional syntax
		name: "Charith",
		age:  38,
		contactnos: []int{
			332222355,
		},
	}

	// Positional syntax
	persony := person{
		"Max",
		67,
		[]int{
			55577888,
		},
	}

	// Anonymous struct
	details := struct{ birthday string }{birthday: "2233-10-12"}

	fmt.Println(personx)
	fmt.Println(persony)
	fmt.Println(details)
}
