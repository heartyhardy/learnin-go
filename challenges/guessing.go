package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func prompt() int {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Guess the secret number: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSuffix(text, "\n")
	guess, _ := strconv.Atoi(text)

	return guess
}

func main() {

	rand.Seed(time.Now().UnixNano())
	rndnum := rand.Intn(100)

	looping := true
	tries := 0
	prevguess := 0

	for looping {

		guess := prompt()

		switch {
		case guess < rndnum:
			fmt.Println("Too Small!")
			if prevguess != guess {
				tries++
				prevguess = guess
			}
		case guess > rndnum:
			fmt.Println("Too Large!")
			if prevguess != guess {
				tries++
				prevguess = guess
			}
		case guess == rndnum:
			fmt.Println("You won!")
			looping = false
			break
		case guess == -1:
			looping = false
			break
		}

	}

	fmt.Printf("Tries: %v\n", tries)
}
