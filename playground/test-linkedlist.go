package main

import (
	"fmt"
	ll "heartyhardy/data-structs"
)

func main() {
	el1 := ll.Element{20, nil}
	el2 := ll.Element{30, nil}

	el1.Next = &el2

	fmt.Println(el1.Data, el1.Next.Data)
}
