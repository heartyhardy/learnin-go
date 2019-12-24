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

func (n *Node) getChildCount() int {
	count := 0
	if n.hasChildren() {
		if n.Left != nil {
			count++
		}
		if n.Right != nil {
			count++
		}
		return count
	}
	return 0
}

func (n *Node) getsOnlyChild() *Node {
	if n.hasChildren() {
		if n.Left != nil {
			return n.Left
		} else if n.Right != nil {
			return n.Right
		}
	}
	return nil
}

func (n *Node) getMinimumLeftNode() *Node {
	if n.hasChildren() {
		currentNode := n.Left

		for {
			if currentNode.Left != nil {
				currentNode = currentNode.Left
			} else {
				break
			}
		}
		return currentNode
	}
	return nil
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

//Remove a node from bst
func (bst *BinarySearchTree) Remove(value int) {
	currentNode := bst.Root

	for {
		if currentNode != nil {
			switch {
			case currentNode.Value == value:
				//Found node
				//See if it has children
				// if YES go right
				// if NO can delete
				if !currentNode.hasChildren() {
					// Delete current node
					currentNode.Value = 0
					currentNode.Left = nil
					currentNode.Right = nil
					currentNode = nil
					return
				}

				switch currentNode.getChildCount() {
				case 1:
					// 1 child method
					n := currentNode.getsOnlyChild()
					currentNode.Value = n.Value
					currentNode.Left = nil
					currentNode.Right = nil
					n = nil
					return
				case 2:
					// 2 children method
					minLeft := currentNode.Right.getMinimumLeftNode()
					if minLeft != nil {
						currentNode.Value = minLeft.Value
						minLeft = nil
						return
					}
					currentNode.Value = currentNode.Right.Value
					currentNode.Right = nil
					return
				}
			case currentNode.Value > value:
				// Go Left
				currentNode = currentNode.Left
			case currentNode.Value < value:
				// Go Right
				currentNode = currentNode.Right
			}
		} else {
			break
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
