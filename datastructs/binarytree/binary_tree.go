package binarytree

import (
	"fmt"
	"strings"
)

//Node - Tree Node
type Node struct {
	key   int
	value int64
	left  *Node
	right *Node
}

//BinaryTree ...
type BinaryTree struct {
	root *Node
}

//Add a node to the tree
func (btree *BinaryTree) Add(node *Node) *BinaryTree {
	if node == nil {
		return nil
	}
	if btree.root == nil {
		btree.root = node
		return btree
	}
	walkToAndAdd(btree.root, node)
	return btree
}

//walkToAndAdd search the tree and adds the child
//after comparing the values of parent and child
func walkToAndAdd(parent *Node, child *Node) {
	if parent.value >= child.value {
		if parent.left == nil {
			parent.left = child
			return
		}
		walkToAndAdd(parent.left, child)
		return
	}
	if parent.right == nil {
		parent.right = child
		return
	}
	walkToAndAdd(parent.right, child)
}

func walkAndPrint(parent *Node, lvl int) {
	if parent == nil {
		return
	}
	padded := strings.Repeat(".", lvl)
	fmt.Printf("\n%v %v\n", padded, parent.value)
	walkAndPrint(parent.left, lvl-5)
	walkAndPrint(parent.right, lvl+5)
}

func walkAndSearch(parent *Node, val int64) *Node {
	if parent == nil {
		return nil
	}
	if parent.value == val {
		return parent
	} else if parent.value >= val {
		return walkAndSearch(parent.left, val)
	}
	return walkAndSearch(parent.right, val)
}

//Search will traverse the tree and find the value
func (btree *BinaryTree) Search(val int64) *Node {
	node := walkAndSearch(btree.root, val)
	return node
}

//Print the tree to console
func (btree *BinaryTree) Print() {
	walkAndPrint(btree.root, 50)
}
