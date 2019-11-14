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

// 	for i := 1; i <= 12; i++ {
// 		fmt.Printf("%v X %v : %v \n", n_num, i, n_num*i)
// 	}
// }
