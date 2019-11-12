package main

import "fmt"

func getX() int {
	return 2
}

func main() {

	if true {
		fmt.Println("true that!")
	}

	if x := getX(); x > 5 {
		fmt.Println("X > 5")
	} else {
		fmt.Println("x < 5")
	}

	// Tagged
	switch x := getX(); x {
	case 10:
		fmt.Println("10 selected")
	case 5:
		fmt.Println("5 selected")
	case 1, 2:
		fmt.Println("1 or 2 selected")
	default:
		fmt.Println("someting else selected")
	}
	// Tagless
	// Can overlap for instance x > 5 , x > 2 can overlap, 1st case is selected
	x := getX()
	switch {
	case x > 0 && x < 3:
		fmt.Println("1-2 range")
		fallthrough // This keyword allows to fall through to next case
	case x > 2 && x < 6:
		fmt.Println("3-5 range")
	case x > 5:
		fmt.Println("6 onwards...")
	}

	// Type switching

	var j interface{} = 10.5
	switch j.(type) {
	case int:
		fmt.Println("Int")
	case float64:
		fmt.Println("Float")
	default:
		fmt.Println("Unknown")
	}

}
