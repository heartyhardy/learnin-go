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
// 	fmt.Print("Enter a number (n): ")

// 	n, _ := reader.ReadString('\n')
// 	n = strings.TrimSuffix(n, "\n")
// 	n_num, _ := strconv.Atoi(n)

// 	fmt.Print("Enter operation 1 - SUM | 2 - PRODUCT : ")

// 	ops, _ := reader.ReadString('\n')
// 	ops = strings.TrimSuffix(ops, "\n")
// 	ops_num, _ := strconv.Atoi(ops)

// 	sum := 0
// 	for i := 0; i <= n_num; i++ {
// 		switch ops_num {
// 		case 1:
// 			sum += i
// 		case 2:
// 			if sum == 0 {
// 				sum = 1
// 			}
// 			sum *= i
// 		}
// 	}

// 	fmt.Printf("Result of the selected operation is: %v \n", sum)
// }
