package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

type Gamer struct {
	Person
	FavGame     string
	TimesPlayed int
}

type Animal struct {
	Name string `required max:"10"`
}

func main() {
	// charith := Gamer{} can use this sytax and then use dot syntax to init
	charith := Gamer{
		FavGame:     "Diablo3",
		TimesPlayed: 10,
	}
	charith.Name = "Charith"
	charith.Age = 23

	chaze := Gamer{
		Person:      Person{Name: "Chaze", Age: 26},
		FavGame:     "Witcher",
		TimesPlayed: 3,
	}

	animal := Animal{Name: "sdsdd"}

	fmt.Println(charith)
	fmt.Println(chaze)

	t := reflect.TypeOf(animal)
	feild, _ := t.FieldByName("Name")
	fmt.Println(feild.Tag)
}
