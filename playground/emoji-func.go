package main

import (
	"fmt"
	"time"
)

func smile(t time.Duration) {
	fmt.Printf("\r%v\t", `😀`)
	time.Sleep(t * time.Second)
}

func meh(t time.Duration) {
	fmt.Printf("\r%v\t", `😑`)
	time.Sleep(t * time.Second)
}

func main() {
	smile(2)
	meh(3)
	smile(2)
	meh(2)
}
