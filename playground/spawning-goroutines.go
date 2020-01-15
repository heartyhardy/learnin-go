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

/*

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
	fmt.Scanln(&input)
}

*/

/*

func fib(n, first, second int) int {
	memo := map[int]int{}
	if v, ok := memo[n]; ok {
		return v
	}
	if n == 0 {
		return first
	}
	if n == 1 {
		return second
	}
	res := fib(n-1, second, first+second)
	memo[n] = res
	return res
}

func calcFibs(ch chan<- struct{ n, r int }) {
	for i := 0; i <= 50; i++ {
		fibres := struct{ n, r int }{i, fib(i, 1, 0)}
		ch <- fibres
	}
}

func printFibs(ch <-chan struct{ n, r int }) {
	for i := 0; i <= 50; i++ {
		res := <-ch
		fmt.Printf("\nFib (%v): %v ", res.n, res.r)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	ch := make(chan struct{ n, r int })
	go calcFibs(ch)
	go printFibs(ch)

	var input string
	fmt.Scanln(&input)
}

*/

/*
	Select
*/

func yang(ch chan<- string) {
	for {
		ch <- `🤡`
		time.Sleep(2 * time.Second)
	}
}

func yin(ch chan<- string) {
	for {
		ch <- `👹`
		time.Sleep(3 * time.Second)
	}
}

func yinyang(yinch chan string, yangch chan string) {
	for {
		select {
		case yinmsg := <-yinch:
			fmt.Printf("\vYin: %v   ", yinmsg)
		case yangmsg := <-yangch:
			fmt.Printf("\vYang: %v   ", yangmsg)
		}
	}
}

func main() {
	yinch := make(chan string)
	yangch := make(chan string)
	go yin(yinch)
	go yang(yangch)
	go yinyang(yinch, yangch)

	var input string
	fmt.Scanln(&input)
}
