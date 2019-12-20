package stack

import (
	"testing"
)

func TestPush(t *testing.T) {
	s := Stack{Start: nil, End: nil, length: 0}
	node := new(Node)
	node.Value = "Hello"

	sptr := s.Push(node)

	if sptr.length <= 0 {
		t.Error("Expected Length: 1 Got:", sptr.length)
	}

	if sptr.Start == nil || sptr.End == nil {
		t.Error("Start and End pointers should not be nil after PUSH")
	}
}

func TestPeek(t *testing.T) {
	s := Stack{Start: nil, End: nil, length: 0}
	node := new(Node)
	node.Value = "Hello"

	sptr := s.Push(node)

	if sptr.Peek() == nil {
		t.Error("Peek() should not return nil!")
	}

	peekval := (sptr.Peek().Value).(*Node)
	if peekval.Value != "Hello" {
		t.Error("Peek() should return Hello")
	}
}

func TestPeekVal(t *testing.T) {
	s := Stack{Start: nil, End: nil, length: 0}
	node := new(Node)
	node.Value = "Hello"

	sptr := s.Push(node)

	if sptr.PeekVal().Value != "Hello" {
		t.Error("PeekVal() should return Hello")
	}
}

func TestPeekEnd(t *testing.T) {
	s := Stack{Start: nil, End: nil, length: 0}
	node1 := new(Node)
	node1.Value = "Hello"
	node2 := new(Node)
	node2.Value = "World"

	sptr := s.Push(node1)
	sptr = s.Push(node2)

	if sptr.PeekEnd() == nil {
		t.Error("PeekEnd should not be nil")
	}

	peekval := (sptr.PeekEnd().Value).(*Node)
	if peekval.Value != "Hello" {
		t.Error("PeekEnd value should be Hello")
	}

	peekend := (sptr.PeekEnd().Prev.Value).(*Node)

	if peekend.Value != "World" {
		t.Error("PeekEnd().Prev.Value should be World")
	}
}
