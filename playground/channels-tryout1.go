// package main

// import (
// 	"fmt"
// 	"sync"
// )

// var wg = sync.WaitGroup{}

// func main() {
// 	ch := make(chan int, 50)

// 	wg.Add(2)
// 	go func(ch <-chan int) {
// 		//res := <-ch
// 		for res := range ch {
// 			fmt.Println("RECIEVED: ", res)
// 		}
// 		wg.Done()
// 	}(ch)

// 	go func(ch chan<- int) {
// 		for send := 1; send < 11; send++ {
// 			ch <- 10 * send
// 		}
// 		close(ch)
// 		wg.Done()
// 	}(ch)

// 	wg.Wait()
// }
