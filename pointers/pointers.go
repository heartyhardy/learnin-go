package main

import "fmt"

func main() {
	a := 42
	b := 33
	fmt.Println(a, b)

	var c *int = &a
	*c = 50
	fmt.Println(a, *c)

	d := &c
	fmt.Println(**d)

	var me *someStruct
	me = &someStruct{name: "Charith", age: 30}
	fmt.Println(*me)

	// Object initialization cannot be used with NEW keyword
	var me2 *someStruct
	me2 = new(someStruct)
	fmt.Println(*me2)

	// If not assigned will be initialized with <nil>
	var me3 *someStruct
	fmt.Println(me3)

	// No need to use deref operator when accessing fields using a pointer
	var me4 *someStruct
	me4 = &someStruct{name: "Max", age: 35}
	me4.age = 40
	fmt.Println(me4.name, me4.age)

}

type someStruct struct {
	name string
	age  int
}
