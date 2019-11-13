package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter some text here: ")
	text, err := reader.ReadString('\n')
	if err == nil {
		fmt.Println("Reading complete! \n")
	} else {
		fmt.Println("Error reading input! \n")
	}
	fmt.Printf("\nYou typed: %v", text)
}
