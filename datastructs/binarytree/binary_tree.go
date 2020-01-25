package binarytree

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
