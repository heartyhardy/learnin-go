// package main

// import "fmt"

// // Function that returns the sum of multiple values
// func sum(values ...int) int {
// 	res := 0
// 	for _, v := range values {
// 		res += v
// 	}
// 	return res
// }

// // Function that calculates product and returns a pointer
// func product(values ...int) *int {
// 	res := 1
// 	for _, v := range values {
// 		res *= v
// 	}
// 	return &res
// }

// // Function that has a named return
// func div2(divisor int, values ...int) (result int) {

// 	result = 0
// 	for _, v := range values {
// 		result += v / divisor
// 	}
// 	return

// }

// func division(a, b int) (int, error) {
// 	if b == 0 {
// 		return 0, fmt.Errorf("Division by zero!")
// 	}
// 	return a / b, nil
// }

// func main() {
// 	fmt.Println(sum(1, 2, 4, 5, 10))

// 	prod := product(2, 3, 7)
// 	*prod *= 10
// 	fmt.Println("Product =", *prod)

// 	div2res := div2(2, 14, 8, 4, 2)
// 	fmt.Println("Div =", div2res)

// 	divres, err := division(10, 10)
// 	fmt.Println(divres, err)

// 	// IIFE function
// 	func() {
// 		msg := "Hello from IIFE"
// 		fmt.Println(msg)
// 	}()

// 	// When IiFE inside a loop and executed async then better to pass variables
// 	// into the function like this...
// 	values := [...]int{1, 2, 5, 6, 7}
// 	for _, i := range values {
// 		func(x int) {
// 			fmt.Println(x * x)
// 		}(i)
// 	}

// 	// can be defined as this..
// 	var f func() = func() {
// 		fmt.Println("Fn as var")
// 	}
// 	f()

// 	// Invoking a method on greeter
// 	g := greeter{
// 		msg:   "Charith",
// 		times: 5,
// 	}
// 	g.greet()

// 	// Defining and using custom type intx
// 	x := intx(10)
// 	fmt.Println(x.increment())
// }

// type greeter struct {
// 	msg   string
// 	times int
// }

// // Method defined on greeter struct
// // g is passed as a value, make it a pointer to pass by ref
// func (g greeter) greet() {
// 	for i := 0; i < g.times; i++ {
// 		fmt.Println(g.msg, "greeted", i+1, "times!")
// 	}
// }

// // Custom type
// type intx int

// // Method defined on the custom type
// func (i intx) increment() intx {
// 	i++
// 	return i
// }
