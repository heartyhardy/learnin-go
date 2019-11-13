package main

import (
	"fmt"
)

// Note that maps are passed by reference

func main() {
	// Order in a map is not guaranteed
	//var maps = make(map[string]int)
	var states = map[string]int{
		"Gampaha": 100,
	}
	// Using array as key, slices dont work
	another := map[[3]int]string{}
	fmt.Printf("%v , %T\n", states, states)
	fmt.Printf("%v , %T\n", another, another)
	fmt.Println(states["Gampaha"])

	states["Colombo"] = 200
	fmt.Println(states)

	delete(states, "Gampaha")
	fmt.Println(states)

	pick, ok := states["Gampaha"]
	fmt.Printf("%v, %v\n", pick, ok)

	_, err := states["Gampaha"]
	fmt.Println(err)
}
