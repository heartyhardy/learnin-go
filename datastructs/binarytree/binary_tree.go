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
	if btree.root == nil {
		btree.root = node
		return btree
	}
	destination := walk(btree.root, node)
	if destination.value <= node.value {
		destination.right = node
		return btree
	}
	destination.left = node
	return btree
}

func walk(node *Node, compare *Node) *Node {
	if node.left == nil && node.right == nil {
		return node
	}
	if node.left != nil && node.value > compare.value {
		return walk(node.left, compare)
	}
	if node.right != nil && node.value <= compare.value {
		return walk(node.right, compare)
	}
	return node
}
