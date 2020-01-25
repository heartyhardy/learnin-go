package binarytree

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"testing"
)

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

func TestSearch(t *testing.T) {

	btree := &BinaryTree{}
	children := []int{40, 7, 39, 37, 14, 42, 43, 24, 13, 44, 12, 30, 20, 16, 26, 38, 29, 50, 9, 28, 25, 10, 1, 21, 41, 48, 45, 2, 49, 18, 5, 8, 32, 47, 27, 33, 4, 34, 17, 36, 15, 31, 19, 46, 6, 23, 35, 11, 22, 3}
	for _, v := range children {
		child := &Node{value: int64(v)}
		btree.Add(child)
	}

	t.Run("It should return nil when value not found", func(t *testing.T) {
		var expected *Node = nil
		actual := btree.Search(int64(9999))

		if actual != nil {
			t.Errorf("Search result should return nil, expected: %v actual %v ", expected, actual)
		}
	})

	for _, child := range children {
		t.Run("It should find and return each value", func(t *testing.T) {
			expected := &Node{value: int64(child)}
			actual := btree.Search(int64(child))

			if actual == nil {
				t.Errorf("Search result should return 50, expected: %v actual %v ", expected.value, actual.value)
			}
			if actual.value != expected.value {
				t.Errorf("Search result should match, expected: %v actual %v ", expected.value, actual.value)
			}
		})
	}

}

func BenchmarkSearch(t *testing.B) {
	btree := &BinaryTree{}
	data, _ := ioutil.ReadFile("benchmark.txt")
	fields := strings.Fields(string(data))

	for _, val := range fields {
		v64, _ := strconv.ParseInt(val, 10, 64)
		node := &Node{value: v64}
		btree.Add(node)
	}
	t.Run("Search a node in this array", func(t *testing.B) {
		var j int64
		for j = 0; j < int64(len(fields)); j = j + 10 {
			v64, _ := strconv.ParseInt(fields[j], 10, 64)
			res := btree.Search(v64)
			fmt.Println("Result: ", res.value)
		}
	})
}

func TestPrint(t *testing.T) {
	btree := &BinaryTree{}
	children := []int{40, 7, 39, 37, 14, 42, 43, 24, 13, 44, 12, 30, 20, 16, 26, 38, 29, 50, 9, 28, 25, 10, 1, 21, 41, 48, 45, 2, 49, 18, 5, 8, 32, 47, 27, 33, 4, 34, 17, 36, 15, 31, 19, 46, 6, 23, 35, 11, 22, 3}
	for _, v := range children {
		child := &Node{value: int64(v)}
		btree.Add(child)
	}
	btree.Print()
}
