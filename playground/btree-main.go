package main

import (
	"fmt"
	"heartyhardy/trees/btree"
)

func main() {
	bst := btree.BinarySearchTree{}
	bst.Insert(5)
	bst.Insert(4)
	bst.Insert(10)
	bst.Insert(11)
	bst.Insert(12)
	bst.Insert(3)

	x := bst.Lookup(7)
	fmt.Println(x.Value)
	//bst.Root.Print(0, "root")
}
