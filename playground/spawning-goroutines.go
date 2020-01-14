package main

import (
	"fmt"
	"time"
)

/*
	Spawning goroutines

func f(n int) {
	for i := 0; i <= n; i++ {
		fmt.Printf("\r%d / %d", i, n)
		time.Sleep(1000 * time.Millisecond)
	}
}

func main() {
	for i := 0; i <= 100; i++ {
		go f(i)
	}
	var input string
	fmt.Scanln(&input)
}

*/

/*
	Channels
*/

func gift(c chan<- string) {
	for _, v := range `🎁🧸💎🍔🍩🍹` {
		c <- string(v)
	}
}

func receive(c <-chan string) {
	for i := 0; i < 10; i++ {
		gift := <-c
		fmt.Println("\rRecieved: ", gift)
		time.Sleep(time.Millisecond * 1000)
	}
}

func main() {
	c := make(chan string)
	go gift(c)
	go receive(c)
	var input string
	defer fmt.Scanln(&input)
}
