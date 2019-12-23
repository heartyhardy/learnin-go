package btree

import "fmt"

//Node represents a valid bst child
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func (n *Node) hasChildren() bool {
	if n == nil {
		return false
	}
	if n.Left != nil || n.Right != nil {
		return true
	}
	return false
}

//BinarySearchTree ...
type BinarySearchTree struct {
	Root   *Node
	Length int
}

//Insert - appends a node after determining its position in the tree
func (bst *BinarySearchTree) Insert(value int) {
	if bst.Root == nil {
		node := new(Node)
		node.Value = value
		bst.Root = node
	} else {
		// compare given value with current node's value
		// if value > current value go right
		// else go left
		// if equal ignore
		// repeat until a proper insert location found
		currentNode := bst.Root
		currentVal := currentNode.Value
		for {
			if value > currentVal {
				if currentNode.Right != nil {
					currentNode = currentNode.Right
				} else {
					currentNode.Right = new(Node)
					currentNode.Right.Value = value
					break
				}
			} else if value < currentVal {
				if currentNode.Left != nil {
					currentNode = currentNode.Left
				} else {
					currentNode.Left = new(Node)
					currentNode.Left.Value = value
					break
				}
			}
		}
	}
}

//Lookup searches for a given value inside the tree
func (bst *BinarySearchTree) Lookup(value int) *Node {
	currentNode := bst.Root
	for {
		if currentNode != nil {
			switch {
			case currentNode.Value == value:
				return currentNode
			case currentNode.Value < value:
				currentNode = currentNode.Right
			case currentNode.Value > value:
				currentNode = currentNode.Left
			}
		} else {
			return currentNode
		}
	}
}

//Print ...
func (n *Node) Print(lvl int, side string) {
	if n == nil {
		return
	}

	if n.hasChildren() {
		fmt.Println(n.Value, lvl, side)
		lvl++
		n.Left.Print(lvl, "left")
		n.Right.Print(lvl, "right")
	} else {
		fmt.Println(n.Value, lvl, side)
	}
}
