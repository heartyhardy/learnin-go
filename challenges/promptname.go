// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strings"
// )

// func main() {
// 	Part I
// 	reader := bufio.NewReader(os.Stdin)
// 	fmt.Print("Enter name: ")
// 	text, err := reader.ReadString('\n')
// 	if err == nil {
// 		fmt.Println("Reading complete! \n")
// 	} else {
// 		fmt.Println("Error reading input! \n")
// 	}
// 	fmt.Printf("\nHello %v", text)

// 	// Part II
// 	reader := bufio.NewReader(os.Stdin)
// 	fmt.Print("Enter name: ")
// 	text, _ := reader.ReadString('\n')
// 	switch strings.ToLower(text) {
// 	case "bob\n":
// 		fmt.Printf("\nHello there %v", text)
// 	case "alice\n":
// 		fmt.Printf("\nHowdy %v", text)

// 	}

// }
