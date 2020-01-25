package binarytree

import "testing"

func TestAdd(t *testing.T) {
	/*
		When Tree is empty
	*/
	t.Run("When Tree is Empty", func(t *testing.T) {
		tree := new(BinaryTree)
		node := new(Node)
		node.value = 10

		actual := tree.Add(node)

		if actual.root == nil {
			t.Errorf("Expected: %10v | Actual: %10v ", node, *actual)
		}
		if actual == nil {
			t.Errorf("Expected: %10v | Actual: %10v ", tree, *actual)
		}
		if actual.root.value != node.value {
			t.Errorf("Expected: %10v | Actual: %10v ", *node, *actual.root)
		}
	})

	/*
		When tree has a root and we insert a node which has
		a larger value than the root
	*/
	t.Run("When root has lower value than new Node", func(t *testing.T) {

		expected := new(BinaryTree)
		rootNode := new(Node)
		rootNode.value = 10
		newNode := new(Node)
		newNode.value = 20

		expected.root = rootNode
		expected.root.right = newNode

		actual := new(BinaryTree)
		actualRoot := new(Node)
		actualRoot.value = 10
		actualNewNode := new(Node)
		actualNewNode.value = 20

		actual = actual.Add(actualRoot).Add(actualNewNode)

		if actual.root.right == nil {
			t.Errorf("\nExpected: %5v \nActual: %5v ", *actual.root.left, *expected.root.left)
		}
		if actual.root.value != expected.root.value {
			t.Errorf("\nExpected: %5v \nActual: %5v ", actual.root.value, expected.root.value)
		}
		if actual.root.left != nil {
			t.Errorf("\nExpected: %5v \nActual: %5v ", actual.root.left.value, expected.root.left.value)
		}

		expected.root.left = new(Node)
		expected.root.left.value = 5

		anotherNewNode := new(Node)
		anotherNewNode.value = 5
		actual.Add(anotherNewNode)

		if actual.root.value != expected.root.value {
			t.Errorf("\nExpected: %5v \nActual: %5v ", actual.root.value, expected.root.value)
		}
		if actual.root.left.value != expected.root.left.value {
			t.Errorf("\nExpected: %5v \nActual: %5v ", actual.root.left.value, expected.root.left.value)
		}
	})

}

func TestWalk(t *testing.T) {
	expected := new(BinaryTree)
	root := new(Node)
	root.value = 10
	newnodeL := new(Node)
	newnodeL.value = 5
	newnodeR := new(Node)
	newnodeR.value = 20
	expected.Add(root)

	compare := new(Node)
	compare.value = 5
	actual := walk(expected.root, compare)

	if actual != expected.root {
		t.Errorf("Expected: %10v | Actual: %10v ", expected.root, actual)
	}

	compare.value = 10
	actual = walk(expected.root, compare)

	if actual != expected.root {
		t.Errorf("Expected: %10v | Actual: %10v ", expected.root, actual)
	}

	result := expected.Add(newnodeL)

	if result.root.left.value != newnodeL.value {
		t.Errorf("Expected: %10v | Actual: %10v ", newnodeL.value, result.root.left.value)
	}

	result = expected.Add(newnodeR)

	if result.root.right.value != newnodeR.value {
		t.Errorf("Expected: %10v | Actual: %10v ", newnodeL.value, result.root.right.value)
	}

}
