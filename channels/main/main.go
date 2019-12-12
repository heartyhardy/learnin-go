package main

import (
	"fmt"
	"heartyhardy/channels/astroids"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan astroids.Coordinate, 500)

	wg.Add(2)

	go func(ch <-chan astroids.Coordinate) {
		for res := range ch {
			fmt.Printf("# ( %v, %v ) \n", res.X, res.Y)
		}
		wg.Done()
	}(ch)

	go func(ch chan<- astroids.Coordinate) {
		astroids := astroids.ReadFile("../input/day-10-input.txt")
		for _, coord := range astroids {
			ch <- coord
		}
		close(ch)
		wg.Done()
	}(ch)

	wg.Wait()
}
