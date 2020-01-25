package main

import (
	"fmt"
	"time"
)

func moveEmoji(emoji string, distance int) {
	for i := 0; i < distance; i++ {
		padding := fmt.Sprint("\r%", i, "v  ")
		//fmt.Println("\033[2J")
		fmt.Printf(padding, emoji)
		time.Sleep(50 * time.Millisecond)
	}
	moveEmojiBack(emoji, distance)
}

func moveEmojiBack(emoji string, distance int) {
	for i := distance; i >= 0; i-- {
		padding := fmt.Sprint("\r%", i, "v  ")
		//fmt.Println("\033[2J")
		fmt.Printf(padding, emoji)
		time.Sleep(50 * time.Millisecond)
	}
}

// func main() {
// 	avatar := `🐯`
// 	distance := 100
// 	fmt.Print("\033[?25l")
// 	fmt.Print("\033[2J")
// 	moveEmoji(avatar, distance)
// 	fmt.Print("\033[?25h")
// }
