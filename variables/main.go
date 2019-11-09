package main

import (
	"fmt"
	"strconv"
)

func main() {
	i := "100"

	res, err := strconv.Atoi(i)

	var j int = 50
	res2 := strconv.Itoa(j)

	if err == nil {
		fmt.Printf("%v, %T \n", res, res)
		fmt.Printf("%v, %T \n", res2, res2)
	}

}
