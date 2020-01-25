package binarytree

import "testing"

func TestAdd(t *testing.T) {
	t.Run("It should add a root", func(t *testing.T) {
		root := &Node{value: 10}
		expected := &BinaryTree{root: root}
		actual := &BinaryTree{}
		actualRoot := &Node{value: 10}
		actual = actual.Add(actualRoot)

		if actual == nil || actual.root == nil {
			t.Error("Should not return nil, expected a btree with a root")
		}
		if actual.root.value != expected.root.value {
			t.Errorf("Values of roots should be identical! expected: %v actual: %v", expected.root.value, actual.root.value)
		}
	})

	t.Run("It should child node to left", func(t *testing.T) {
		root := &Node{value: 10}
		expected := &BinaryTree{root: root}
		expectedNewChild := &Node{value: 5}
		expected.root.left = expectedNewChild
		actual := &BinaryTree{}
		actualRoot := &Node{value: 10}
		actualNewChild := &Node{value: 5}
		actual = actual.Add(actualRoot).Add(actualNewChild)

		if actual.root.left == nil {
			t.Errorf("Should not return nil, expected: %v ", expected.root.left)
		}
		if actual.root.left.value != expected.root.left.value {
			t.Errorf("Values of roots should be identical! expected: %v actual: %v", expected.root.left.value, actual.root.left.value)
		}
	})

	t.Run("It should add child node to right", func(t *testing.T) {
		root := &Node{value: 10}
		expected := &BinaryTree{root: root}
		expectedNewChild := &Node{value: 15}
		expected.root.right = expectedNewChild
		actual := &BinaryTree{}
		actualRoot := &Node{value: 10}
		actualNewChild := &Node{value: 15}
		actual = actual.Add(actualRoot).Add(actualNewChild)

		if actual.root.right == nil {
			t.Errorf("Should not return nil, expected: %v ", expected.root.right)
		}
		if actual.root.right.value != expected.root.right.value {
			t.Errorf("Values of roots should be identical! expected: %v actual: %v", expected.root.right.value, actual.root.right.value)
		}
	})

	t.Run("It should return nil because input is nil", func(t *testing.T) {
		root := &Node{value: 10}
		expected := &BinaryTree{root: root}
		var expectedNewChild *Node = nil
		expected.root.left = expectedNewChild
		actual := &BinaryTree{}
		actualRoot := &Node{value: 10}
		var actualNewChild *Node = nil
		actual = actual.Add(actualRoot).Add(actualNewChild)

		if actual != nil {
			t.Errorf("Should return nil, expected: %v actual %v ", nil, actual)
		}
	})
}

func TestWalkTo(t *testing.T) {
	t.Run("Left node should not be null", func(t *testing.T) {
		expectedParent := &Node{value: 10}
		expectedChild := &Node{value: 5}
		expectedParent.left = expectedChild
		actualParent := &Node{value: 10}
		actualChild := &Node{value: 5}

		walkToAndAdd(actualParent, actualChild)

		if actualParent.left == nil {
			t.Errorf("Left should not be null, expected: %v actual %v ", expectedParent.left, actualParent.left)
		}
		if actualParent.left.value != expectedParent.left.value {
			t.Errorf("Should add child to parent -> left, expected: %v actual %v ", expectedParent.left.value, actualParent.left.value)
		}
	})

	t.Run("Right node should not be null", func(t *testing.T) {
		expectedParent := &Node{value: 10}
		expectedChild := &Node{value: 15}
		expectedParent.right = expectedChild
		actualParent := &Node{value: 10}
		actualChild := &Node{value: 15}

		walkToAndAdd(actualParent, actualChild)

		if actualParent.right == nil {
			t.Errorf("right should not be null, expected: %v actual %v ", expectedParent.right, actualParent.right)
		}
		if actualParent.right.value != expectedParent.right.value {
			t.Errorf("Values should be identical: %v actual %v ", expectedParent.right.value, actualParent.right.value)
		}
	})

	t.Run("Left node should match the inserted value", func(t *testing.T) {
		expectedParent := &Node{value: 10}
		expectedChild := &Node{value: 5}
		expectedChildL := &Node{value: 3}
		expectedParent.left = expectedChild
		expectedParent.left.left = expectedChildL
		actualParent := &Node{value: 10}
		actualChild := &Node{value: 5}
		actualChildL := &Node{value: 3}

		walkToAndAdd(actualParent, actualChild)
		walkToAndAdd(actualParent, actualChildL)

		if actualParent.left.left == nil {
			t.Errorf("right should not be null, expected: %v actual %v ", expectedParent.left.left, actualParent.left.left)
		}
		if actualParent.left.left.value != expectedParent.left.left.value {
			t.Errorf("Values should be identical, expected: %v actual %v ", expectedParent.right.value, actualParent.right.value)
		}
	})

	t.Run("Right node should match the inserted value", func(t *testing.T) {
		expectedParent := &Node{value: 10}
		echild1 := &Node{value: 5}
		echild2 := &Node{value: 15}
		echild3 := &Node{value: 3}
		echild4 := &Node{value: 20}
		expectedParent.left = echild1
		expectedParent.right = echild2
		expectedParent.left.left = echild3
		expectedParent.right.right = echild4

		actualParent := &Node{value: 10}
		achild1 := &Node{value: 5}
		achild2 := &Node{value: 15}
		achild3 := &Node{value: 3}
		achild4 := &Node{value: 20}

		walkToAndAdd(actualParent, achild1)
		walkToAndAdd(actualParent, achild2)
		walkToAndAdd(actualParent, achild3)
		walkToAndAdd(actualParent, achild4)

		if actualParent.left.left == nil {
			t.Errorf("right should not be null, expected: %v actual %v ", expectedParent.left.left, actualParent.left.left)
		}
		if actualParent.left.left.value != expectedParent.left.left.value {
			t.Errorf("Values should be identical, expected: %v actual %v ", expectedParent.left.left.value, actualParent.left.left.value)
		}
		if actualParent.right.right == nil {
			t.Errorf("right should not be null, expected: %v actual %v ", expectedParent.right.right, actualParent.right.right)
		}
		if actualParent.right.right.value != expectedParent.right.right.value {
			t.Errorf("Values should be identical, expected: %v actual %v ", expectedParent.right.right.value, actualParent.right.right.value)
		}
	})
}

func TestPrint(t *testing.T) {

}
