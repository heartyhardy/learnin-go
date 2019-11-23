// package main

// import "fmt"

// func main() {
// 	var w Writer = ConsoleWriter{
// 		data: []byte("Interfaces are awesome!"),
// 	}
// 	n, err := w.Write()
// 	fmt.Printf("%v bytes written. Errors: %v \n", n, err)
// }

// // Writer ...
// type Writer interface {
// 	Write() (int, error)
// }

// // ConsoleWriter ...
// type ConsoleWriter struct {
// 	data []byte
// }

// func (cw ConsoleWriter) Write() (int, error) {
// 	n, err := fmt.Println(string(cw.data))
// 	return n, err
// }
