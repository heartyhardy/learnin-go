// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// func main() {
// 	reader := bufio.NewReader(os.Stdin)
// 	fmt.Print("Enter a number: ")
// 	text, _ := reader.ReadString('\n')

// 	text = strings.TrimSuffix(text, "\n")

// 	num, err := strconv.Atoi(text)

// 	if err != nil {
// 		fmt.Println("Invalid Input!")
// 	} else {
// 		sum := 0
// 		for i := 0; i <= num; i++ {
// 			if i%3 == 0 || i%5 == 0 {
// 				sum += i
// 			}
// 		}

// 		fmt.Printf("\nSum of numbers from 1 to %v is %v \n", num, sum)
// 	}
// }
