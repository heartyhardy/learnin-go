package main

import (
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

	bst.Root.Traverse(0)
}
