package main

import (
	"fmt"
	"heartyhardy/packages/linkedlist"
)

func main() {
	newRobot1 := new(linkedlist.Robot)
	newRobot2 := new(linkedlist.Robot)

	newRobot1.ID = 1
	newRobot1.Name = "Andrea"

	newRobot2.ID = 2
	newRobot2.Name = "Leah"

	ll := new(linkedlist.Robots)
	ll.Append(newRobot1)
	ll.Append(newRobot2)
	fmt.Println(ll.Get().Next())
}
