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

//Traverse ...
func (n *Node) Traverse(lvl int) {
	if n == nil {
		return
	}

	if n.hasChildren() {
		fmt.Println(n.Value, lvl)
		lvl++
		n.Left.Traverse(lvl)
		n.Right.Traverse(lvl)
	} else {
		fmt.Println(n.Value, lvl)
	}
}
