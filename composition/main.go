package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Gamer struct {
	Person
	FavGame     string
	TimesPlayed int
}

func main() {
	// charith := Gamer{} can use this sytax and then use dot syntax to init
	charith := Gamer{
		FavGame:     "Diablo3",
		TimesPlayed: 10,
	}
	charith.Name = "Charith"
	charith.Age = 23

	fmt.Println(charith)
}
